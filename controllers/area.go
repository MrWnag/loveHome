package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
	"time"
	"encoding/json"
	//"fmt"
)

type AreaController struct {
	beego.Controller
}

func(this*AreaController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("connect success")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)

	//从redis缓存中拿数据拿数据
	cache_conn, err := cache.NewCache("redis", `{"key":"lovehome","conn":":6370","dbNum":"0"}`)

	if areaData := cache_conn.Get("area");areaData!=nil{
		beego.Info("get data from cache===========")
		resp["data"] = areaData
		return
	}



	//beego.Info("cache_conn.aa =",cache_conn.Get("aaa"))
	//fmt.Printf("cache_conn ,conn[aa]= %s\n",cache_conn.Get("aaa"))

	//从mysql数据库拿到area数据
	var areas  []models.Area

	o := orm.NewOrm()
	num ,err :=o.QueryTable("area").All(&areas)

	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
 	}
 	if num ==0{
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}

	resp["data"] = areas

	//把数据转换成json格式存入缓存
	json_str,err := json.Marshal(areas)
	if err != nil{
		beego.Info("encoding err")
		return
	}

	cache_conn.Put("area",json_str,time.Second*3600)


		//打包成json返回给前段
	beego.Info("query data sucess ,resp =",resp,"num =",num)


}
