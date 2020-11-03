
package samo

import (
    "github.com/levigross/grequests"
)

type Node struct {

    private_k string
    public_k string
}

func newNode() *Node {
    return &Node {
        public_k: genKey(),
        private_k: genKey(),
    }
}

func Connect() {

    resp, err := grequests.Get("http://localhost:8080/connect", nil)

    if err != nil {
        fmt.Println("error in request to connect", err)
    }

    fmt.Println(resp.String())

}
