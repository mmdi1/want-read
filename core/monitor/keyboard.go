package monitor

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
	"want-read/configs"
	"want-read/server/api/ws"

	"github.com/lxn/win"
)

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 0x0100
	WM_KEYUP       = 0x0101
	WM_SYSKEYDOWN  = 0x0104
	WM_SYSKEYUP    = 0x0105
)

type KBDLLHOOKSTRUCT struct {
	VkCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DwExtraInfo uint32
}
type keyOpt struct {
	Code   uint32
	IsDown bool
}

const (
	PrevPage = iota
	NextPage
	HidePanel
)

type operationkey struct {
	name    string
	handler int
	group   []uint32
}

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	setWindowsHookEx = user32.MustFindProc("SetWindowsHookExW")
	callNextHookEx   = user32.MustFindProc("CallNextHookEx")
	// getMessage          = user32.MustFindProc("GetMessageW")
	// translateMessage    = user32.MustFindProc("TranslateMessage")
	// dispatchMessage     = user32.MustFindProc("DispatchMessageW")
	unhookWindowsHookEx = user32.MustFindProc("UnhookWindowsHookEx")
	hookID              uintptr
	err                 error
	downChan            = make(chan keyOpt, 1)
	downAllKeys         = map[uint32]bool{}
	settings            = []operationkey{
		{
			name:    "上一页",
			handler: PrevPage,
			group:   []uint32{164, 188},
		},
		{
			name:    "下一页",
			handler: NextPage,
			group:   []uint32{164, 190},
		},
		{
			name:    "隐藏",
			handler: NextPage,
			group:   []uint32{164, 77},
		},
	}
)

func hasgroupKey() {
	for i := 0; i < len(settings); i++ {
		need_len := 0
		down_len := map[uint32]bool{}
		for x := 0; x < len(settings[i].group); x++ {
			need_len = len(settings[x].group)
			_, ok := downAllKeys[settings[i].group[x]]
			if ok {
				down_len[settings[i].group[x]] = true
			}
		}
		if need_len == len(downAllKeys) && need_len == len(down_len) {
			out := map[string]any{}
			out["msgId"] = settings[i].handler
			switch settings[i].handler {
			case PrevPage:
				configs.CurrentPage--
				if configs.CurrentPage < 0 {
					configs.CurrentPage = 0
					return
				}
				out["data"] = string(configs.CurrentReadBook[configs.CurrentPage])
			case NextPage:
				if configs.CurrentPage == len(configs.CurrentReadBook) {
					return
				}
				configs.CurrentPage++
				out["data"] = string(configs.CurrentReadBook[configs.CurrentPage])
			case HidePanel:
				configs.IS_HIDE = !configs.IS_HIDE
				out["data"] = configs.IS_HIDE
			}
			ws.Conn.SendMsg(out)
		}
	}
}
func handler() {
	for {
		opt := <-downChan
		if opt.IsDown {
			ok := downAllKeys[opt.Code]
			if ok {
				continue
			}
			downAllKeys[opt.Code] = true
			hasgroupKey()
		} else {
			delete(downAllKeys, opt.Code)
		}
	}

}
func keyboardProc(nCode, wParam, lParam uintptr) uintptr {
	// ok := lock.TryLock()
	// if !ok {
	// 	return 0
	// }
	// defer lock.Unlock()
	kbd := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
	switch wParam {
	case WM_KEYDOWN, WM_SYSKEYDOWN:
		downChan <- keyOpt{
			IsDown: true,
			Code:   kbd.VkCode,
		}
		fmt.Println("按下Key pressed:", kbd.VkCode)
	case WM_KEYUP, WM_SYSKEYUP:
		fmt.Println("抬起Key pressed:", kbd.VkCode)
		downChan <- keyOpt{
			IsDown: false,
			Code:   kbd.VkCode,
		}
	}
	ret, _, _ := callNextHookEx.Call(hookID, nCode, wParam, lParam)
	return ret
}
func KeyBoardHandler() {
	hookID, _, err = setWindowsHookEx.Call(
		WH_KEYBOARD_LL,
		syscall.NewCallback(keyboardProc),
		0,
		0,
	)
	if err != nil {
		log.Println("hook call err:", err)
	}
	if hookID == 0 {
		fmt.Println("Failed to set up hook")
		return
	}
	defer unhookWindowsHookEx.Call(hookID)
	// 无限循环，等待事件发生
	go handler()
	for {
		var msg win.MSG
		ok := win.GetMessage(&msg, 0, 0, 0)
		if ok == win.FALSE {
			fmt.Println("GetMessage error:", err)
			break
		}
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}
