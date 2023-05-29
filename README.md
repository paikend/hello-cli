### 1. Install
1. go mod init hello-cli
2. go get -u github.com/spf13/cobra
3. cobra-cli init
4. cobra-cli add greet
5. update cmd/greet.go
   1. Print Hello World
   2. Print Hello ${name}, default name is World
