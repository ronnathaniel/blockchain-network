
package chain

import (
    "net"
    "fmt"
    "github.com/gin-gonic/gin"
)



func RunNet () {
    network := initNet()
    r := gin.Default()
    r.GET("/new-node", func(c *gin.Context) {
        newNode(c, network)
    })

    r.Run(":8080")
}


func newNode(c *gin.Context, n *Net) {
    public_k := "public_k"
    node := initNode(public_k)
    n.AddNode(node);

    c.String(200, public_k)
}


type Net struct {
    Nodes []*Node
}

func initNet() *Net {
    return &Net{}
}

func (net *Net) AddNode(node *Node) {
    net.Nodes = append(net.Nodes, node)
    net.PrintNodes()
}

func (net *Net) PrintNodes() {
    // for _, elem := range(net.Nodes) {
    //     fmt.Printf("")
    // }
    fmt.Printf("net add - %p\n", net)
    fmt.Printf("there are %d nodes connected\n", len(net.Nodes))
}

func LocalAddresses () {
    ifaces, err := net.Interfaces()
    if err != nil {
        fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
        return
    }
    for _, i := range ifaces {
        addrs, err := i.Addrs()
        if err != nil {
            fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
            continue
        }
        for _, a := range addrs {
            switch v := a.(type) {
            case *net.IPAddr:
                fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())

            case *net.IPNet:
                fmt.Printf("Name %v : ip type %s [%v/%v]\n", i.Name, v, v.IP, v.Mask)
            }

        }
    }
}
