package main

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"

	"github.com/golang/glog"
)

//Check collects all information and methods about checks
type Check struct {
	errorHistory              list.List
	trueCounter, falseCounter int
}

// AddError adds an error to the list of errors,
// format and args are format used to create a formatted error message
func (t *Check) AddError(format string, args ...interface{}) {
	errn := fmt.Sprintf("error #%d: ", t.errorHistory.Len())
	glog.V(2).Infof(errn+format, args...)
	t.errorHistory.PushBack(fmt.Errorf(errn+format, args...))
}

func (t *Check) bookkeep(b bool, err error) {
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
func (t *Check) Equals(ruleValue, treeValue string) {
	t.bookkeep(ruleValue == treeValue, nil)
}

// Matches tracks if r matches s as a regular expression
func (t *Check) Matches(r, s string) {
	match, err := regexp.MatchString(r, s)
	t.bookkeep(match, err)
}

// Length compares length to the len of the arry
func (t *Check) Length(length string, array []interface{}) {
	l, err := strconv.Atoi(length)
	t.bookkeep(l == len(array), err)
}

// Max compares max to val
func (t *Check) Max(max string, val int) {
	m, err := strconv.Atoi(max)
	t.bookkeep(m >= val, err)
}

// Min compares min to the val
func (t *Check) Min(min string, val int) {
	m, err := strconv.Atoi(min)
	t.bookkeep(m <= val, err)
}

// Eval evaluates an expression
func (t *Check) Eval(r, s string) {
	//ToDO
	glog.V(2).Infof("Not implemented")
}
