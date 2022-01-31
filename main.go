package main

import (
	"menu-manage/api"
	// "menu-manage/config"
)

func main() {
	api.NewApiServer().Run()

}
