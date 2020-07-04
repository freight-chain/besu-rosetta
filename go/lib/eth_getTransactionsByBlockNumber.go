package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "http://127.0.0.1:8545"
  method := "POST"

  payload := strings.NewReader("{\n    \"jsonrpc\": \"2.0\",\n    \"method\": \"eth_getTransactionByBlockNumberAndIndex\",\n    \"params\": [\n        \"latest\",\n        \"0x0\"\n    ],\n    \"id\": 1\n}")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}