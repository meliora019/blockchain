# A simple implementation of blockchain in golang
---

# Usage

## Starting a node

You can start as many nodes as you want with the following command

`go run -port=<port-number>`

### Show the Blockchain of a node

* `GET 127.0.0.1:8080/chain`

### Generating a new block

* `GET 127.0.0.1:8080/mine`

### Adding a new transaction

* `POST 127.0.0.1:8080/transactions`

* __Body__: A transaction to be added

  ```json
  {
    "Sender": "sender-address",
    "Recipient": "recipient-address",
    "Amount": 1
  }
  ```
  ```
  curl --verbose --header "Content-Type: application/json" --request POST --data '{"Sender":"sender1","Recipient":"recipient1","Amount":1.02}' http://localhost:8080/transactions
  ```

### Register a new node in the network
Currently you must add each new node to each running node.

* `POST 127.0.0.1:8080/nodes`

* __Body__: A list of nodes to add

  ```json
  {
     "nodes": ["http://127.0.0.1:8080", <more-nodes>]
  }
  ```
  ```
  curl --verbose --header "Content-Type: application/json" --request POST --data '{"nodes":["http://127.0.0.1:8080","http://127.0.0.1:8081"]}' http://localhost:8080/nodes
  ```

### Resolving Blockchain differences in each node

* `GET 127.0.0.1:8080/consensus`
