package routers

import (
	"loveHome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{},"get:GetArea")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:GetHouseIndex")
	//api/v1.0/session
	beego.Router("/api/v1.0/session", &controllers.SessionController{},"get:GetSessionData;delete:DeleteSessionData")
	///api/v1.0/users
	beego.Router("/api/v1.0/users", &controllers.UserController{},"post:Reg")


}
