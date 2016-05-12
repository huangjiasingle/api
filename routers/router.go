package routers

import (
	"api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/member", &controllers.MemberController{}, "POST:Save")
	beego.Router("/api/v1/member", &controllers.MemberController{}, "GET:All")
	beego.Router("/api/v1/member", &controllers.MemberController{}, "PUT:Put")
	beego.Router("/api/v1/member", &controllers.MemberController{}, "DELETE:Delete")
	beego.Router("/api/v1/member/auth", &controllers.MemberController{}, "POST:Auth")

	beego.Router("/api/v1/rights", &controllers.RightsController{}, "POST:Save")
	beego.Router("/api/v1/rights", &controllers.RightsController{}, "GET:All")
	beego.Router("/api/v1/rights/count", &controllers.RightsController{}, "GET:Count")
	beego.Router("/api/v1/rights", &controllers.RightsController{}, "PUT:Put")
	beego.Router("/api/v1/rights", &controllers.RightsController{}, "DELETE:Delete")
	beego.Router("/api/v1/open", &controllers.RightsController{}, "GET:Open")
	beego.Router("/api/v1/close", &controllers.RightsController{}, "GET:Close")

	beego.Router("/api/v1/memrights", &controllers.MemberRightsController{}, "POST:Save")
	beego.Router("/api/v1/memrights", &controllers.MemberRightsController{}, "GET:All")
	beego.Router("/api/v1/memrights", &controllers.MemberRightsController{}, "PUT:Put")
	beego.Router("/api/v1/memrights", &controllers.MemberRightsController{}, "DELETE:Delete")
}
