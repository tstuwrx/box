package cat_test

import (
        "fmt"
        "github.com/tstuwrx/box/cmds/cat"
)

func ExampleCat() {
        text, err := cat.Cat(`testdata/test.txt`)
        if err != nil {
                fmt.Println(err)
        }
        for _, line := range text {
                fmt.Println(line)
        }

        // Output:
        // foo
        // bar
        // fizz
        // buzz
        // and
        // or
}
