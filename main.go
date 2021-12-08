package main

import (
  "yourwishes/application"
  "yourwishes/application/registry"
  "flag"
  "fmt"
)

func main() {

  appMap := map[string]func() application.RegistryContract{
    "myapp": registry.NewMyapp(),
  }

  flag.Parse()

  app, exist := appMap[flag.Arg(0)]
  if exist {
    application.Run(app())
  } else {
    fmt.Println("You may try 'go run main.go <app_name>' :")
    for appName := range appMap {
      fmt.Printf(" - %s\n", appName)
    }
  }

}
