package main //Play function
import (
	"fmt"
	"os"
)

// Play calls traverse check a json files by the rules in the second json file
func Play(args []string) error {
	if len(args) != 3 {
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

	var tr Traverse
	tr.check = &Check{}

	tr.traverse("", tree, ruletree)

	fmt.Printf("Errors       : %d\n", tr.check.errorHistory.Len())
	fmt.Printf("Checks   true: %d\n", tr.check.trueCounter)
	fmt.Printf("Checks  false: %d\n", tr.check.falseCounter)

	return nil
}

func main() {
	err := Play(os.Args)
	if err != nil {
		panic(err)
	}
}
