# Golang Package and Module demo

### Notes

- Always start your module with `go mod init $URL`
- A package is nothing but a directory inside your Go workspace
- Variable, function, struct scopes are depends on placement of the package
- Same package will the under single dir, all functions, variables are exported by default
- Variables, Functions exported to other packages by camelcase naming pattern

### Example

- `hi.go` comes under `main` package. So all the variables, functions exported by default ( myName, sayHello)
- `bye/bye.go` has `goodbye` package which has `SayBye()` function. ( camelcase). So it is exported by default but not `byeMessage` variable
- It is very difficult to find `goodbye` package as we have `bye & hello` folders. So always create folder with package name. It is a best practice -`hello\hello.go` has `hello` package whih has `SayHello` ( camelcase). So it is exported by default

### 3rd party app expxort level

- If any 3rd party app wants our package, then it need to import it via `go get` or import the URL
- This module is consumed by [https://github.com/akilans/golang-class/tree/main/week5/01-pack-mod-demo]
