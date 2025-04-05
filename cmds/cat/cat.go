package cat

import (
        "bufio"
        "fmt"
        "os"
)

func Main() {
        args := os.Args[2:]
        for _, arg := range args {
                text, err := Cat(arg)
                if err != nil {
                        fmt.Println(err)
                }
                for _, line := range text {
                        fmt.Println(line)
                }
        }
}

func Cat(path string) ([]string, error) {
        var lines []string
        fp, err := os.Open(path)
        defer fp.Close()
        if err != nil {
                fmt.Println(err)
                return nil, err
        }
        s := bufio.NewScanner(fp)
        for s.Scan() {
                lines = append(lines, s.Text())
        }
        return lines, nil
}
