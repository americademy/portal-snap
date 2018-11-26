package main

import (
  "net/http"
  "os/exec"
)

func navigateTo(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Query().Get("url")

  output1, err1 := exec.Command("snapctl", "set", "url=" + url).CombinedOutput()
  if err1 != nil {
    w.Write([]byte(err1.Error()))
  }

  output2, err2 := exec.Command("snapctl", "restart", "codeverse-portal.chromium-mir-kiosk").CombinedOutput()
  if err2 != nil {
    w.Write([]byte(err2.Error()))
  }

  w.Write([]byte("OK" + string(output1) + string(output2)))
}

func main() {
  http.HandleFunc("/navigateTo", navigateTo)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
