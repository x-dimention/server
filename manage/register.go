package manage

import (
	"net/http"

	"github.com/urfave/negroni"
	tool "github.com/zgr126/x-tool"
	"github.com/zgr126/x/server/manage/common"
	hhttp "github.com/zgr126/x/server/manage/hamster/http"
	"github.com/zgr126/x/server/manage/middleware/response"
	"github.com/zgr126/x/server/router"
)

type registerSelfField struct {
	Port   string
	Router http.Handler
	Http   hhttp.HttpServe
}

func (r *registerSelfField) loadHamster() {
	r.loadHamster_Httprouter()
	r.loadHamster_httpPort()
}

func (r *registerSelfField) loadHamster_httpPort() {
	var _config struct {
		HTTP struct {
			Port string
		}
	}

	err := tool.ReadConfigFromFile(&_config)
	if err != nil {
		tool.Console.Panic("HTTP", err)
	}
	r.Http = hhttp.New()
	port := _config.HTTP.Port
	r.Http.SetPort(port)
	common.CV.HttpPort = port
}

func (r *registerSelfField) loadHamster_Httprouter() {
	r.Router = router.GetRouter()
	common.CV.Router = r.Router
}

func (r *registerSelfField) StartServe() {

	n := negroni.Classic() // Includes some default middlewares
	n.Use(response.NewMiddleware())
	n.UseHandler(r.Router)
	r.Http.SetHandler(n)
	r.Http.Run()
}
