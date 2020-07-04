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

  payload := strings.NewReader("{\n    \"jsonrpc\": \"2.0\",\n    \"method\": \"debug_storageRangeAt\",\n    \"params\": [\n        \"0x2b76b3a2fc44c0e21ea183d06c846353279a7acf12abcc6fb9d5e8fb14ae2f8c\",\n        0,\n        \"0x0e0d2c8f7794e82164f11798276a188147fbd415\",\n        \"0x0000000000000000000000000000000000000000000000000000000000000000\",\n        1\n    ],\n    \"id\": 1\n}")

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