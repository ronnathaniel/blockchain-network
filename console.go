
package main

import (
    "github.com/abiosoft/ishell"
    "github.com/ronnathaniel/blockchain-network/chain"
)


func main() {
	runShell()
}


func runShell() {
    shell := ishell.New()
	shell.Println("\nSamoChain Interactive Shell")
	shell.Println("SamoChain is a subsidiary of SamoCorp\n")
	addShellCommands(shell)

	shell.Run()
}

func addShellCommands(shell *ishell.Shell) {
	shell.AddCmd(&ishell.Cmd{
		Name: "connect",
		Help: "seek help",
		Func: func(c *ishell.Context) {
			chain.Connect()

		},
	})

}
