package main

import (
	"fmt"
	"log"
)

func callBackMap(conf *config) error {

	resp, err := conf.pokeapiclient.ListLocationAreas(conf.locationAreaNext)
	if err != nil {
		log.Fatalln(err)
	}
	conf.locationAreaNext = resp.Next
	conf.locationAreaPre = resp.Previous
	fmt.Println(resp)
	return nil
}
