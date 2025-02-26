package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type TodoList struct {
	TodoCount int
	Todos     []string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// create a writer that will allow us to write to the browser
func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	checkErr(err)
}

// APIS
func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	write(writer, "Hello Internet")
}

// method to get strings from the file
func getString(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	checkErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	checkErr(scanner.Err())
	return lines
}

// TO viewTodo with TOdo files
func viewTodoHandler(writer http.ResponseWriter, request *http.Request) {
	toDoVals := getString("todos.txt")

	// Parse todos to template
	view_template, err := template.ParseFiles("view.html")
	checkErr(err)

	// create todos structs
	todos := TodoList{
		TodoCount: len(toDoVals),
		Todos:     toDoVals,
	}

	err = view_template.Execute(writer, todos)
	checkErr(err)
}

func homeTodoHandler(writer http.ResponseWriter, request *http.Request) {
	create_template, err := template.ParseFiles("home.html")
	checkErr(err)
	err = create_template.Execute(writer, nil)
	checkErr(err)
}

// function to create todos that will come from homeTodoHandler bcoz that is func we used to create todos and home.html will have the create todo view which will use this func
func createTodoHandler(writer http.ResponseWriter, request *http.Request) {
	todo := request.FormValue("todo")
	options := os.O_APPEND | os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	checkErr(err)
	defer file.Close()
	_, err = fmt.Fprintln(file, todo)
	checkErr(err)
	http.Redirect(writer, request, "/view", http.StatusFound)
}

func getUserId(writer http.ResponseWriter, request *http.Request) {
	user_id := request.PathValue("user_id")
	fmt.Println(user_id)
}

func main() {
	http.HandleFunc("/hello", helloWorld)

	// Todos
	http.HandleFunc("/view", viewTodoHandler)
	http.HandleFunc("/home", homeTodoHandler)
	http.HandleFunc("/create", createTodoHandler)

	http.HandleFunc("POST /dynamic_route/{user_id}", getUserId)

	// Run it
	err := http.ListenAndServe("localhost:8080", nil)
	checkErr(err)
}
