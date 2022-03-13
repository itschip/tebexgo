# TebexGo
Go package for the Tebex Plugin API

### Prerequisites
Golang

### Installing

```
go get github.com/itschip/tebexgo
```

### Getting started

#### Retrieve a package by id

```go
import (
  "fmt"

  "github.com/itschip/tebexgo"
)

func main() {
	// Creates a new session
	s := tebexgo.New("your_secret")
	
	// You should handle your error here
	pkg, err := s.GetPackage("package_id")
	
	fmt.Println(pkg.Name) // Prints out the package name
}
```