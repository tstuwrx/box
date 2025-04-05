package tail_test

import (
        "fmt"
        "github.com/tstuwrx/box/cmds/tail"
)

func ExampleTail() {
        lines, err := tail.Tail(`testdata/test.txt`, 5)
        if err != nil {
                fmt.Println(err)
        }
        for _, line := range lines {
                fmt.Printf("%s\n", line)
        }

        // Output:
        // six
        // seven
        // eight
        // nine
        // ten
}
