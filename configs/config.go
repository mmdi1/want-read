package configs

import (
	"context"
	"want-read/core/db"

	"github.com/lxn/win"
)

var (
	OLD_GWL_EXSTYLE int32 = 327937
	APP_HUND        win.HWND
	APP_CTX         context.Context
	APP_Width       = 600
	APP_Height      = 396
	IS_HIDE         = false
	ReadBook        *db.Book
	CurrentPage     int
	SettingModel    = &db.Setting{}
	SettingGroup    = []db.Operationkey{
		{
			Name:    "上一页",
			Handler: db.WsID_PrevPage,
			Group:   []uint32{164, 188},
		},
		{
			Name:    "下一页",
			Handler: db.WsID_NextPage,
			Group:   []uint32{164, 190},
		},
		{
			Name:    "隐藏",
			Handler: db.WsID_HidePanel,
			Group:   []uint32{164, 77},
		},
	}
)
