## Usage on local

```
$ make run
$ grpcurl -plaintext -d '{"name": "John"}' localhost:7654 helloworld.Greeter/SayHello
-> {"message": "Hello John!"}
```


## Not working on GKE


#### Create cluster
```
$ gcloud container clusters create test-agones-cluster --cluster-version=1.20 --enable-autoscaling --num-nodes=3 --min-nodes=1 --max-nodes=5 --scopes=gke-default --no-enable-autoupgrade --machine-type=e2-micro --tags=game-server --zone=us-west1-a --node-locations=us-west1-a
```

#### Create firewall
```
$ gcloud compute firewall-rules create game-server-firewall \
  --allow udp:7000-8000 \
  --target-tags game-server \
  --description "Firewall to allow game server udp traffic"
```

#### Get static IP
```
$ gcloud compute addresses create test-agones-ip --region=us-west1
-> ${STATIC_IP}
```

#### Install Agones
```
$ helm repo add agones https://agones.dev/chart/stable
$ helm repo update
$ kubectl create namespace agones-system
$ helm install my-release agones/agones --set "agones.allocator.http.loadBalancerIP=${STATIC_IP}" --namespace agones-system
```

#### Apply fleet(hello gRPC server) and fleetautoscaler
```
$ kubectl apply -f ./agones/
```

#### Try to access but failing
<img width="969" alt="スクリーンショット 2021-09-17 0 47 54" src="https://user-images.githubusercontent.com/38310693/133643939-878cb587-0236-49ef-899a-33ff0ca29c6a.png">



## Versions

k8s: 1.20+
agones: 1.17
