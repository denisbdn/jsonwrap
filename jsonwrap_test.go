package jsonwrap

import (
	"fmt"
	"testing"
	"time"
)

type Simple struct {
	First  int    `json:"first" jscmm:"it's First field"`
	Second uint   `json:"second" jscmm:"it's Second field"`
	Third  string `json:"third" jscmm:"it's Third field"`
	Fourth bool   `json:"fourth" jscmm:"it's Fourth field"`
}

type ComplexSimple struct {
	Id     int             `json:"id" jscmm:"it's Id field"`
	Caller Simple          `json:"caleer" jscmm:"it's Caller"`
	Called *Simple         `json:"called" jscmm:"it's Called"`
	All    []Simple        `json:"all" jscmm:"it's all"`
	KV     map[int]*Simple `json:"kv" jscmm:"it's all"`
	Inter  interface{}     `json:"interface" jscmm:"it's call bee anything"`
	Time   time.Time       `json:"time" jscmm:"it's call time field"`
}

func TestJsonWraper(t *testing.T) {
	jsonWraper := New()
	// jsonWraper.NewLine = ""
	// jsonWraper.NewField = "  "
	arr1, err1 := jsonWraper.Marshal(Simple{})
	if err1 != nil {
		t.Error(err1)
	}
	str1 := string(arr1[:])
	fmt.Println(str1)
	t.Log(str1)

	arr2, err2 := jsonWraper.Marshal(ComplexSimple{})
	if err2 != nil {
		t.Error(err2)
	}
	str2 := string(arr2[:])
	fmt.Println(str2)
	t.Log(str2)
}
