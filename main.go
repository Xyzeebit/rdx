package main

import (
    "fmt"
    _"io/ioutil"
    _"io/fs"
    _"bytes"
    "flag"
    "os"
    
    _"github.com/microcosm-cc/bluemonday"
    _"github.com/russross/blackfriday/v2"
)

type config struct {
    root string
    depth int
    port int
}

func main() {
    help := flag.String("help", "", "Usage help");
    root := flag.String("path", ".", "Directory to search");
    port := flag.Int("p", 4559, "PORT to preview");
    depth := flag.Int("d", 3, "Maximum recursive depth");
    flag.Parse();
    
    c := config {
        root: *root, 
        depth: *depth, 
        port: *port,
    }
    
    if *help != "" {
        flag.Usage();
        os.Exit(1);
    }
    
    if err := run(c); err != nil {
        os.Exit(1);
    }
    
}

func run(c config) error {
    fmt.Println(c);
    return nil
}