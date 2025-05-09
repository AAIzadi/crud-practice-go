package main

import (
	appconfig "crud-practice-go/internal/config"
	"fmt"
)

func main() {

	environment := "local"
	appConfig, _ := appconfig.GetConfig(environment)

	fmt.Println(appConfig)
}
