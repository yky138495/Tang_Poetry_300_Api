// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"tangshi300/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/tangshi",
			beego.NSInclude(
				&controllers.Tangshi300Controller{},
			),
		),
		// beego.NSNamespace("/main",
		// 	beego.NSInclude(
		// 		&controllers.MainController{},
		// 	),
		// ),
	)
	beego.AddNamespace(ns)
}
