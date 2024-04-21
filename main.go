package main

import "toyproject_recruiting_community/infra"

func main() {
	infra.Init()
	db := infra.SetupDB()

}
