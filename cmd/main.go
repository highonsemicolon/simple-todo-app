package main

import (
  "log"
  "encoding/json"
  "sync"
  "net/http"
  "github.com/gorilla/mux"
)

type Todo struct{
  ID string `json:id`
  Title string `json:title`
  Status string `json:status`
}

var (
  todos = make(map[string]Todo)
  mu = sync.Mutex{}
)

func main() {
  
  router := mux.NewRouter()
  
  router.HandleFunc("/", index).Methods("GET")
  router.HandleFunc("/todos", createTodo).Methods("POST")
  router.HandleFunc("/todos", getTodos).Methods("GET")

  log.Println("server is running... ")
  log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
  return
}

func createTodo(w http.ResponseWriter, r *http.Request) {
  var todo Todo
  
  log.Println("in here")
  json.NewDecoder(r.Body).Decode(&todo)


  mu.Lock()
  todos[todo.ID] = todo
  mu.Unlock()

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(todo)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  var todoList[]Todo
  for _, todo := range todos {
    todoList = append(todoList, todo)
  }

  mu.Unlock()

  json.NewEncoder(w).Encode(todoList)
}
