package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func handleFatal(err error) {
	if err != nil {
		panic(err)
	}
}

func getJSON() map[string]string {
	nc := make(map[string]string)
	jbuf, err := ioutil.ReadFile("authdata.json")
	handleFatal(err)
	err = json.Unmarshal(jbuf, &nc)
	handleFatal(err)
	return nc
}

func authenticate(user, token string) bool {
	view := getJSON()
	expt, ok := view[user]
	if !ok {
		return false
	} else if expt != token {
		return false
	} else {
		return true
	}
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(7)
	}
	mbuf, err := ioutil.ReadFile(os.Args[1])
	handleFatal(err)
	dslice := strings.Split(string(mbuf), "\n")
	if authenticate(dslice[0], dslice[1]) == true {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
