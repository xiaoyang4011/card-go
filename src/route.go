package main

import (
    "net/http"
    "log"
    "html/template"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/" {
      http.Redirect(w, r, "/login/index", http.StatusFound)
  }

  t, err := template.ParseFiles("template/html/404.html")
  if (err != nil) {
      log.Println(err)
  }
  t.Execute(w, nil)
}
