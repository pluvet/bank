package main

import (
	"github.com/pluvet/bank/app/infraestructure/frameworks"
)

func main() {
	framework := new(frameworks.GinFramework)
	framework.Run()
}
