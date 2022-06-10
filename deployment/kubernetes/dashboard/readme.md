# kubernetes dashboard

# [helm install guide](https://artifacthub.io/packages/helm/k8s-dashboard/kubernetes-dashboard)

## Add kubernetes-dashboard repository
```
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/ 
```
## Deploy kubernetes-dashboard chart
```
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard -n kubernetes-dashboard --create-namespace
```

## Create a simple user for login
```
kubectl apply -f ./dashboard/access.yaml
```

## Getting a Bearer Token
```
kubectl -n kubernetes-dashboard create token admin-user
```

## Expose dashboard

```
export POD_NAME=$(kubectl get pods -n kubernetes-dashboard -l "app.kubernetes.io/name=kubernetes-dashboard,app.kubernetes.io/instance=kubernetes-dashboard" -o jsonpath="{.items[0].metadata.name}")

 kubectl -n kubernetes-dashboard port-forward $POD_NAME 8443:8443

 ```

## Access from local
```
https://127.0.0.1:8443/
```


## Uninstall/Delete
```
helm delete kubernetes-dashboard -n kubernetes-dashboard
```