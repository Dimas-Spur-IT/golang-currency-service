package routers

import (
	"golang-currency-service/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/rates", &controllers.RateController{}, "get:GetAll")
	web.Router("/rate", &controllers.RateController{}, "get:GetByDate")
}
