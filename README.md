For Japanese, see "日本語".

# About This repo
This repo includes aggregated information of Operator with Golang.  
Almost all of contents are for beginners because author of this repo is one of them.  
If anything goes wrong in this repo, please open issue.(e.g. uncorrent or out of date knowledge, defect of explanation or copyright violation)  
There are several element related to operator such as SRE, but they are out of scope in this repo.  

# Prerequirements
## components

Components to develop operator in Golang and verify sample operators of this repo as below.  
```
Kubernetes Cluster: v.1.11 or later
kubectl command: v1.15.2 or later
runtime of Golang: go1.14.2 or later
Operator SDK: v0.17.0 or later
```

## key concepts
Before diving into operator, it is indispensalbe to understand some concept related to Kubernetes.  
keywords:  
kubernetes api-server  
declarative programming model  
reconcile loop  
Custom Resource Definition  
Custom Resource  
  
