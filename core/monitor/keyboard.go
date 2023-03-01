package monitor

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

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
)

func keyboardProc(nCode, wParam, lParam uintptr) uintptr {
	kbd := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
	fmt.Println("==============>", wParam, kbd.VkCode)
	switch wParam {
	case WM_KEYDOWN, WM_SYSKEYDOWN:
		fmt.Println("按下Key pressed:", kbd.VkCode)
	case WM_KEYUP:
		fmt.Println("抬起Key pressed:", kbd.VkCode)
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
