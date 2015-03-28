package main

import (
	"fmt"
	"github.com/masayukioguni/go-cloudatcost/cloudatcost"
	"os"
)

func main() {
	Login := os.Getenv("CLOUDATCOAT_API_LOGIN")
	Key := os.Getenv("CLOUDATCOAT_API_TOKEN")
	client, _ := cloudatcost.NewClient(&cloudatcost.Option{Login: Login, Key: Key})

	res, hr, err := client.ConsoleService.Console("254516103")

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return
	}

	fmt.Printf("%v,%v\n", res, err)

}