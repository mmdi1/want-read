package setting

import (
	"encoding/json"
	"log"
	"sync"
	"want-read/configs"
	"want-read/core/db"
	"want-read/server/api/ws"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	sync.Mutex
}

func NewApp() *App {
	return &App{}
}
func updateSetting(parm db.Setting) {
	configs.SettingGroup = []db.Operationkey{
		{
			Name:    "上一页",
			Handler: db.WsID_PrevPage,
			Group:   parm.PrevGroup,
		},
		{
			Name:    "下一页",
			Handler: db.WsID_NextPage,
			Group:   parm.NextGroup,
		},
		{
			Name:    "隐藏",
			Handler: db.WsID_HidePanel,
			Group:   parm.HideGroup,
		},
	}
}
func initSettingModel() *db.Setting {
	return &db.Setting{
		PrevGroup: []uint32{164, 188},
		NextGroup: []uint32{164, 190},
		HideGroup: []uint32{164, 77},
		FontSize:  12,
		FontColor: "rgb(120,120,120)",
		ShowSize:  50,
	}
}

// 保存配置
func (sef *App) SaveSetting(parm db.Setting) bool {
	if parm.FontSize < 12 {
		parm.FontSize = 12
	}
	if parm.ShowSize <= 0 {
		parm.ShowSize = 50
	}
	if len(parm.PrevGroup) == 0 {
		parm.PrevGroup = []uint32{164, 188}
	}
	if len(parm.NextGroup) == 0 {
		parm.NextGroup = []uint32{164, 190}
	}
	if len(parm.HideGroup) == 0 {
		parm.HideGroup = []uint32{164, 77}
	}
	err := db.Set(db.K_Setting, parm)
	if err == nil {
		updateSetting(parm)
		ws.Conn.SendMsg(db.WsMsgModel{Id: db.WsID_Update, Data: parm})
		runtime.MessageDialog(configs.APP_CTX, runtime.MessageDialogOptions{
			Type:          runtime.InfoDialog,
			Title:         "提示",
			Message:       "配置已生效!",
			DefaultButton: "ok",
		})
		return true
	}
	log.Println("save setting err:", err)
	return false
}
func GetSetting() (*db.Setting, error) {
	bt, err := db.Get(db.K_Setting)
	if err == nil {
		out := db.Setting{}
		err := json.Unmarshal(bt, &out)
		if err != nil {
			return initSettingModel(), err
		}
		updateSetting(out)
		return &out, nil
	}
	return initSettingModel(), err
}

func (sef *App) InitSetting() (*db.Setting, error) {
	return GetSetting()
}
func (sef *App) SetReadPanel(width, height int) bool {
	next := sef.TryLock()
	if !next {
		return false
	}
	defer sef.Unlock()
	configs.SettingModel.ReadHeight = height
	configs.SettingModel.ReadWidth = width
	err := db.Set(db.K_Setting, configs.SettingModel)
	return err == nil
}
