package models

import (
	"github.com/weilaihui/fdfs_client"
	"github.com/astaxie/beego"
)


func TestUploadByFilename(fileName string)(groupName string,fileId string,err error){


	fdfsClient ,errClient := fdfs_client.NewFdfsClient("conf/client.conf")
	if errClient != nil {
		beego.Info("New FdfsClient error %s", errClient.Error())
		return "","",errClient
	}
	uploadResponse, errUpload :=fdfsClient.UploadByFilename(fileName)
	if errUpload != nil {
		beego.Info("New FdfsClient error %s", errUpload.Error())
		return "","",errUpload
	}
	beego.Info("=================groupNmae = ",uploadResponse.GroupName)
	beego.Info("=================fileId = ",uploadResponse.RemoteFileId)
	return uploadResponse.GroupName,uploadResponse.RemoteFileId,nil
}