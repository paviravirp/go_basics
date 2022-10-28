# go_basics

To enable dependency tracking
```
go mod init hello_world
```

To install dependencies
```
cd hello_world
go mod tidy
```

To run a go file
```
cd hello_world
go run .
```

To import a module
```
go mod edit -replace multi_module/input=../input

