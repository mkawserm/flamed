package main

import "github.com/mkawserm/flamed/pkg/app"

func main() {
	err := app.GetApp().Execute()
	if err != nil {
		panic(err)
	}
}
