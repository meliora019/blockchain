package main

import (
  "net/http"
  "flag"
  _ "blockchain/node"
  "blockchain/api"
)

func main() {
  port := flag.String("port", "8080", "Port for our Web Server")
  flag.Parse()

  http.HandleFunc("/mine", api.Mine) //GET
	http.HandleFunc("/transactions", api.Transactions) //POST
	http.HandleFunc("/chain", api.Chain) //GET
  http.HandleFunc("/nodes", api.Nodes) //POST
  http.HandleFunc("/consensus", api.Consensus) //GET

	err := http.ListenAndServe(":" + *port, nil)
  if err != nil {
    panic(err)
  }
}
