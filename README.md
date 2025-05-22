# 🧭 GoLang Learning Roadmap

A structured roadmap to learn GoLang from beginner to advanced levels.

---
### [Projects I have Completed ✅](#projects-list)

## 📘 Phase 1: Basics (Week 1–2)

**🎯 Goal:** Understand the fundamentals and syntax.

### 🔑 Topics:

- Installing Go & setting up `GOPATH`
- Basic Syntax
- Variables, Constants, Data Types
- Functions
- Conditionals (`if-else`, `switch`)
- Loops (`for`)
- Arrays, Slices, Maps
- Structs & Methods
- Pointers

### 📚 Resources:

**YouTube:**
- [Tech With Tim – Golang for Beginners](https://www.youtube.com/playlist?list=PLzMcBGfZo4-lB8MZfHPLTEHO9zJDDLpYj)

**Web:**
- [Tour of Go (official)](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [Learn Go in Y Minutes](https://learnxinyminutes.com/docs/go/)
- [Go.dev Docs](https://go.dev/doc/)

### ✅ Practice Projects (Beginner):

- Simple Calculator (CLI)
- In-Memory To-Do List
- Temperature Converter
- Word Counter (reads from text file)

---

## 🚀 Phase 2: Intermediate Concepts (Week 3–4)

**🎯 Goal:** Learn Go-specific concepts and standard library features.

### 🔑 Topics:

- Interfaces
- Error Handling (`errors`, `panic`, `recover`)
- Packages & Modules (`go mod`)
- Concurrency (`goroutines`, `channels`)
- File I/O
- HTTP Server using `net/http`
- JSON Handling (`marshal`, `unmarshal`)
- Unit Testing with `testing` package

### 📚 Resources:

**YouTube:**
- [JustForFunc (Google Developers)](https://www.youtube.com/playlist?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ)

**Web:**
- [Go Standard Library Tour](https://pkg.go.dev/std)
- [Go Concurrency Patterns – Go Blog](https://go.dev/blog/pipelines)

### ✅ Practice Projects (Intermediate):

- REST API for Task Manager
- URL Shortener
- CLI Password Manager
- Markdown to HTML Converter
- JSON to CSV Converter

---

## 🔬 Phase 3: Advanced Topics (Optional)

Explore deeper concepts in Go:

- Advanced use of Goroutines & Channels
- Worker Pools
- Middleware in HTTP Servers
- Using Third-party Packages
- Dependency Injection
- Build Tools: `go build`, `go test`, `go fmt`
- ORM (e.g., GORM)
- Docker + Go projects

---

## 💡 Project Ideas for Portfolio

- **Blog API:** CRUD with SQLite/PostgreSQL
- **Chat Server:** Real-time using WebSockets
- **Weather CLI App:** Integrate with external API
- **Personal Finance Tracker:** Manage expenses and budgets

---

## 📦 Helpful Libraries

| Feature             | Package              |
|---------------------|----------------------|
| HTTP Router         | `gorilla/mux`, `chi` |
| ORM                 | `gorm`               |
| Configuration       | `viper`              |
| CLI Apps            | `cobra`              |

---

## 🛠️ Tools to Install

- Go compiler: `go install`
- IDE: VS Code + Go Extension
- Debugger: [Delve](https://github.com/go-delve/delve)
- API Testing: Postman (for backend & API work)

---

Happy Coding! 🎉

### Projects List 
- Simple Calculator (CLI) 
- To-do List (In Memory)

---

# 🧭 GoLang Learning Roadmap

A structured roadmap to learn GoLang from beginner to advanced levels.

---

## 📘 Key Learnings from To-do List Project

In the **To-do List (In Memory)** project, I learned:

### 🔢 Use of `strconv.Atoi()`

- **Purpose:** Converts a string to an integer.
- **Example:** User input from `bufio.Reader` returns a string like `"3"`.  
  Using `strconv.Atoi("3")` returns the integer `3`.

```go
index, err := strconv.Atoi(indexInput)
```

- **Error Handling:**  
  If the input is not a valid number (e.g., `"abc"`), it returns an error.
  Always check for errors to avoid crashes.

```go
if err != nil {
    fmt.Println("Invalid input:", err)
    return
}
```

---

### ✂️ Deleting an Element from a Slice

Let’s say you have a slice:

```go
todoList := []string{"Task1", "Task2", "Task3", "Task4"}
```

You want to delete the element at `index = 2` (1-based, i.e., delete `"Task2"`):

```go
todoList = append(todoList[:index-1], todoList[index:]...)
```

- `todoList[:index-1]` → elements before `"Task2"` (i.e., `"Task1"`)
- `todoList[index:]` → elements after `"Task2"` (i.e., `"Task3"`, `"Task4"`)

`append()` combines them, effectively removing `"Task2"`.

📌 **Important:**  
`todoList[index:]` means:  
> From index **`index`** **onward**, including the element at that index — **not** the next one.
