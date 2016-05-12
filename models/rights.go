package models

import (
	"fmt"
)

type Rights struct {
	Id            int64  `json:"id" xorm:"not null pk autoincr INT(11)"`
	Forwardaddr   byte   `json:"forwardaddr" xorm:"not null INT(11)"`
	Controboxaddr byte   `json:"controboxaddr" xorm:"not null INT(11)"`
	Relayaddr     byte   `json:"relayaddr" xorm:"not null INT(11)"`
	Correspondnum string `json:"correspondnum" xorm:"not null VARCHAR(255)"`
	Remark        string `json:"remark" xorm:"VARCHAR(512)"`
	Status        int64  `json:"status" xorm:"default 0 INT(2)"`
	Type          int64  `json:"type" xorm:"not null default 0 INT(11)"`
}

func CreateRights(u *Rights) (int64, error) {
	id, err := engine.Insert(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateRights(u *Rights) (int64, error) {
	count, err := engine.Where("id=?", u.Id).Update(u)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func UpdateSatusRights(id int64, status int) (int64, error) {
	res, err := engine.Exec("update rights set status=? where id=?", status, id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func DeleteRights(id int) (int64, error) {
	count, err := engine.Where("id=?", id).Delete(new(Rights))
	if err != nil {
		return 0, err
	}
	return count, nil
}

//分页条件查询
func QueryRights(where map[string]interface{}, start, lenght int) ([]*Rights, error) {
	fmt.Sprintf("%v", 1)
	list := []*Rights{}
	sql := " 1 AND "
	if len(where) > 0 {
		if where["id"] != nil {
			sql += "id = " + fmt.Sprintf("%d", where["id"]) + " AND "
		}
		if where["forwardaddr"] != nil {
			sql += "forwardaddr = " + fmt.Sprintf("%v", where["forwardaddr"]) + " AND "
		}
		if where["Controboxaddr"] != nil {
			sql += "controboxaddr = " + fmt.Sprintf("%v", where["controboxaddr"]) + " AND "
		}
		if where["relayaddr"] != nil {
			sql += "relayaddr = " + fmt.Sprintf("%v", where["relayaddr"]) + " AND "
		}
		if where["correspondnum"] != nil {
			sql += "correspondnum like '%" + fmt.Sprintf("%v", where["correspondnum"]) + "%" + "' AND "
		}
		if where["remark"] != nil {
			sql += "remark = " + fmt.Sprintf("%v", where["remark"]) + " AND "
		}
		if where["type"] != nil {
			sql += "type = " + fmt.Sprintf("%v", where["type"]) + " AND "
		}
		if where["status"] != nil {
			sql += "status = " + fmt.Sprintf("%d", where["status"]) + " AND "
		}
	}
	sql += " 1 "
	if err := engine.Where(sql).Limit(lenght, start).Asc("correspondnum").Find(&list); err != nil {
		return list, err
	}
	return list, nil
}

func QueryByCorrespondNum(correspondNum, types string) (*Rights, error) {
	r := &Rights{}
	has, err := engine.Where("correspondNum=? and type = ?", correspondNum, types).Get(r)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return r, nil
}

func Count(status, types string) (int64, error) {
	r := &Rights{}
	count, err := engine.Where("status=? and type=?", status, types).Count(r)
	if err != nil {
		return 0, err
	}
	return count, nil
}
