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
	files, err := ioutil.ReadDir(".")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
