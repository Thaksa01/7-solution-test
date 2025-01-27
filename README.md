# 7 Solution Test

Hello, I'm thaksa nanan, graduated from Mae Fah Luang University, majoring in Software Engineering. I would like to apply for a Backend Developer position.


## Test Question ( backend-challenge )

https://github.com/7-solutions/backend-challenge


# How to run project
firstly, i must setup gRPC service before 
```cmd
make setup
```
or
```cmd
- step 1 
protoc --go_out=. --go-grpc_out=. --proto_path=. api.proto
- step 2 
protoc --grpc-gateway_out=. --go_out=. --go-grpc_out=. --proto_path=. api.proto
```

after than  you can run with this command

```cmd
make start or go run server.go
```
# Optional
- you can download and open API spec with POSTMAN Json file.
## contact
- linkedIn : [https://www.linkedin.com/in/thaksa-nanan-061661236/](doc:linking-to-pages#anchor-links)
- phone : 0826975094
- line with phone number : 0826975094

