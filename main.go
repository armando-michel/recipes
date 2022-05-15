package main

import (
	"recipes.com/recipes"
)

func main() {
	recipes.ConnectDB()
	recipes.Init(":4001") // <---- default :3000 si no especificar :4000
}
