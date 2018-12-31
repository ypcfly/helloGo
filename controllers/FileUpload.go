package controllers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type FileController struct {
	beego.Controller
}

// struct接收文件
type Images struct {
	File string
}

func (f *FileController) URLMapping() {
	f.Mapping("Upload",f.Upload)
	f.Mapping("DoUpload",f.DoUpload)
	f.Mapping("Base64",f.Base64)
}

// @router /file/upload [get]
func (f *FileController) Upload() {
	logs.Info(">>>> forward to file upload page start <<<<")

	f.TplName = "file/file.html"
}

// @router /file/doUpload [post]
func (f *FileController) DoUpload() {
	logs.Info(">>>> upload file start <<<<")
	file,header,_ := f.GetFile("file")

	defer file.Close()
	err := f.SaveToFile("file","static/img/" + header.Filename)
	if err != nil {
		log.Fatal(err)
		f.Ctx.Redirect(302,"/false")
	} else {
		f.Ctx.Redirect(302,"/success")
	}
}

// @router /file/base64 [post]
func (f *FileController) Base64() {
	logs.Info(">>>> receive base64 file start <<<<")

	/*直接从requestbody读取二进制，并转成string*/
	bys,_ := ioutil.ReadAll(f.Ctx.Request.Body)

	/*将requestBody数据转成一个对象*/
	var files = Images{}
	er := json.Unmarshal(bys,&files)
	if er != nil {
		log.Fatal(er)
	}
	fs := files.File


	//base := f.Ctx.Request.Form.Get("file")
	//strs := strings.Split(base,",")
	strs := strings.Split(fs,",")
	/*base64字符串的头和体*/
	head := strs[0]
	body := strs[1]

	s := strings.LastIndex(head,"/")
	e := strings.LastIndex(head,";")
	name := string([]byte(head))[s+1:e]
	bytes,err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		log.Fatal(err)
		logs.Error(">>>> decode base64 string to bytes failed <<<<")
	}

	err = ioutil.WriteFile("./static/upload/" + strconv.Itoa(time.Now().Nanosecond()) + "." + name,bytes,0666)
	if err != nil {
		log.Fatal(err)
		logs.Error(">>>> write to file failed <<<<")
	}
	user := User{UserName:"hello",Age:22,Address:"US",Mobile:"158725544"}
	result := Result{Code:"200",Msg:"success",Rs:&user}
	f.Data["json"] = &result
	f.ServeJSON()
}

type Result struct {
	Code string `json:"code"`
	Msg string	`json:"msg"`
	Rs *User	`json:"rs"`
}

type User struct {
	UserName string `json:"userName"`
	Age int			`json:"age"`	
	Mobile string	`json:"mobile"`
	Address string	`json:"address"`
}
