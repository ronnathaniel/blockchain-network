
package net

import (
    "fmt"
    "github.com/levigross/grequests"
)

type Node struct {

    private_k string
    Public_k string
}

func newNode() *Node {
    return &Node {
        Public_k: genKey(),
        private_k: genKey(),
    }
}

func Connect(client *Node) {

    resp, err := grequests.Get("http://localhost:8080/connect", nil)

    if err != nil {
        fmt.Println("error in request to connect", err)
    }

    resp.JSON(client)

}
