# [Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/) --- Stephen Grider

- Install golang and create a go project

```bash
sudo dnf install golang
go version
mkdir -p ~/code/go-proj
cd ~/code/go-proj
go mod init go-proj
```

- Compile and Run

```go
go build main.go // compile
./main // run

go run main.go // compile and run
```

Package types: Executable, Reusable

Types: bool, string, int, float64...
