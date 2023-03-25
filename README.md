## About
This module is used to interact with **Student Services**.  
As a result of this research, we will learn to implement a
a more modern and attractive gRPC architecture.

## Development Guide
1. Install the dependency `buf.build` and `protoc-gen-go`
    ```bash
    go install github.com/bufbuild/buf/cmd/buf@latest
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    ```

2. Building the server stub and mocks  
    ```bash
    buf generate proto/
    ```
3. Installing main project dependency
    ```
    go mod tidy
    ```
4. Running the server
    ```
    go run main.go 
    ```
   
By default service exposed at port `:5080` 

## Whats Next ?
1. Fork this repo
2. Implement `Update Student Method`

Happy Coding!

## Author
Khafid Prayoga: khafidp@pm.me
- [Github](https://github.com/khafidprayoga/student-svc)
- [Medium](https://khafidprayoga.medium.com)
- [LinkedIn](https://www.linkedin.com/in/khafidprayoga/)
