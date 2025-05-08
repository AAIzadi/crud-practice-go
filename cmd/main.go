package main

import (
	"crud-practice-go/internal/config"
	"fmt"
)

func main() {

	environment := "local"
	appConfig, _ := config.GetConfig(environment)

	fmt.Println(appConfig.Postgres.Port)
}
