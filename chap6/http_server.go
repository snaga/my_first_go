package main

import (
    "fmt"
    "net/http"
)

var todos []string;

func main() {
    todos = make([]string, 0)
    http.HandleFunc("/todos", showToDo)
    http.HandleFunc("/todos/new", addToDo)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("err ", err)
    }
}

func showToDo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "<html>")
    fmt.Fprintln(w, "<head><title>Todo List</title></head>")

    fmt.Fprintln(w, "<body>")
    fmt.Fprintln(w, "<h1>Todo list</h1>")

    fmt.Fprintln(w, "<ul>")
    for _, todo := range todos {
        fmt.Fprintf(w, "<li>%s</li>\n", todo)
    }
    fmt.Fprintln(w, "</ul>")

    fmt.Fprintln(w, "<h2>Add Todo</h2>")
    fmt.Fprintln(w, `<form method="post" action="/todos/new">`)
    fmt.Fprintln(w, `<input type="text" name="todo">`)
    fmt.Fprintln(w, `<input type="submit" name="Add">`)
    fmt.Fprintln(w, `</form>`)

    fmt.Fprintln(w, "</body>")
    fmt.Fprintln(w, "</html>")
}

func addToDo(w http.ResponseWriter, r *http.Request) {
    receiveValue := r.FormValue("todo")
    todos = append(todos, receiveValue)
    http.Redirect(w, r, "/todos", 303)
}
