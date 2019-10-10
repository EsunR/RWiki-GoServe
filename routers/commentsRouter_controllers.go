package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: `/getUserInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:BaseController"],
        beego.ControllerComments{
            Method: "UpdateToken",
            Router: `/updateToken`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:ProjectController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateProject",
            Router: `/createProject`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RWiki-GoServe/controllers:ProjectController"] = append(beego.GlobalControllerRouter["RWiki-GoServe/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectListByUid",
            Router: `/getProjectListByUid`,
            AllowHTTPMethods: []string{"get"},
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
            Method: "TestData",
            Router: `/testData`,
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
