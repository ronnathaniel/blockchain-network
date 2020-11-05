
package net

import (
    "fmt"
    // "strings"

    "github.com/gin-gonic/gin"
)






func RunNet () {
    net := DefaultNet()
    // r := gin.Default()
    netPaths(net)
    net.Run(":8080")
}

func netPaths (net *Net){
    net.Engine.GET("/connect", func(c *gin.Context) {
        AddNode(c, net)
    })
}


func AddNode(c *gin.Context, net *Net) {
    node := newNode()
    net.AddNode(node);

    // c.String(200, fmt.Sprintf("public_k - %s\n", node.public_k))
    c.String(200, "Welcome to SamoChain :-)")
    c.String(200, "Please keep these credentials private:\n")
    c.JSON(200, gin.H{
        "public_k": node.Public_k,
        "private_k": node.private_k,
    })
}


type Net struct {
    Nodes []*Node
    Engine *gin.Engine
}

func DefaultNet() *Net {
    return &Net{
        Nodes: []*Node{},
        Engine: gin.Default(),
    }
}


func (net *Net) Run(port string) {
    // param port:
    //   format - ":${port}"
    net.Engine.Run(port)
}

func (net *Net) AddNode(node *Node) {
    net.Nodes = append(net.Nodes, node)
    net.PrintNodes()
}


func (net *Net) NodeCount() int {
    return len(net.Nodes)
}


func (net *Net) PrintNodes() {
    fmt.Printf("there are %d nodes connected\n", net.NodeCount())
}

func (net *Net) GetNodes() []*Node {
    // var nodes []*Node
    nodes := []*Node{}

    for _, n := range net.Nodes {
        nodes = append(nodes, n)
    }

    return nodes
}
