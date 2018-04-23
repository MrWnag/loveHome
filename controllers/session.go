package controllers
import (
	"github.com/astaxie/beego"
	"loveHome/models"

	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type SessionController struct {
	beego.Controller
}

func(this*SessionController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this*SessionController)GetSessionData(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	user := models.User{}
	//user.Name = "wyj"
//
	//resp["errno"] = 0
	//resp["errmsg"] = "OK"


	resp["errno"] =models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] =models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user

	}

}

func (this*SessionController)DeleteSessionData(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")

	resp["errno"] =models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)


}


func (this*SessionController)Login(){

//1.得到用户信息
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)

	//beego.Info("======name = ",resp["mobile"],"=======password =",resp["password"])


//2.判断是否合法
	if resp["mobile"] == nil || resp["password"] == nil{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)

		//beego.Info("111111name = ",resp["mobile"],"=======password =",resp["password"])
		return
	}


//3.与数据库匹配判断账号密码正确

	o := orm.NewOrm()
	user := models.User{Name:resp["mobile"].(string)}

	qs := o.QueryTable("user")
	err := qs.Filter("mobile", "7777").One(&user)

	if err != nil{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		//beego.Info("222222=======errr =",err)
		return
	}

	if user.Password_hash != resp["password"]{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		beego.Info("333333name = ",resp["mobile"],"=======password =",resp["password"])
		return
	}




	//4.添加session
	this.SetSession("name",resp["mobile"])
	this.SetSession("mobile",resp["mobile"])
	this.SetSession("user_id",user.Id)



//5.返回json数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}