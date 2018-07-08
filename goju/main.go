package main //Play function
import (
	"flag"
	"os"

	"github.com/endocode/goju"
	"github.com/golang/glog"
)

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

	err := goju.Play(json, rule)
	if err != nil {
		glog.V(0).Infof("Error: %s", err)
		os.Exit(2)
	}
	os.Exit(0)
}
