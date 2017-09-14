package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fatalErr("Please provide path to a JSON file as a parameter")
	}
	file := os.Args[1]
	data, err := ioutil.ReadFile(file)
	errCheck(err)

	var js map[string]interface{}
	err = json.Unmarshal(data, &js)
	errCheck(err)
}

func fatalErr(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}

func errCheck(err error) {
	if err != nil {
		fatalErr(err.Error())
	}
}
