## Usage

```
$ make run
$ grpcurl -plaintext -d '{"name": "John"}' localhost:7654 helloworld.Greeter/SayHello
-> {"message": "Hello John!"}
```
