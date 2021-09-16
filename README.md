## Usage

```
$ make run
$ grpcurl -plaintext -d '{"name": "John"}' localhost:7654 helloworld.Greeter/SayHello
-> {"message": "Hello John!"}
```


## Not working on GKE

<img width="961" alt="スクリーンショット 2021-09-16 16 02 28" src="https://user-images.githubusercontent.com/38310693/133565500-4d1b4d14-09c3-4403-b2ae-f55a3cd266fe.png">


## Versions

k8s: 1.20+
agones: 1.17
