## Usage on local

```
$ make run
$ grpcurl -plaintext -d '{"name": "John"}' localhost:7654 helloworld.Greeter/SayHello
-> {"message": "Hello John!"}
```


## Not working on GKE


#### 1. Create cluster
```
$ gcloud container clusters create test-agones-cluster --cluster-version=1.20 --enable-autoscaling --num-nodes=3 --min-nodes=1 --max-nodes=5 --scopes=gke-default --no-enable-autoupgrade --machine-type=e2-micro --tags=game-server --zone=us-west1-a --node-locations=us-west1-a
```

#### 2. Create firewall
```
$ gcloud compute firewall-rules create game-server-firewall \
  --allow udp:7000-8000 \
  --target-tags game-server \
  --description "Firewall to allow game server udp traffic"
```

#### 3. Get static IP
```
$ gcloud compute addresses create test-agones-ip --region=us-west1
-> ${STATIC_IP}
```

#### 4. Install Agones
```
$ helm repo add agones https://agones.dev/chart/stable
$ helm repo update
$ kubectl create namespace agones-system
$ helm install my-release agones/agones --set "agones.allocator.http.loadBalancerIP=${STATIC_IP}" --namespace agones-system
```

#### 5. Apply fleet(hello gRPC server) and fleetautoscaler
```
$ kubectl apply -f ./agones/
```

#### 6. Get ①tls.crt ②tls.key ③tls-ca.crt from secrets
mapping encoded ①~③ in local.
```
$ ls
-> ... tls-ca.crt tls.crt tls.key
```


#### 7. Try to access but failing
<img width="978" alt="スクリーンショット 2021-09-18 1 30 22" src="https://user-images.githubusercontent.com/38310693/133822630-0ca16700-db09-461b-9223-361aa7db5e75.png">



## Versions

k8s: 1.20+
agones: 1.17
