package main

import (
	"fmt"
	"github.com/mkawserm/flamed/pkg/app"
)

func main() {
	err := app.GetApp().Execute()
	if err != nil {
		fmt.Println(err)
	}
}
