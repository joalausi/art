package main

import "art/web"

func main() {
	server := &web.ArtServer{}
	server.Run()
}
