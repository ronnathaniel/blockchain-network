
package console

import (
    // "encoding/json"
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
            // c.shell.Println(resp.String())
            // node := net.EmptyNode()
            var n net.Node
            resp.JSON(&n)
            c.shell.Println(n)

            // resp_bytes := resp.Bytes()
            // c.shell.Println("bytes", resp_bytes)
            // err := json.Unmarshal(resp_bytes, &node)
            // if err != nil {
            //     c.shell.Println("error! - ", err)
            // }
            // c.shell.Println(node)



            // WORKS
            // c.shell.Println(Bytes())
            //
            // c.shell.Println(node)
            //
            // err := resp.JSON(*node)
            //
            // if err != nil {
            //     c.shell.Println("error!! - ", err)
            // } else {
            //     c.shell.Println("no error")
            // }
            //
            // c.shell.Println(node)

        },
    })

    c.shell.AddCmd(&ishell.Cmd{
        Name: "address",
        Help: "address - returns address of your account",
        Func: func(*ishell.Context) {

        },
    })
}
