package main

import "github.com/zgr126/x/server/manage"

func main() {
	server := manage.New()
	server.Run()
}
