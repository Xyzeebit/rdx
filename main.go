package main

import (
    "fmt"
    "path/filepath"
    "io/ioutil"
    _"io/fs"
    "bytes"
    "flag"
    "os"
    
    "html/template"
    
    "github.com/microcosm-cc/bluemonday"
    "github.com/russross/blackfriday/v2"
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
    
    // If root is a README parse README
    if isMarkdownReadMe(c.root) {
        readme, err := parseReadme(c.root, "template.tmpl");
        if err != nil {
            return err;
        }
        fmt.Println(string(readme))
    } else {
        //walkDir(c.root);
    }
    return nil
}


func isMarkdownReadMe(path string) bool {
    name := filepath.Base(path);
    
    if name == "README.md" {
        return true;
    }
    return false;
}

func parseReadme(root, tFilename string) ([]byte, error) {
    content, err := ioutil.ReadFile(root);
    if err != nil {
        return nil, err
    }
    // parse content with blackfriday
    output := blackfriday.Run(content);
    text := bluemonday.UGCPolicy().SanitizeBytes(output);
    
    var buffer bytes.Buffer;
    
    t, err := template.ParseFiles(tFilename);
    if err != nil {
        return nil, err;
    }
    
    type data struct {
        Title string
        Body template.HTML
    }
    
    c := data {
        Title: "README EXPLORER",
        Body: template.HTML(text),
    }
    
    if err := t.Execute(&buffer, c); err != nil {
        return nil, err;
    }
    
    return buffer.Bytes(), nil;
    
}