## go-practice
Basic syntax repo for programming language GO.

## Resources
12 factors: https://12factor.net/

## Most Commonly Used Commands
```bash
# compile and run Go program
go run hello.go

# module maintenance
go mod xxx   # eg: go mod tidy

# compile a Go file to binary
go build hello.go   # with output directory: go build -o bin/mybinary

# execute the binary
./hello

# format a Go file
go fmt hello.go

# add dependencies to current module and install them
go get dependencies

# compile and install packages and dependencies
go install xxx

# test packages
go test xxx

# run specific go tool
go tool xxx

# report likely mistakes in packages, static code checking
go vet hello.go

# misc - environment variables
GOOS=linux GOARCH=amd64 go build

# misc - full list
$GOROOT/src/go/build/syslist.go
```

## License
MIT
