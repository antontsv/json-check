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

	if len(data) < 0 || (data[0] != '{' && data[0] != '[') {
		errCheck(fmt.Errorf("expected beginning of JSON object or array at position 0"))
	}

	// Just use validation part, not hydrate any data structure
	err = json.Unmarshal(data, nil)
	e, ok := err.(*json.SyntaxError)
	if ok {
		err = fmt.Errorf("syntax error near position %d: %s", e.Offset, e.Error())
		errCheck(err)
	}

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
