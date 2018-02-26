package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Traverse is the object collection all data on a traversal
type Traverse struct {
	lastError error
}

func (t Traverse) applyRule(offset string, treeValue reflect.Value, rulesValue reflect.Value, rules interface{}) {
	if rulesValue.Kind() == reflect.Map {
		m, ok := rules.(map[string]interface{})
		if ok {
			first := true
			for k, v := range m {
				if !first {
					fmt.Println(offset + "\t")
				}
				fmt.Printf("rules %q (%q, %q)", k, v, treeValue.Interface())
				first = false
			}
		} else {
			t.lastError = fmt.Errorf("found unknown rule %v with value %q", rules, rules)
		}
	}
}

func (t Traverse) traverse(offset string, tree interface{}, rules interface{}) {
	if tree == nil || rules == nil {
		fmt.Printf(offset+"< traverse t is nil=%t r is nil=%t>\n", tree == nil, rules == nil)
		return
	}
	treeValue := reflect.ValueOf(tree)
	rulesValue := reflect.ValueOf(rules)
	fmt.Printf("\n%s< traverse %v\n", offset, treeValue.Type())

	switch treeValue.Kind() {
	case reflect.Slice, reflect.Array:
		for i, vi := range tree.([]interface{}) {
			index := fmt.Sprintf("%d:", i)
			t.traverse(index+offset+"\t", vi, rules)
		}
	case reflect.Map:
		for k, v := range tree.(map[string]interface{}) {
			r, ok := rulesValue.Interface().(map[string]interface{})
			if ok {
				t.traverse(offset+"\t "+k+": ", v, r[k])
			} else {
				t.applyRule(offset, treeValue, rulesValue, r)
			}
		}

	case reflect.String, reflect.Float64, reflect.Bool:
		t.applyRule(offset, treeValue, rulesValue, rules)
	default:
		fmt.Printf(" == unknown %q", treeValue)
		t.lastError = fmt.Errorf("found unknown type %v with value %q", treeValue, treeValue)
	}
	fmt.Println(offset + ">")
}

//Play function
func Play(pods string) {

	var y map[string]interface{}
	json.Unmarshal([]byte(pods), &y)

	item1 := y["items"].([]interface{})
	fmt.Printf("%v\n", reflect.TypeOf(item1))
	item2 := item1[1]
	fmt.Printf("%v\n", reflect.TypeOf(item2))
	item3 := item2.(map[string]interface{})
	fmt.Printf("%v\n", reflect.TypeOf(item3))
	item4 := item3["metadata"]
	fmt.Printf("%v\n", reflect.TypeOf(item4))
	item5 := item4.(map[string]interface{})
	fmt.Printf("%v\n", reflect.TypeOf(item5))
	item6 := item5["generateName"]

	fmt.Printf("%v\n", reflect.TypeOf(item6))
	t := Traverse{}
	t.traverse("", y, nil)

	//	fmt.Println("nodes = %s", pods)
}

func main() {
	Play("")
}
