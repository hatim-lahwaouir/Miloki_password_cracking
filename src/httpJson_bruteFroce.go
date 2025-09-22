package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func (e *HttpJsonEndpointBruteForce) Parse() {
	var (
		hInfo  []string
		BFinfo []string
		url    string
		info   string
	)

	flag.StringVar(&url, "url", "", "http endpoint specify the endpoint and method, exmaple:\nPOST:exmaple.com/login")
	flag.StringVar(&info, "infos", "", "infos to used for brute-force, example:\nusername:file_of_usernames:password:file_of_passwords, the key here will be used as key in json body")
	flag.Parse()

	hInfo = strings.SplitN(url, ":", 2)
	BFinfo = strings.Split(info, ":")

	if len(hInfo) != 2 {
		ErrorLog("Invalid url  try --help for more info ", info)
	}

	if len(BFinfo) != 4 {
		ErrorLog("Invalid brute-force info  try --help for more info ", info)
	}

	e.EndPoint = hInfo[0]
	e.Method = hInfo[1]
	e.UsernameKey = BFinfo[0]
	e.UsernamesFilePath = BFinfo[1]
	e.PasswordKey = BFinfo[2]
	e.PassowrdFilePath = BFinfo[3]
}

func (e *HttpJsonEndpointBruteForce) Client() {
    e.Client = &http.Client{}
}

func (e *HttpJsonEndpointBruteForce) Try() error {

	var (
		req    *http.Request
		resp   *http.Response
	)

    req = NewRequest(e.Method,e.EndPoint, nil)

    resp, _ :=  e.Client.Do(req)
	fmt.Println(req, client)


	return nil
}
