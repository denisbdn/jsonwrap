# jsonwrap

This library help Marshal you struct to json string with c-like comments

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

	{
		"id": 0, // int, it's Id field 
		"caleer": {
			"first": 0, // int, it's First field 
			"second": 0, // uint, it's Second field 
			"third": "", // string, it's Third field 
			"fourth": false // bool, it's Fourth field 
		}, // struct, it's Caller 
		"called": {
			"first": 0, // int, it's First field 
			"second": 0, // uint, it's Second field 
			"third": "", // string, it's Third field 
			"fourth": false // bool, it's Fourth field 
		}, // ptr, it's Called 
		"all": [
			{
				"first": 0, // int, it's First field 
				"second": 0, // uint, it's Second field 
				"third": "", // string, it's Third field 
				"fourth": false // bool, it's Fourth field 
			}
		], // slice, it's all 
		"kv": {
			"intKeyMap": {
				"first": 0, // int, it's First field 
				"second": 0, // uint, it's Second field 
				"third": "", // string, it's Third field 
				"fourth": false // bool, it's Fourth field 
			}
		}, // map, it's all 
		"interface": interface{} // interface, it's call bee anything 
	}
	{  "id": 0, /* int, it's Id field */  "caleer": {  "first": 0, /* int, it's First field */  "second": 0, /* uint, it's Second field */  "third": "", /* string, it's Third field */  "fourth": false /* bool, it's Fourth field */  }, /* struct, it's Caller */  "called": {  "first": 0, /* int, it's First field */  "second": 0, /* uint, it's Second field */  "third": "", /* string, it's Third field */  "fourth": false /* bool, it's Fourth field */  }, /* ptr, it's Called */  "all": [  {  "first": 0, /* int, it's First field */  "second": 0, /* uint, it's Second field */  "third": "", /* string, it's Third field */  "fourth": false /* bool, it's Fourth field */  }  ], /* slice, it's all */  "kv": {  "intKeyMap": {  "first": 0, /* int, it's First field */  "second": 0, /* uint, it's Second field */  "third": "", /* string, it's Third field */  "fourth": false /* bool, it's Fourth field */  }  }, /* map, it's all */  "interface": interface{} /* interface, it's call bee anything */  }
	{
		"id": 0, // int, it's Id field 
		"caleer": {
			"first": 0, // int, it's First field 
			"second": 0, // uint, it's Second field 
			"third": "", // string, it's Third field 
			"fourth": false // bool, it's Fourth field 
		}, // struct, it's Caller 
		"called": {
			"first": 0, // int, it's First field 
			"second": 0, // uint, it's Second field 
			"third": "", // string, it's Third field 
			"fourth": false // bool, it's Fourth field 
		}, // ptr, it's Called 
		"all": [
			{
				"first": 0, // int, it's First field 
				"second": 0, // uint, it's Second field 
				"third": "", // string, it's Third field 
				"fourth": false // bool, it's Fourth field 
			}
		], // slice, it's all 
		"kv": {
			"intKeyMap": {
				"first": 0, // int, it's First field 
				"second": 0, // uint, it's Second field 
				"third": "", // string, it's Third field 
				"fourth": false // bool, it's Fourth field 
			}
		}, // map, it's all 
		"interface": interface{} // interface, it's call bee anything 
	}
	
Additionaly you can set your own format function insted inner
	
	multyLine.Format = func(wraper *jsonwrap.JsonWraper, kind reflect.Kind, name string, comment string) string {
		return "..."
	}
