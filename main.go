package goju //Play function
import (
	"flag"
	"os"

	"github.com/golang/glog"
)

// Play calls traverse check a json files by the rules in the second json file
func Play(json, rule string) error {
	//usage := fmt.Sprintf("usage: %s [options] <data> <rules>\n\noptions are:\n\n", os.Args[0])

	var tree, ruletree map[string]interface{}
	err := ReadFile(json, &tree)
	if err != nil {
		return err
	}
	err = ReadFile(rule, &ruletree)
	if err != nil {
		return err
	}

	tr := &TreeCheck{Check: &Check{}}

	tr.traverse("", tree, ruletree)

	glog.V(1).Infof("Errors       : %d\n", tr.Check.errorHistory.Len())
	glog.V(1).Infof("Checks   true: %d\n", tr.Check.trueCounter)
	glog.V(1).Infof("Checks  false: %d\n", tr.Check.falseCounter)

	return nil
}

func main() {
	var json, rule string
	flag.StringVar(&json, "json", "", "json file to check")
	flag.StringVar(&rule, "rule", "", "check rules")
	flag.Set("logtostderr", "true")
	flag.Set("v", "1")
	flag.Parse()

	if json == "" || rule == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := Play(json, rule)
	if err != nil {
		glog.V(0).Infof("Error: %s", err)
		os.Exit(2)
	}
	os.Exit(0)
}
