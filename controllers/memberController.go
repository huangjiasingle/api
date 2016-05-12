package controllers

import (
	"api/models"
	"encoding/json"
)

type MemberController struct {
	APIController
}

func (this *MemberController) Save() {
	m := &models.Member{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	}

	mp, _, err := models.QueryMemByTel(m.Telephone)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	}

	if mp != nil {
		this.err = models.Mem_EXSIT
		this.data = map[string]interface{}{"id": nil}
	} else {
		id, err := models.CreateMem(m)
		this.version = 1.0
		if err != nil {
			this.err = err
			this.data = map[string]interface{}{"id": nil}
		} else {
			this.err = nil
			this.data = map[string]interface{}{"id": id}
		}
	}

	this.Finish()
}

func (this *MemberController) Put() {
	m := &models.Member{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"affectrows": nil}
	}
	count, err := models.UpdateMem(m)
	this.version = 1.0
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"affectrows": nil}
	} else {
		this.err = nil
		this.data = map[string]interface{}{"affectrows": count}
	}

	this.Finish()
}

func (this *MemberController) Delete() {
	m := map[string]int{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), &m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	}
	count, err := models.DeleteMem(m["id"])
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	} else {
		this.err = nil
		this.data = map[string]interface{}{"affectrows": count}
	}

	this.Finish()
}

func (this *MemberController) All() {
	where := map[string]interface{}{}
	start, _ := this.GetInt("start")
	lenght, _ := this.GetInt("lenght")
	list, err := models.QueryMem(where, start, lenght)
	this.version = 1.0
	if err != nil {
		this.err = err
		this.data = []map[string]interface{}{}
	} else {
		this.err = nil
		this.data = list
	}
	this.Finish()
}

//登陆认证
func (this *MemberController) Auth() {
	ControllerLog.Info(this.GetString("telephone"))
	ControllerLog.Info(this.GetString("password"))
	telephone := this.GetString("telephone")
	password := this.GetString("password")
	this.version = 1.0
	m, r, err := models.QueryMemByTel(telephone)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"success": false}
	}
	if m != nil {
		if m.Password == password {
			this.Ctx.SetCookie("user", telephone)
			this.SetSession("username", telephone)
			this.SetSession(telephone, r)
			this.data = map[string]interface{}{"success": true, "rights": r}
		} else {
			this.err = models.PASS_OR_NAME_ERR
			this.data = map[string]interface{}{"success": false}
		}
	} else {
		this.err = models.Mem_NOT_EXSIT
		this.data = map[string]interface{}{"success": false}
	}

	this.Finish()
}
