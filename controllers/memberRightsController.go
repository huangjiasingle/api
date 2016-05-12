package controllers

import (
	"api/models"
	"encoding/json"
)

type MemberRightsController struct {
	APIController
}

func (this *MemberRightsController) Save() {
	m := &models.MemRight{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	}
	id, err := models.CreateMemRights(m)
	this.version = 1.0
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	} else {
		this.err = nil
		this.data = map[string]interface{}{"id": id}
	}

	this.Finish()
}

func (this *MemberRightsController) Put() {
	m := &models.MemRight{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"affectrows": nil}
	}
	count, err := models.UpdateMemRights(m)
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

func (this *MemberRightsController) Delete() {
	m := map[string]int{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), &m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	}
	count, err := models.DeleteMemRights(m["id"])
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	} else {
		this.err = nil
		this.data = map[string]interface{}{"affectrows": count}
	}

	this.Finish()
}

func (this *MemberRightsController) All() {
	where := map[string]interface{}{}
	start, _ := this.GetInt("start")
	lenght, _ := this.GetInt("lenght")
	list, err := models.QueryMemRights(where, start, lenght)
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
