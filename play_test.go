package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readFile(f string, t interface{}) error {
	b, err := ioutil.ReadFile("data/" + f + ".json")
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &t)
}

func TestPlayWithSimpleItemArray(t *testing.T) {
	testPodWithRules(t, "itempods", "itemrule", 2, 0, 0, 1)
}

func TestPlayWithSimplePods(t *testing.T) {
	testPodWithRules(t, "simplepods", "simplerule", 2, 0, 0, 4)
}

func TestPlayWithImagePod(t *testing.T) {
	testPodWithRules(t, "imagepod", "imagerule", 1, 0, 0, 2)
}

func TestPlayWithImageFullPod(t *testing.T) {
	testPodWithRules(t, "fullpods", "fullpodimagerule", 4, 0, 1, 30)
}

func TestPlayWithUnknownMethod(t *testing.T) {
	testPodWithRules(t, "itempods", "unknownmethodrule", 2, 1, 0, 0)
}

func TestPlayWithoutFile(t *testing.T) {
	var tree map[string]interface{}
	var tr Traverse
	err := readFile("notexisting", &tree)
	tr.bookkeep(true, err)
	assert.NotNil(t, tr.errorHistory.Front())
}

func testPodWithRules(t *testing.T, treeFile, ruleFile string,
	treeLengthExpected, errorLengthExpected,
	falseExpected, trueExpected int) {
	var tr Traverse
	var tree, ruletree map[string]interface{}
	assert.Nil(t, readFile(treeFile, &tree), treeFile)
	assert.Nil(t, readFile(ruleFile, &ruletree), ruleFile)

	assert.Len(t, tree, treeLengthExpected, "tree length")
	tr.traverse("", tree, ruletree)
	assert.Equal(t, errorLengthExpected, tr.errorHistory.Len(), "errors")
	if errorLengthExpected == 0 {
		assert.Nil(t, tr.errorHistory.Front(), "error history")
	}
	assert.Equal(t, falseExpected, tr.falseCounter, "falseCounter")
	assert.Equal(t, trueExpected, tr.trueCounter, "trueCounter")
}
func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
