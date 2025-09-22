package main

import (
    "net/http"
)

// burte force interface each protocol need to implement
type BruteForce interface {
	Parse()
	Client()
	Try() error
}

// for http json bruteForce
type HttpJsonEndpointBruteForce struct {
	EndPoint          string
	Method            string
	UsernameKey       string
	UsernamesFilePath string
	PasswordKey       string
	PassowrdFilePath  string
    Client            *http.Client
}
