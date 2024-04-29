package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func DebugCWD() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Working Directory: ", cwd)
	ioutil.ReadFile("helper/testdata/must-gather.sample/namespaces/openshift-etcd.yaml")

}
