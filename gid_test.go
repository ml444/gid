package gid

import (
	"fmt"
	"github.com/ml444/gid/server"
	"strconv"
	"testing"
	"time"
)

var idInst IdServer = NewIdService(1, 2, 0, 0, "sync")
var id uint64
var testId uint64 = 2307230695474331649
var mid uint64 = 2307230695474331649
var idObj = &server.Id{
	Machine: 1,
	Seq:     0,
	Time:    1323400746,
	Method:  2,
	Type:    0,
	Version: 0,
}
var tf string
var ts int64

func TestIdService_All(t *testing.T) {
	//idsvc := NewIdService(1, 2, 1, 0, "sync")
	//var idInst IdServer = idsvc // 引用传递而不是值传递
	for i := 0; i < 2; i++ {
		id = idInst.GenId()
		fmt.Println(id)
		idObj = idInst.ExpId(id)
		fmt.Printf("%#v", *idObj)
		ts, tf = idInst.TransTime(idObj.GetTime())
		fmt.Println(tf, ts)
		// 延时
		time.Sleep(300 * time.Millisecond)
	}
	mid = idInst.MakeId(1, 0, 1610780776324, 2, 0, 0)
	fmt.Println(id)
}

func TestIdService_GenId(t *testing.T) {
	id = idInst.GenId()
	if len(strconv.FormatUint(id, 10)) < 19 {
		t.Error("Error: len(id) < 19!!!")
	}
}

func TestIdService_ExpId(t *testing.T) {
	idObj = idInst.ExpId(id)
	fmt.Printf("%#v", *idObj)
	if idObj.Time < 0 {

	}

	id := idInst.MakeId(1, 0, 1610780776324, 2, 0, 0)
	fmt.Println(id)
}

func TestIdService_TransTime(t *testing.T) {
	for i := 0; i < 2; i++ {
		id := idInst.GenId()
		fmt.Println(id)
		idObj := idInst.ExpId(id)
		fmt.Printf("%#v", *idObj)
		tsf, ts := idInst.TransTime(idObj.GetTime())
		fmt.Println(tsf, ts)
		// 延时
		time.Sleep(300 * time.Millisecond)
	}
	id := idInst.MakeId(1, 0, 1610780776324, 2, 0, 0)
	fmt.Println(id)
}

func TestIdService_MakeId(t *testing.T) {
	for i := 0; i < 2; i++ {
		id := idInst.GenId()
		fmt.Println(id)
		idObj := idInst.ExpId(id)
		fmt.Printf("%#v", *idObj)
		tsf, ts := idInst.TransTime(idObj.GetTime())
		fmt.Println(tsf, ts)
		// 延时
		time.Sleep(300 * time.Millisecond)
	}
	id := idInst.MakeId(1, 0, 1610780776324, 2, 0, 0)
	fmt.Println(id)
}
