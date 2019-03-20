package main

import (
	"IMTools/websever"
)

func main() {

	router := websever.InitRouter()

	router.Run(":8000")

}
