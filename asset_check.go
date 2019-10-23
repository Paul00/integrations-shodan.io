package main

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
)

// global variable declaration
const API_KEY = "<INSERT API KEY>"
const URI_LINK = "https://api.shodan.io/shodan/host/"

// establish connection
func Connection(host string) string { 
  uri := URI_LINK + host + "?key=" + API_KEY
  req, _ := http.NewRequest("GET", uri, nil)
  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  return string(body)
}

func main() {
  var ret string
  ret = Connection(os.Args[1])
  fmt.Println(ret)
}

