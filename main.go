package main

import (
  "os"
  "strings"
)

func main() {
  a := App{}
  a.Initialize(
    os.Getenv("APP_DB_USERNAME"),
    strings.TrimSpace(os.Getenv("APP_DB_PASSWORD")),
    os.Getenv("APP_DB_NAME"),
  )
  a.Run(":8001")
}
