package lechecker

import (
	"fmt"
	"log"
	"regexp"

	"github.com/levigross/grequests"
)

const baseURL = "https://crt.sh/"

func request(ro *grequests.RequestOptions) string {
	res, err := grequests.Get(baseURL, ro)
	if err != nil {
		log.Fatal("Unable to make request:", err)
	}
	return res.String()
}

func getCAIds() []string {
	ro := grequests.RequestOptions{
		Params: map[string]string{"CAName": "%s Encrypt%"},
	}
	res := request(&ro)
	caIDRe := regexp.MustCompile(`\?caid=(\d+)`)
	match := caIDRe.FindAllStringSubmatch(res, -1)
	ids := []string{}
	for _, v := range match {
		ids = append(ids, v[1])
	}
	return ids
}

// CheckDomain makes the request to crt.sh
func CheckDomain(domain string) {
	ro := grequests.RequestOptions{
		Params: map[string]string{"CAName": "%s Encrypt%"},
	}

	res, err := grequests.Get(baseURL, &ro)

	if err != nil {
		log.Fatal("Unable to make request:", err)
	}

	fmt.Println(res)
}
