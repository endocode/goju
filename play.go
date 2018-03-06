package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// Traverse is the object collection all data on a traversal
type Traverse struct {
	check *Check
}

func cutString(i interface{}, l int) string {
	var out string
	if i == nil {
		out = "nada"
	} else {
		out = fmt.Sprintf("%s", i)
	}
	if len(out) > l {
		return out[0:l] + " ..."
	}
	return out
}

func (t *Traverse) applyRule(offset string, treeValue reflect.Value,
	rulesValue reflect.Value, rules interface{}) {
	if rulesValue.Kind() == reflect.Map {
		m, ok := rules.(map[string]interface{})
		tv := treeValue.Interface()
		if ok {
			for k, v := range m {
				capMethod := strings.Title(k)

				method := reflect.ValueOf(t.check).MethodByName(capMethod)
				if method.IsValid() {
					fmt.Printf("%s\t rules %q (%q, %q)\n", offset, capMethod, v, cutString(tv, 40))
					method.Call([]reflect.Value{reflect.ValueOf(v), reflect.ValueOf(tv)})
				} else {
					switch treeValue.Kind() {
					case reflect.String, reflect.Float64, reflect.Bool:
						{
							t.check.AddError("unknown method %q requested with args(%q, %q)", capMethod, v, cutString(tv, 40))
						}
					}
				}
			}
		} else {
			t.check.AddError("found unknown rule %v with value %q", rules, rules)
		}
	}
	//	fmt.Printf("# errors %d %d\n", t.falseCounter, t.trueCounter)
}

func (t *Traverse) traverse(offset string, tree interface{}, rules interface{}) {
	if tree == nil || rules == nil {
		//		fmt.Printf(offset+"< traverse t is nil=%t r is nil=%t>\n", tree == nil, rules == nil)
		return
	}
	treeValue := reflect.ValueOf(tree)
	rulesValue := reflect.ValueOf(rules)
	fmt.Printf("%s< traverse %v\n", offset, treeValue.Type())

	switch treeValue.Kind() {

	case reflect.Slice, reflect.Array:
		t.applyRule(offset, treeValue,
			rulesValue, rules)

		for i, vi := range tree.([]interface{}) {
			index := fmt.Sprintf("%d:", i)
			index = ""
			t.traverse(offset+index+"\t", vi, rules)
		}

	case reflect.Map:
		for k, v := range tree.(map[string]interface{}) {
			r, ok := rulesValue.Interface().(map[string]interface{})
			if ok {
				// fmt.Printf("### ok key %q %v =: %q \n", k, cutString(v, 30), cutString(r[k], 30))
				t.traverse(offset+"\t ", v, r[k])
			} else {
				// fmt.Printf("#### not ok")
				t.applyRule(offset, treeValue, rulesValue, r)
			}
		}

	case reflect.String, reflect.Float64, reflect.Bool:
		t.applyRule(offset, treeValue, rulesValue, rules)
	default:
		fmt.Printf(" == unknown %v\n", treeValue)
		t.check.AddError("found unknown type %v with value %q", treeValue, treeValue)
	}
	fmt.Println(offset + ">")
}

//ReadFile reads file f and unmarshal it into t, reporting the error
func ReadFile(f string, t interface{}) error {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &t)
}
