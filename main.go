package main

import (
    "fmt"
    "path/filepath"
    "io/ioutil"
    "io/fs"
    "flag"
    "os"
    
    "strings"
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
    Path string
    Title string
    Body template.HTML
}

var templateData map[string]*data;
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
    
    templateData = make(map[string] *data);
    
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
        walkPath(configData.root, configData.depth);
        startServer(configData.port)
    }
    return nil
}


func isMarkdownReadMe(path string) bool {
    name := filepath.Base(path);
    name = strings.ToLower(name);
    if name == "readme.md" {
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
    
    if _, ok := templateData[root]; !ok {
        d := data {
            Path: root,
            Title: "README EXPLORER",
            Body: template.HTML(text),
        }
        templateData[root] = &d;
    }
    
    return nil;
    
}

func walkPath(path string, maxDepth int) error {
    err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err;
        }
        if d.IsDir() && strings.Count(path, string(os.PathSeparator)) > maxDepth {
            return fs.SkipDir
        }
        
        base := filepath.Base(path);
        base = strings.ToLower(base);
        if base == "readme.md" {
            parseReadme(path);
        }
        return nil
    });
    return err;
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
    if r.URL.Path == "/" {
        if err := t.Execute(w, templateData); err != nil {
            os.Exit(1);
        }
    }
    //m, ok := templateData[r.URL.Path];
    
    
}