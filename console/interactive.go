
package console

import (
    "github.com/abiosoft/ishell"
    "github.com/ronnathaniel/blockchain-network/net"
)

func RunConsole() {

    // console := ishell.New()
    c := emptyConsole()
    initConsole(c)
    c.run()
}

func initConsole(c *console) {
    c.shell.Println("\nSamoChain Interactive Console")
    c.shell.Println("SamoChain is a subsidiary of SamoCorp\n")
    addConsoleCommands(c)
}

func addConsoleCommands(c *console) {
    c.shell.AddCmd(&ishell.Cmd{
        Name: "create",
        Help: "creates account & wallet at public_k address",
        Func: func(*ishell.Context) {
            resp := net.Connect()
            c.shell.Println(resp.String())
        },
    })

    c.shell.AddCmd(&ishell.Cmd{
        Name: "address",
        Help: "address - returns address of your account",
        Func: func(*ishell.Context) {

        },
    })
}
