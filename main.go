package main //Play function
import (
	"fmt"
	"os"

	"github.com/golang/glog"
)

// Play calls traverse check a json files by the rules in the second json file
func Play(args []string) error {
	if len(args) != 3 {
		glog.V(2).Infof("usage: %s <data> <rules>", os.Args[0])
		return fmt.Errorf("usage: %s <data> <rules>", os.Args[0])
	}
	var tree, ruletree map[string]interface{}
	err := ReadFile(args[1], &tree)
	if err != nil {
		return err
	}
	err = ReadFile(args[2], &ruletree)
	if err != nil {
		return err
	}

	tr := &Traverse{check: &Check{}}

	tr.traverse("", tree, ruletree)

	glog.V(2).Infof("Errors       : %d\n", tr.check.errorHistory.Len())
	glog.V(2).Infof("Checks   true: %d\n", tr.check.trueCounter)
	glog.V(2).Infof("Checks  false: %d\n", tr.check.falseCounter)

	return nil
}

func main() {
	err := Play(os.Args)
	if err != nil {
		glog.V(1).Infof("Error: %s", err)
		panic(err)
	}
}
