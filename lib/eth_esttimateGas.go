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

  payload := strings.NewReader("{\n    \"jsonrpc\": \"2.0\",\n    \"method\": \"eth_estimateGas\",\n    \"params\": [\n        {\n            \"from\": \"0x687422eea2cb73b5d3e242ba5456b782919afc85\",\n            \"to\": \"0xdd37f65db31c107f773e82a4f85c693058fef7a9\",\n            \"value\": \"0x1\"\n        },\n        \"latest\"\n    ],\n    \"id\": 1\n}")

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