# jsonwrap

This library help Marshal you struct to json string witj c like comments

Usage.
Download library:

go get github.com/denisbdn/jsonwrap

In code:


import (

	"fmt"

	"time"

	"github.com/denisbdn/jsonwrap"
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

	All    []Simple        `json:"all" jscmm:"it's all"

	KV     map[int]*Simple `json:"kv" jscmm:"it's all"`

	Inter  interface{}     `json:"interface" jscmm:"it's call bee anything"`

	Time   time.Time       `json:"time" jscmm:"it's call time field"`

}


func main() {

	multyLine := jsonwrap.New()

	if arr, err := multyLine.Marshal(ComplexSimple{}); err == nil {

		fmt.Println(string(arr[:]))

	}


	singleLine := jsonwrap.New()

	singleLine.NewLine = ""

	singleLine.NewField = "  "

	if arr, err := singleLine.Marshal(ComplexSimple{}); err == nil {

		fmt.Println(string(arr[:]))

	}


	complex := jsonwrap.New()

	if arr, err := complex.Marshal(ComplexSimple{}); err == nil {

		fmt.Println(string(arr[:]))

	}

}

Cout:
