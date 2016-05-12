package models

import (
	"fmt"
	"testing"
	"time"
)

// func TestCreateMem(t *testing.T) {
// 	m := &Member{Telephone: "13452254487", Password: "123456", Email: "123@qq.com", Status: 1}
// 	id, err := CreateMem(m)
// 	if err != nil {
// 		t.Log(err.Error())
// 	}
// 	t.Log(id)
// 	// DeleteMem(id)
// }

// func TestDelete(t *testing.T) {
// 	num, err := Delete(25)
// 	if err != nil {
// 		t.Log(err)
// 	}
// 	t.Log(num)
// }

func TestQuery(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go QueryMem(nil, 0, 10)
		// list, err := QueryMem(nil, 0, 10)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		fmt.Println(i)
	}

	time.Sleep(time.Second * 600)
}
