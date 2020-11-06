
package net

import (
    "fmt"
    "github.com/levigross/grequests"
)

type Node struct {

    private_k string
    Public_k string
}

func DefaultNode() *Node {
    return &Node {
        Public_k: genKey(),
        private_k: genKey(),
    }
}

func EmptyNode() *Node {
    return &Node {
        Public_k: nil,
        private_k: nil,
    }
}

func Connect() *grequests.Response {

    resp, err := grequests.Get("http://localhost:8080/connect", nil)

    if err != nil {
        fmt.Println("error in request to connect", err)
    }

    return resp

}
