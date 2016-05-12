package models

import (
	"fmt"
)

type MemRight struct {
	Id       int64 `json:"id" xorm:"not null pk autoincr INT(11)"`
	MemberId int64 `json:"memberId" xorm:"not null index INT(11)"`
	RightsId int64 `json:"rightsId" xorm:"not null index INT(11)"`
}

func CreateMemRights(u *MemRight) (int64, error) {
	id, err := engine.Insert(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateMemRights(u *MemRight) (int64, error) {
	count, err := engine.Where("id=?", u.Id).Update(u)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func DeleteMemRights(id int) (int64, error) {
	count, err := engine.Where("id=?", id).Delete(new(MemRight))
	if err != nil {
		return 0, err
	}
	return count, nil
}

//分页条件查询
func QueryMemRights(where map[string]interface{}, start, lenght int) ([]*MemRight, error) {
	fmt.Sprintf("%v", 1)
	list := []*MemRight{}
	sql := " 1 AND "
	if len(where) > 0 {
		if where["id"] != nil {
			sql += "id = " + fmt.Sprintf("%d", where["id"]) + " AND "
		}
		if where["memberId"] != nil {
			sql += "memberId = " + fmt.Sprintf("%v", where["memberId"]) + " AND "
		}
		if where["rightsId"] != nil {
			sql += "rightsId = " + fmt.Sprintf("%d", where["rightsId"]) + " AND "
		}
	}
	sql += " 1 "
	if err := engine.Where(sql).Limit(lenght, start).Desc("id").Find(&list); err != nil {
		return list, err
	}
	return list, nil
}
