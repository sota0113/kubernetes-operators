# Description
This is a sample Golang operator of "Kubernetes Operators: Automating the Container Orchestration Platform".  
Whole sample codes in this book are published on [GitHub](https://github.com/kubernetes-operators-book/chapters).  
You can get the book as free e-book from [RedHat official page](https://www.redhat.com/ja/resources/oreilly-kubernetes-operators-automation-ebook) or
[oreilly](https://learning.oreilly.com/library/view/kubernetes-operators/9781492048039/).
NOTE: Some code of this operator has been modifed or defferent from original one due to defference of components' version.  

# Usage
Here are usage of this operator.
## Operator for test on local machine
Operator would work on local machine not as container but code itself.  
### 0. preparation
Make sure Kubernetes Cluster is ready, `kubectl` command is configured, `go` and `operator-sdk` command is ready.  
Cluster version 1.11 or later is preferable.  
This operator was developend and verified on environment as below.
```
OS: macOS Mojave
Kubernetes Cluster Dist : minikube
Minikube version : 1.9.2
Minikube VM driver : virtualbox
Kubernetes Cluster version 1.18.0
kubectl client version : 1.15.2
go version : 1.14.2
operator-sdk version: 0.17.0
```

Check current status of the cluster by executing each commands as below.
```
kubectl get deploy
kubectl get po
kubectl get svc
kubectl get secret
```

### 1. Deploy Custom Resource Definition(CRD)
Deploy CRD in `default` namespace on Cluster.
```
kubectl apply -f deploy/crds/example.com_visitorsapps_crd.yaml
```
Verify API endpoint has been created and it works.  
```
kubectl get VisitorsApp
```
### 2. Start operator.
Start operator in local mode.  
```
operator-sdk run --local --watch-namespace default
```
### 3. Deploy Custom Resource(CR)
Deploy CR in `default` namespace on Cluster.  
```
kubectl apply -f deploy/crds/example.com_v1_visitorsapp_cr.yaml
```
Confirm new `VisitorsApp` resource has been created.  
```
kubectl get VisitorsApp
```
### 4. 
Verify each child resource has been created by Operator.
```
kubectl get deploy
kubectl get po
kubectl get svc
kubectl get secret
```

### Stop Operator
Execute `Ctrl+C` on shell in which operator works.
## Operator as pod on Kubernetes
