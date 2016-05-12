package models

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Member struct {
	Id        int64  `json:"id" xorm:"not null pk autoincr index INT(11)"`
	Telephone string `json:"telephone" xorm:"not null unique pk VARCHAR(11)"`
	Password  string `json:"password" xorm:"not null VARCHAR(255)"`
	Email     string `json:"email" xorm:"not null VARCHAR(255)"`
	Status    int    `json:"status" xorm:"not null INT(11)"`
}

var (
	engine           *xorm.Engine
	Mem_EXSIT        = errors.New("用户已经被注册过了")
	Mem_NOT_EXSIT    = errors.New("用户不存在")
	PASS_OR_NAME_ERR = errors.New("用户或者密码错误")
)

func init() {
	engine, _ = xorm.NewEngine("mysql", "root:root@tcp(120.76.123.47:3306)/maker?charset=utf8")
	engine.ShowSQL(true)
}

func CreateMem(u *Member) (int64, error) {
	id, err := engine.Insert(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateMem(u *Member) (int64, error) {
	count, err := engine.Where("id=?", u.Id).Update(u)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func UpdateSatusMem(id, status int) (int64, error) {
	res, err := engine.Exec("update member set status=? where id=?", status, id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func DeleteMem(id int) (int64, error) {
	count, err := engine.Where("id=?", id).Delete(new(Member))
	if err != nil {
		return 0, err
	}
	return count, nil
}

//分页条件查询
func QueryMem(where map[string]interface{}, start, lenght int) ([]*Member, error) {
	list := []*Member{}
	sql := " 1 AND "
	if len(where) > 0 {
		if where["id"] != "" {
			sql += "id = " + fmt.Sprintf("%d", where["id"]) + " AND "
		}
		if where["telephone"] != "" {
			sql += "telephone = " + fmt.Sprintf("%v", where["telephone"]) + " AND "
		}
		if where["password"] != "" {
			sql += "password = " + fmt.Sprintf("%v", where["password"]) + " AND "
		}
		if where["email"] != "" {
			sql += "email = " + fmt.Sprintf("%v", where["email"]) + " AND "
		}
		if where["status"] != "" {
			sql += "status = " + fmt.Sprintf("%d", where["status"]) + " AND "
		}
	}
	sql += " 1 "
	if err := engine.Where(sql).Limit(lenght, start).Desc("id").Find(&list); err != nil {
		return list, err
	}

	return list, nil
}

func QueryMemByTel(tel string) (*Member, []*Rights, error) {
	m := &Member{}
	has, err := engine.Where("telephone=?", tel).Get(m)
	if err != nil {
		return nil, nil, err
	} else if !has {
		return nil, nil, Mem_NOT_EXSIT
	}

	r := []*Rights{}
	if err := engine.Sql("SELECT id, forwardaddr, controboxaddr, relayaddr, correspondnum, remark, status, type FROM rights").Where("id in (select rights_id from mem_right where member_id=1?)", m.Id).Find(&r); err != nil {

	}
	return m, r, nil
}
