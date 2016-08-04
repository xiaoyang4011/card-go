package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/login/", loginHandler)
  http.HandleFunc("/admin/", adminHandler)
}
