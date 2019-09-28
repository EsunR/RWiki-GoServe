package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:SysController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:SysController"],
        beego.ControllerComments{
            Method: "SysInit",
            Router: `/sysInit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:SysController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:SysController"],
        beego.ControllerComments{
            Method: "SysStatus",
            Router: `/sysStatus`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:TestController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:TestController"],
        beego.ControllerComments{
            Method: "TestRouter",
            Router: `/testRouter`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:TestController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:TestController"],
        beego.ControllerComments{
            Method: "TestUsers",
            Router: `/testUsers`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
