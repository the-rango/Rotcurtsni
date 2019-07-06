package main

import (
    "log"
    "net/http"
    "strings"
)

func redirect(w http.ResponseWriter, r *http.Request) {
    path, ok := r.URL.Query()["path"]
    if !ok || len(path[0]) < 1 {
        log.Println("URL Param 'path' is missing")
        return
    }
    path := path[0] //to rmp/ee

    name, ok := r.URL.Query()["name"]
    if !ok || len(name[0]) < 1 {
        log.Println("URL Param 'name' is missing")
        return
    }
    name := name[0]
    comma := strings.Index(name, ",")
    lastName := name[:comma]

    if path == "ee"{  //to eatereval
      http.Redirect(w, r, "https://eaterevals.eee.uci.edu/browse/instructor#"+lastName, 301)
    } else {          //to ratemyprofessor
      http.Redirect(w, r, "https://www.ratemyprofessors.com/search.jsp?queryBy=teacherName&schoolName=university+of+california+irvine&queryoption=HEADER&facetSearch=true&query="+lastName, 301)
    }
}

func main() {
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
