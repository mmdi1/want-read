package main

import (
	"context"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// App struct
type App struct {
	ctx context.Context
}

var (
	OLD_GWL_EXSTYLE int32 = 327937
	APP_HUND        win.HWND
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}
func settingWinDefault() {
	win.SetWindowLong(APP_HUND, win.GWL_EXSTYLE, win.WS_EX_LAYERED)
}
func settingWinTransfer() {
	win.SetWindowLong(APP_HUND, win.GWL_EXSTYLE, win.WS_EX_LAYERED)
}

const (
	WINEVENT_OUTOFCONTEXT    = 0x0000
	EVENT_SYSTEM_FOREGROUND  = 0x0003
	EVENT_OBJECT_FOCUS       = 0x8005
	EVENT_OBJECT_STATECHANGE = 0x800A
	WM_KEYDOWN               = 0x0100
	WM_KEYUP                 = 0x0101
	WM_SYSKEYDOWN            = 0x0104
	WM_SYSKEYUP              = 0x0105
)

var (
	user32DLL                = syscall.NewLazyDLL("user32.dll")
	getForegroundWindow      = user32DLL.NewProc("GetForegroundWindow")
	getWindowThreadProcessId = user32DLL.NewProc("GetWindowThreadProcessId")
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

func winEventProc(hWinEventHook win.HWINEVENTHOOK, event uint32, hwnd win.HWND, idObject, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr {
	switch event {
	case EVENT_SYSTEM_FOREGROUND:
		var processID, threadID uint32
		threadID = win.GetWindowThreadProcessId(hwnd, &processID)
		if threadID == idEventThread {
			fmt.Println("Foreground window changed:", hwnd)
		}
		return uintptr(processID)
	case WM_KEYDOWN:
		fmt.Println("----------->")
	case WM_SYSKEYDOWN:
		fmt.Println("WM_SYSKEYDOWN")
	case WM_SYSKEYUP:
		fmt.Println("WM_SYSKEYUP")
	default:
		fmt.Println("default", event)
	}
	return uintptr(0)
}

// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	str, err := syscall.UTF16PtrFromString(title)
	if err == nil {
		APP_HUND = win.FindWindow(nil, str)
	}
	// // settingWinTransfer()
	// hook, err := win.SetWinEventHook(
	// 	EVENT_SYSTEM_FOREGROUND,
	// 	WM_SYSKEYUP,
	// 	0,
	// 	winEventProc,
	// 	0,
	// 	0,
	// 	WINEVENT_OUTOFCONTEXT,
	// )
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// if hook == 0 {
	// 	fmt.Println("SetWinEventHook failed")
	// 	return
	// }
	// defer win.UnhookWinEvent(hook)

	// var msg win.MSG
	// for win.GetMessage(&msg, 0, 0, 0) != 0 {
	// 	win.TranslateMessage(&msg)
	// 	win.DispatchMessage(&msg)
	// }
	user32 := syscall.MustLoadDLL("user32.dll")
	defer user32.Release()
	// 注册键盘按下事件
	eventHook, err := win.SetWinEventHook(0, win.EVENT_OBJECT_VALUECHANGE+1000000, 0, winEventCallback, 0, 0, win.WINEVENT_OUTOFCONTEXT)
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

	// 卸载事件钩子
	win.UnhookWinEvent(eventHook)
}

// 定义回调函数，处理事件
func winEventCallback(hWinEventHook win.HWINEVENTHOOK, eventType uint32, hwnd win.HWND, idObject int32, idChild int32, dwEventThread uint32, dwmsEventTime uint32) uintptr {
	switch eventType {
	default:
		if eventType != 32779 {
			fmt.Println("event", eventType)
		}
		return 0
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ReceiveFilePath(path string) {

}
func (a *App) SettingDefaultWindows() {
	settingWinDefault()
}
