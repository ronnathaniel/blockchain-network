
package client

import (
    "github.com/abiosoft/ishell"
    "github.com/ronnathaniel/blockchain-network/net"
)

func RunConsole() {

    console := ishell.New()
    initConsole(console)
    console.Run()
}

func initConsole(shell *ishell.Shell) {
    shell.Println("\nSamoChain Interactive Console")
    shell.Println("SamoChain is a subsidiary of SamoCorp\n")
    addConsoleCommands(shell)
}

func addConsoleCommands(shell *ishell.Shell) {
    shell.AddCmd(&ishell.Cmd{
        Name: "create",
        Help: "creates account & wallet at public_k address",
        Func: func(c *ishell.Context) {
            resp := net.Connect()
            shell.Println(resp.Bytes())
        },
    })

    shell.AddCmd(&ishell.Cmd{
        Name: "address",
        Help: "address - returns address of your account",
        Func: func(*ishell.Context) {

        },
    })
}
