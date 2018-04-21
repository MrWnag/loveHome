package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)


type UserController struct {
	beego.Controller
}

func(this*UserController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this*UserController)Reg(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)

//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)
/*
	mobile:"111"
	password:"111"
	sms_code:"111"

	beego.Info(`resp["mobile"] =`,resp["mobile"])
	beego.Info(`resp["password"] =`,resp["password"])
	beego.Info(`resp["sms_code"] =`,resp["sms_code"])
*/

//插入数据库
o:=orm.NewOrm()
user := models.User{}
user.Password_hash = resp["password"].(string)
user.Name = resp["mobile"].(string)
user.Mobile = resp["mobile"].(string)

id,err:=o.Insert(&user)
if err != nil{
	resp["errno"] = 4002
	resp["errmsg"] = "注册失败"
	return
}

beego.Info("reg success ,id = ",id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	this.SetSession("name",user.Name)




}