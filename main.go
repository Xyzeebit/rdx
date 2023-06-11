package main

import (
    "fmt"
    "path/filepath"
    "io/ioutil"
    _"io/fs"
    "flag"
    "os"
    
    "strconv"
    
    "net/http"
    "html/template"
    
    "github.com/microcosm-cc/bluemonday"
    "github.com/russross/blackfriday/v2"
)

type config struct {
    root string
    depth int
    port int
}

type data struct {
    Title string
    Body template.HTML
}

var templateData data;
var configData config;

func main() {
    help := flag.String("help", "", "Usage help");
    root := flag.String("path", ".", "Directory to search");
    port := flag.Int("p", 4559, "PORT to preview");
    depth := flag.Int("d", 3, "Maximum recursive depth");
    flag.Parse();
    
    configData = config {
        root: *root, 
        depth: *depth, 
        port: *port,
    }
    
    if *help != "" {
        flag.Usage();
        os.Exit(1);
    }
    
    if err := run(); err != nil {
        os.Exit(1);
    }
    
}

func run() error {
    
    // If root is a README parse README
    if isMarkdownReadMe(configData.root) {
        err := parseReadme(configData.root);
        if err != nil {
            return err;
        }
        startServer(configData.port);
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

func parseReadme(root string) error {
    content, err := ioutil.ReadFile(root);
    if err != nil {
        return err
    }
    // parse content with blackfriday and bluemonday
    output := blackfriday.Run(content);
    text := bluemonday.UGCPolicy().SanitizeBytes(output);
    
    templateData.Title = "README EXPLORER";
    templateData.Body = template.HTML(text);
    
    return nil;
    
}

func startServer(port int) {
    
    server := http.Server {
        Addr: string("127.0.0.1:" + strconv.Itoa(port)),
    }
    
    http.HandleFunc("/", requestHandler);
    
    fmt.Println("Server started on port:", port);
    
    server.ListenAndServe();
    
    
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
    
    t, err := template.ParseFiles("template.tmpl");
    if err != nil {
        os.Exit(1);
    }
    
    if err := t.Execute(w, templateData); err != nil {
        os.Exit(1);
    }
    
    
}