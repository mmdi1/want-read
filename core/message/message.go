package message

import (
	"want-read/configs"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func TipInfo(msg string) {
	runtime.MessageDialog(configs.APP_CTX, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "提示",
		Message:       msg,
		DefaultButton: "确定",
	})
}

func TipErr(msg string) {
	runtime.MessageDialog(configs.APP_CTX, runtime.MessageDialogOptions{
		Type:          runtime.ErrorDialog,
		Title:         "错误",
		Message:       msg,
		DefaultButton: "确定",
	})
}
