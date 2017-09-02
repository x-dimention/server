package main

import (
	"log"

	"github.com/zgr126/x-tool"
	"github.com/zgr126/x-tool/odm"
	"github.com/zgr126/x/server/model"
)

func main() {
	var _config struct {
		DB odm.ConnectConfig
	}

	err := tool.ReadConfigFromFile(&_config)
	if err != nil {
		tool.Console.Panic("DB", err)
	}

	err = model.SetDefaultDB(_config.DB)
	if err != nil {
		tool.Console.Panic("DB", err)
	}

	//////////、
	o := model.New()

	type ddd struct {
		BB int `json:"bb,int,omitempty"`
	}
	result := []ddd{}
	_, err = o.Query(nil, "FOR d IN page LIMIT 10 RETURN d", nil, &result)
	log.Printf("len:%v, Addr:%p", len(result), &result)
	for i, v := range result {
		log.Println(i, v)
	}
}
