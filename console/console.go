
package console

import (
    "github.com/abiosoft/ishell"
    "github.com/ronnathaniel/blockchain-network/net"
)


type console struct {
    shell *ishell.Shell
    credentials *net.Node
}

func NewConsole() *console {
    return &console{
        shell: shell.New(),
        credentials: net.EmptyNode(),
    }
}

func (c *console) run() {
    c.shell.Run()
}
