
package chain




type Node struct {

    private_k string
    public_k string
}

func initNode (public_k string) *Node {
    return &Node {
        public_k: public_k,
        private_k: "private_k",
    }
}
