
package console

import (
    "github.com/abiosoft/ishell"
    "github.com/ronnathaniel/blockchain-network/net"
)


type console struct {
    shell *ishell.Shell
    credentials *net.Node
}

func emptyConsole() *console {
    return &console {
        shell: ishell.New(),
        credentials: net.EmptyNode(),
    }
}

func customConsole(public_k, private_k string) *console {
    return &console {
        shell: ishell.New(),
        credentials: net.CustomNode(public_k, private_k),
    }
}

func (c *console) run() {
    c.shell.Run()
}
