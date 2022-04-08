Deploy:
```
kubectl apply -f mysql-pv.yaml  
kubectl apply -f mysql-sv.yaml  
kubectl apply -f srv.yaml  
```
  
IP:  
```
kubectl describe pod srv  
```
Node: minikube/IP  
  
APIs:  
&emsp;GET:  
&emsp;&emsp;IP/v1/say/hi  
&emsp;&emsp;IP/v1/get/allcontent  
&emsp;POST:  
&emsp;&emsp;IP/v1/post/content  
&emsp;&emsp;&emsp;application/json {"content": "string"}  
