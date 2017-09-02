package manage

import (
	tool "github.com/zgr126/x-tool"
	"github.com/zgr126/x/server/model"
)

func detection() {
	detectionDatabase()
}

func detectionDatabase() {
	var _Config struct {
		DefaultDB model.ConfigStruct
	}

	err := tool.ReadConfigFromFile(&_Config)
	if err != nil {
		tool.Console.Panic("DB", err)
	}

	err = model.Init(_Config.DefaultDB)
	if err != nil {
		tool.Console.Panic("DB", err)
	}
}
