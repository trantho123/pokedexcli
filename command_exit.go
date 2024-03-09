package main

import "os"

func callBackExit(conf *config) error {
	os.Exit(0)
	return nil
}
