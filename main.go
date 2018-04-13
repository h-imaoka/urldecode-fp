package main

import (
    "fmt"
    "io/ioutil"
    "net/url"
    "os"
    "path/filepath"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("help -- urldecode-fp [dir]")
        os.Exit(1)
    }
    dirwalk(os.Args[1])
}

func urldecode(enc string) string {
    u, err := url.QueryUnescape(enc)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    return u
}

func dirwalk(dir string) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        var ofp, rfp string

        ofp = filepath.Join(dir, file.Name())
        rfp = urldecode(ofp)
        if ofp != rfp {
            fmt.Printf("convert %s to %s\n", ofp, rfp)
            if err:= os.Rename(ofp, rfp); err != nil {
                fmt.Println(err)
            }
        }
        if file.IsDir() {
            dirwalk(rfp)
        }
    }

}
