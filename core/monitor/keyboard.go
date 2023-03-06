package monitor

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
	"want-read/configs"
	"want-read/core/db"
	"want-read/server/api/ws"

	"github.com/lxn/win"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

var (
	user32              = syscall.MustLoadDLL("user32.dll")
	setWindowsHookEx    = user32.MustFindProc("SetWindowsHookExW")
	callNextHookEx      = user32.MustFindProc("CallNextHookEx")
	unhookWindowsHookEx = user32.MustFindProc("UnhookWindowsHookEx")
	hookID              uintptr
	err                 error
	downChan            = make(chan keyOpt, 1)
	downAllKeys         = map[uint32]bool{}
)

// 更新读书进度
func updateReadProcess() {
	configs.ReadBook.ReadSize = configs.CurrentPage * configs.SettingModel.ShowSize
	configs.ReadBook.UpdateAt = time.Now()
	db.Set(db.K_Read+configs.ReadBook.IdKey, configs.ReadBook)
}

func hasGroupKey() {
	for i := 0; i < len(configs.SettingGroup); i++ {
		need_len := 0
		down_len := map[uint32]bool{}
		for x := 0; x < len(configs.SettingGroup[i].Group); x++ {
			need_len = len(configs.SettingGroup[x].Group)
			_, ok := downAllKeys[configs.SettingGroup[i].Group[x]]
			if ok {
				down_len[configs.SettingGroup[i].Group[x]] = true
			}
		}
		if need_len == len(downAllKeys) && need_len == len(down_len) {
			out := db.WsMsgModel{}
			out.Id = configs.SettingGroup[i].Handler
			switch configs.SettingGroup[i].Handler {
			case db.WsID_PrevPage:
				configs.CurrentPage--
				if configs.CurrentPage < 0 {
					configs.CurrentPage = 0
					return
				}
				updateReadProcess()
				out.Data = string(configs.ReadBook.Conetent[configs.CurrentPage])
			case db.WsID_NextPage:
				if configs.CurrentPage == len(configs.ReadBook.Conetent) {
					return
				}
				configs.CurrentPage++
				updateReadProcess()
				out.Data = string(configs.ReadBook.Conetent[configs.CurrentPage])
			case db.WsID_HidePanel:
				fmt.Println("=============>", need_len, down_len, downAllKeys)
				configs.IS_HIDE = !configs.IS_HIDE
				if configs.IS_HIDE {
					runtime.WindowHide(configs.APP_CTX)
				} else {
					runtime.WindowShow(configs.APP_CTX)
				}
				out.Data = configs.IS_HIDE
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
			hasGroupKey()
		} else {
			delete(downAllKeys, opt.Code)
		}
	}
}
func keyboardProc(nCode, wParam, lParam uintptr) uintptr {
	kbd := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
	fmt.Println("================>", kbd.VkCode)
	switch wParam {
	case WM_KEYDOWN, WM_SYSKEYDOWN:
		fmt.Println("按下kkkkk:", kbd.VkCode)
		downChan <- keyOpt{
			IsDown: true,
			Code:   kbd.VkCode,
		}
	case WM_KEYUP, WM_SYSKEYUP:
		fmt.Println("takkkkk:", kbd.VkCode)
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
		log.Println("failed to set up hook")
		return
	}
	defer unhookWindowsHookEx.Call(hookID)
	// 无限循环，等待事件发生
	go handler()
	for {
		var msg win.MSG
		ok := win.GetMessage(&msg, 0, 0, 0)
		if ok == win.FALSE {
			log.Println("GetMessage error:", err)
			break
		}
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}
