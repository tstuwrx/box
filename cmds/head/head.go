package head

import (
        "bufio"
        "flag"
        "fmt"
        "os"
)


func Main() {
        var header bool
        size := flag.Int("n", 10, "")
        flag.Parse()
        args := os.Args[2:]
        if len(args) > 1 {
                header = true
        }
        for _, arg := range args {
                lines, err := Head(arg, *size)
                if err != nil {
                        fmt.Println(err)
                }
                if header {
                        fmt.Printf("==> %s <==\n", arg)
                }
                for _, line := range lines {
                        fmt.Println(line)
                }
        }
}

func Head(path string, n int) ([]string, error) {
        lines := make([]string, 0, n)
        fp, err := os.Open(path)
        defer fp.Close()
        if err != nil {
                return nil, err
        }
        s := bufio.NewScanner(fp)
        for i := 0; s.Scan() && i < n; i++ {
                lines = append(lines, s.Text())
        }
        return lines, nil
}
