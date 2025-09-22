package main

import (
	"fmt"
	"os"
)

// endpoing we will handle only json
// tacking textfile of usernames and password

// Method:Url infos=username:file.txt:password:password.txt

func ErrorLog(msg string, val string) {
	fmt.Println(val, ":", msg)
	os.Exit(1)
}

func BruteForceRun(e BruteForce) {
	e.Parse()
	e.Try()
	fmt.Println(e)
}

func main() {

	var h HttpJsonEndpointBruteForce

	BruteForceRun(&h)

}
