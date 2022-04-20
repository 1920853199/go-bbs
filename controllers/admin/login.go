package admin

import (
	"go-bbs/models/admin"
	"go-bbs/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (ctl *LoginController) Sign() {
	o := orm.NewOrm()
	var setting admin.Setting
	o.QueryTable(new(admin.Setting)).Filter("name", "title").One(&setting)
	ctl.Data["title"] = setting.Value
	ctl.TplName = "admin/login.html"
}
func (ctl *LoginController) Login() {

	username := ctl.GetString("username")
	password := ctl.GetString("password")

	password = utils.PasswordMD5(password, username)

	response := make(map[string]interface{})

	if user, ok := admin.Login(username, password); ok {
		ctl.SetSession("User", *user)
		response["code"] = 200
		response["msg"] = "登录成功！"
	} else {
		response["code"] = 500
		response["msg"] = "登录失败！"
	}

	ctl.Data["json"] = response
	ctl.ServeJSON()

}

func (ctl *LoginController) Logout() {
	ctl.DelSession("User")
	ctl.Redirect("/admin", 302)
}
