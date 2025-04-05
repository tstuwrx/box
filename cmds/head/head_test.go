package head_test

import (
        "fmt"
        "github.com/tstuwrx/box/cmds/head"
)

func ExampleHead() {
        lines, err := head.Head(`testdata/test.txt`, 5)
        if err != nil {
                fmt.Println(err)
        }
        for _, line := range lines {
                fmt.Printf("%s\n", line)
        }

        // Output:
        // one
        // two
        // three
        // four
        // five
}
