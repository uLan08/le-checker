package lechecker

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

const baseURL = "https://crt.sh/"

// CheckDomain makes the request to crt.sh
func CheckDomain(domain string) {
	ro := grequests.RequestOptions{
		Params: map[string]string{"Identity": domain},
	}

	res, err := grequests.Get(baseURL, &ro)

	if err != nil {
		log.Fatal("Unable to make request:", err)
	}

	fmt.Println(res)
}
