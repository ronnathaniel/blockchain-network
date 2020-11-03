
package samo

import (
    // "net"
    "fmt"
    "github.com/gin-gonic/gin"
    "strings"
    "github.com/satori/go.uuid"
)



func RunNet () {
    net := NewNet()
    r := gin.Default()
    netPaths(r, net)
    r.Run(":8080")
}

func netPaths (r *gin.Engine, net *Net){
    r.GET("/connect", func(c *gin.Context) {
        ConnectNode(c, net)
    })
}


func ConnectNode(c *gin.Context, net *Net) {
    node := newNode()
    net.AddNode(node);

    // c.String(200, fmt.Sprintf("public_k - %s\n", node.public_k))
    c.String(200, "Welcome to SamoChain :-)\n")
    c.JSON(200, gin.H{
        "public_k": node.public_k,
        "private_k": node.private_k,
    })
}


type Net struct {
    Nodes []*Node
}

func NewNet() *Net {
    return &Net{}
}

func (net *Net) AddNode(node *Node) {
    net.Nodes = append(net.Nodes, node)
    net.PrintNodes()
}


func genKey() string {

    uuid_v4 := uuid.Must(uuid.NewV4(), nil)
    id := fmt.Sprintf("%s", uuid_v4)
    key := strings.ReplaceAll(id, "-", "")

    // fmt.Printf("private key -- 0x%x", key)
    key = fmt.Sprintf("0x%s", key)
    return key
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
