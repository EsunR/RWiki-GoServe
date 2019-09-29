package controllers

import (
	"RWiki-GoServe/filters"
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProjectController struct {
	beego.Controller
}

// @router /createProject [post]
func (c *ProjectController) CreateProject() {
	var resp _struct.Resp

	uid := int(filters.TokenData["uid"].(float64))
	projectName := filters.BodyData["projectName"].(string)
	desc := filters.BodyData["desc"].(string)

	o := orm.NewOrm()
	var user models.Users
	user.Id = uid
	_ = o.Read(&user)

	var project models.Projects
	project.ProjectName = projectName
	project.Desc = desc
	project.Creator = &user
	_, _ = o.Insert(&project)

	m2m := o.QueryM2M(&project, "Contributors")
	_, _ = m2m.Add(&user)

	resp.Msg = "ok"
	resp.Data = project
	_ = c.Ctx.Output.JSON(resp, false, false)
}

// @router /getProjectListByUid [get]
func (c *ProjectController) GetProjectListByUid() {
	var resp _struct.Resp
	uid := int(filters.TokenData["uid"].(float64))

	o := orm.NewOrm()
	qs := o.QueryTable("projects")
	var projects []models.Projects
	_, err := qs.Filter("creator_id", uid).RelatedSel().All(&projects)
	if err != nil {
		resp.Msg = "未找到相关项目"
	} else {
		// 封装数据
		var data []interface{}
		for i := range projects {
			item := map[string]interface{}{
				"_id": projects[i].Id,
				"projectInfo": map[string]interface{}{
					"creator": map[string]interface{}{
						"uid":  projects[i].Creator.Id,
						"name": projects[i].Creator.Name,
					},
					"desc":        projects[i].Desc,
					"cover":       projects[i].Cover,
					"createTime":  projects[i].CreateTime,
					"updateTime":  projects[i].UpdateTime,
					"projectName": projects[i].ProjectName,
				},
			}
			data = append(data, item)
		}

		resp.Msg = "ok"
		resp.Data = data
	}
	_ = c.Ctx.Output.JSON(resp, false, false)
}
