
package main

import (
    "github.com/abiosoft/ishell"
    "github.com/ronnathaniel/blockchain-network/net"
)


func main() {
	runShell()
}


func runShell() {
    client := &net.Node{}

    shell := ishell.New()
	shell.Println("\nSamoChain Interactive Shell")
    shell.Println("SamoChain is a subsidiary of SamoCorp\n")
	addShellCommands(shell, client)

	shell.Run()
}

func addShellCommands(shell *ishell.Shell, client *net.Node) {
	  shell.AddCmd(&ishell.Cmd {
		    Name: "connect",
		    Help: "connect helps you join SamoNet",
            Func: func(c *ishell.Context) {
                net.Connect(client)
            },
	  })

    shell.AddCmd(&ishell.Cmd {
        Name: "public_k",
        Help: "public_k is your public address",
        Func: func(c *ishell.Context) {
            shell.Println(client.Public_k)
        },
    })
}
