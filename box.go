package main

import (
        "os"

        "github.com/tstuwrx/box/cmds/head"
        "github.com/tstuwrx/box/cmds/tail"
        "github.com/tstuwrx/box/cmds/cat"
)

func main() {
        arg := os.Args[1]
        switch arg {
        case "head":
                head.Main()
        case "tail":
                tail.Main()
        case "cat":
                cat.Main()
        }
}
