package main

import (
	"recipes.com/recipes"
)

func main() {
	recipes.ConnectDB()
	recipes.Init(":3000") // <---- default :3000 si no especificar :4000
}
