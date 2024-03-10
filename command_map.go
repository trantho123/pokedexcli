package main

import (
	"errors"
	"fmt"
	"log"
)

func callBackMapNext(conf *config) error {
	resp, err := conf.pokeapiclient.ListLocationAreas(conf.locationAreaNext)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Location areas :")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	conf.locationAreaNext = resp.Next
	conf.locationAreaPre = resp.Previous
	return nil
}

func callBackMapPrev(conf *config) error {
	if conf.locationAreaPre == nil {
		return errors.New("you're in the first page")
	}
	resp, err := conf.pokeapiclient.ListLocationAreas(conf.locationAreaPre)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Location areas :")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	conf.locationAreaNext = resp.Next
	conf.locationAreaPre = resp.Previous
	return nil
}
