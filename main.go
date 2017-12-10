package main

import (
    "fmt"
    "net/http"
    "github.com/akamensky/argparse"
    "os"
    "strconv"
    "github.com/danielskowronski/sauron3_core"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path=="/" {
        fmt.Fprintf(w, "%s", index_html)
    } else if r.URL.Path=="/style.css" {
        w.Header().Set("Content-Type", "text/css")
        fmt.Fprintf(w, "%s", style_css)
    } else if r.URL.Path=="/script.js" {
        fmt.Fprintf(w, "%s", script_js)
    } else {
        fmt.Fprintf(w, "Requested URL was not handled by this application.\n%s", r.URL.Path)
    } 
}
func getDefinitions(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, sauron3.DatabaseJson)
}
func getLivecheck(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintf(w, sauron3.CheckAll())
}

func main() {
    parser := argparse.NewParser("sauron3", "a real time eye on your network")

    cfg := parser.String("c", "config", 
        &argparse.Options{Required: false, Help: "Path to config file (default is ./config.yml)"})
    port := parser.String("p", "port",  
        &argparse.Options{Required: false, Help: "Port for webgui (default is 8888)"})
    host := parser.String("b", "bind",  
        &argparse.Options{Required: false, Help: "IP to bind (default is 0.0.0.0)"})
    err := parser.Parse(os.Args)
    if err != nil {
        fmt.Print(parser.Usage(err))
        os.Exit(0)
    }

    if *cfg =="" { *cfg="config.yml" }
    if *port=="" { *port="8888" } else {
        portNumeric, err := strconv.Atoi(*port)
        if err!=nil || portNumeric<1 || portNumeric>65535 {
            fmt.Println("[!!] Invalid port number. Exiting")
            os.Exit(1)
        }
    }

    bindString:=*host+":"+*port
    bindStringFull:=bindString
    if *host=="" { 
        bindStringFull="0.0.0.0:"+*port
    }

    fmt.Println("Welcome to sauron runner!")
    fmt.Println("Copyright 2017 Daniel Skowroński <daniel@dsinf.net>")
    fmt.Println("")

    http.HandleFunc("/", staticHandler)
    http.HandleFunc("/definitions/", getDefinitions)
    http.HandleFunc("/probe/", getLivecheck)
    
    fmt.Println("Loading config file "+*cfg)
    sauron3.LoadConfig(*cfg)
    fmt.Println("Starting to listen at http://"+bindStringFull)
    http.ListenAndServe(bindString, nil)
}

//go:generate go run scripts/packTextAssets.go
