package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"
	"net"

	"api/utils"
	"github.com/astaxie/beego"
)

type RightsController struct {
	APIController
}

func (this *RightsController) Save() {
	m := &models.Rights{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"affectrows": nil}
	}
	id, err := models.CreateRights(m)
	this.version = 1.0
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"affectrows": nil}
	} else {
		this.err = nil
		this.data = map[string]interface{}{"affectrows": id}
	}

	this.Finish()
}

func (this *RightsController) Put() {
	m := &models.Rights{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"affectrows": nil}
	}
	count, err := models.UpdateRights(m)
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

func (this *RightsController) Delete() {
	ControllerLog.Info(this.GetString("id"))
	m := map[string]int{}
	err := json.Unmarshal((this.Ctx.Input.RequestBody), &m)
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	}
	count, err := models.DeleteRights(m["id"])
	if err != nil {
		this.err = err
		this.data = map[string]interface{}{"id": nil}
	} else {
		this.err = nil
		this.data = map[string]interface{}{"affectrows": count}
	}

	this.Finish()
}

func (this *RightsController) All() {
	ControllerLog.Info("rights is runuing")
	where := map[string]interface{}{}
	correspondnum := this.GetString("correspondnum")
	if correspondnum != "" {
		where["correspondnum"] = correspondnum
	}
	status := this.GetString("status")
	if status != "" {
		where["status"] = status
	}
	tp := this.GetString("type")
	if tp != "" {
		where["type"] = tp
	}
	start, _ := this.GetInt("start")
	lenght, _ := this.GetInt("lenght")

	list, err := models.QueryRights(where, start, lenght)
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

func (this *RightsController) Open() {
	ControllerLog.Info("open is running")
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", beego.AppConfig.String("tcp_server"))

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		ControllerLog.Debug(err.Error())
		this.Data["json"] = map[string]interface{}{"success": false}
	}
	conn.SetKeepAlive(true)
	defer conn.Close()

	ControllerLog.Debug("connected!")

	ControllerLog.Info(this.GetString("correspondnum"))
	correspondnum := this.GetString("correspondnum")
	tp := this.GetString("type")
	v, err := models.QueryByCorrespondNum(correspondnum, tp)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = map[string]interface{}{"success": false}
	}
	if v != nil {
		ControllerLog.Info("v is %v ,%v,%v", v.Forwardaddr, v.Controboxaddr, v.Relayaddr)
		bt := utils.GenerateCmd(v.Forwardaddr, v.Controboxaddr, v.Relayaddr, true)
		ControllerLog.Info("is %v", bt)
		_, err1 := conn.Write(bt)

		if err1 != nil {
			ControllerLog.Debug(err.Error())
			this.Data["json"] = map[string]interface{}{"success": false}
		}

		err = onMessageRecived(conn)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"success": false}
		}
		this.Data["json"] = map[string]interface{}{"success": true}
		if _, err := models.UpdateSatusRights(v.Id, 1); err != nil {
			this.Data["json"] = map[string]interface{}{"success": false}
		}
	} else {
		this.Data["json"] = map[string]interface{}{"success": false}
	}
	this.ServeJSON()
}

func (this *RightsController) Close() {
	ControllerLog.Info("close is running")
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", beego.AppConfig.String("tcp_server"))

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		ControllerLog.Debug(err.Error())
	}

	conn.SetKeepAlive(true)
	defer conn.Close()
	ControllerLog.Debug("connected!")

	correspondnum := this.GetString("correspondnum")
	tp := this.GetString("type")
	v, err := models.QueryByCorrespondNum(correspondnum, tp)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = map[string]interface{}{"success": false}
	}
	ControllerLog.Info("v is %v", v)

	if v != nil {
		bt := utils.GenerateCmd(v.Forwardaddr, v.Controboxaddr, v.Relayaddr, false)
		_, err1 := conn.Write(bt)

		if err1 != nil {
			ControllerLog.Debug(err.Error())
		}

		err = onMessageRecived(conn)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"success": false}
		}
		this.Data["json"] = map[string]interface{}{"success": true}
		if _, err := models.UpdateSatusRights(v.Id, 0); err != nil {
			this.Data["json"] = map[string]interface{}{"success": false}
		}
	} else {
		this.Data["json"] = map[string]interface{}{"success": false}
	}

	this.ServeJSON()
}

var quitSemaphore chan bool

func onMessageRecived(conn *net.TCPConn) error {
	b := make([]byte, 12)
	_, err := conn.Read(b)
	if err != nil {
		fmt.Println("read:", err)
		return err
	}
	ControllerLog.Info("read byte is %v", b)
	return nil
}

func (this *RightsController) Count() {
	status := this.GetString("status")
	types := this.GetString("type")
	count, err := models.Count(status, types)
	this.version = 1.0
	if err != nil {
		this.err = err
		this.Data["json"] = map[string]interface{}{"count": 0}
	} else {
		this.Data["json"] = map[string]interface{}{"count": count}
	}

	this.Finish()
}
