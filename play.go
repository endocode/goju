package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Traverse is the object collection all data on a traversal
type Traverse struct {
	errorHistory              list.List
	trueCounter, falseCounter int
}

// AddError adds an error to the list of errors,
// format and args are format used to create a formatted error message
func (t *Traverse) AddError(format string, args ...interface{}) {
	errn := fmt.Sprintf("error #%d: ", t.errorHistory.Len())
	fmt.Printf(errn+format+"\n", args...)
	t.errorHistory.PushBack(fmt.Errorf(errn+format, args...))
}

func (t *Traverse) bookkeep(b bool, err error) {
	if err == nil {
		if b {
			t.trueCounter++
		} else {
			t.falseCounter++
		}
	} else {
		errn := fmt.Errorf("error #%d: %s", t.errorHistory.Len(), err.Error())
		t.errorHistory.PushBack(errn)
	}
}

// Equals tracks if both strings are equal
func (t *Traverse) Equals(ruleValue, treeValue string) {
	t.bookkeep(ruleValue == treeValue, nil)
}

// Matches tracks if r matches s as a regular expression
func (t *Traverse) Matches(r, s string) {
	match, err := regexp.MatchString(r, s)
	t.bookkeep(match, err)
}

// Length compares length to the len of the arry
func (t *Traverse) Length(length string, array []interface{}) {
	l, err := strconv.Atoi(length)
	t.bookkeep(l == len(array), err)
}

// Max compares max to val
func (t *Traverse) Max(max string, val int) {
	m, err := strconv.Atoi(max)
	t.bookkeep(m >= val, err)
}

// Min compares min to the val
func (t *Traverse) Min(min string, val int) {
	m, err := strconv.Atoi(min)
	t.bookkeep(m <= val, err)
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

				method := reflect.ValueOf(t).MethodByName(capMethod)
				if method.IsValid() {
					fmt.Printf("%s\t rules %q (%q, %q)\n", offset, capMethod, v, cutString(tv, 40))
					method.Call([]reflect.Value{reflect.ValueOf(v), reflect.ValueOf(tv)})
				} else {
					switch treeValue.Kind() {
					case reflect.String, reflect.Float64, reflect.Bool:
						{
							t.AddError("unknown method %q requested with args(%q, %q)", capMethod, v, cutString(tv, 40))
						}
					}
				}
			}
		} else {
			t.AddError("found unknown rule %v with value %q", rules, rules)
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
		fmt.Printf(" == unknown %q", treeValue)
		t.AddError("found unknown type %v with value %q", treeValue, treeValue)
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
