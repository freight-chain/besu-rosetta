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

  payload := strings.NewReader("{\n    \"jsonrpc\": \"2.0\",\n    \"method\": \"eth_getTransactionReceipt\",\n    \"params\": [\n        \"0x96c6830efd87a70020d4d1647c93402d747c05ecf6beeb068dee621f8d13d8d1\"\n    ],\n    \"id\": 1\n}")

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