
package net

import (
    "fmt"
    "github.com/levigross/grequests"
)

type Node struct {

    private_k   string  `json:"private_k"`
    public_k    string  `json:"public_k"`
}

func DefaultNode() *Node {
    return &Node {
        public_k: genKey(),
        private_k: genKey(),
    }
}

func EmptyNode() *Node {
    return &Node {
        public_k: "",
        private_k: "",
    }
}

func CustomNode(public_k, private_k string) *Node {
    return &Node {
        public_k: public_k,
        private_k: private_k,
    }
}

func (node *Node) Public_k() string {
    return node.public_k
}

func Connect() *grequests.Response {

    resp, err := grequests.Get("http://localhost:8080/connect", nil)

    if err != nil {
        fmt.Println("error in request to connect", err)
    }

    return resp

}
