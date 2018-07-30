Install project fn to minikube

```
minikube start --vm-driver=kvm2 --cpus 3 --memory 8192
helm init
git clone git@github.com:fnproject/fn-helm.git && cd fn-helm
helm dep build fn
helm install --name my-release fn --set rbac.enabled=true
```

Install fn command line:

```
curl -LSs https://raw.githubusercontent.com/fnproject/cli/master/install | sh
```

Create a function:

```
fn init --runtime go add-vet
cd add-vet
cp $GOPATH/src/github.com/solo-io/gloo/example/projectfn/add-vet/* .
dep ensure
```

Deploy (replace FN_REGISTRY=soloio with your registry)
```
export POD_NAME=$(kubectl get pods --namespace default -l "app=my-release-fn,role=fn-service" -o jsonpath="{.items[0].metadata.name}")
kubectl port-forward --namespace default $POD_NAME 8080:80 &
export FN_API_URL=http://127.0.0.1:8080

FN_REGISTRY=soloio  fn deploy --app myapp
```

Make sure its available:
```
http://$FN_API_URL/r/myapp/add-vet
```

Test with

curl 'http://localhost:8080/r/myapp/add-vet' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H 'Content-Type: application/x-www-form-urlencoded' --data 'firstName=John&lastName=Smith&city=Cambridge&specialty=Dogs'


Install gloo and pet clinic:

```
glooctl install kube
kubectl apply -f $GOPATH/go/src/github.com/solo-io/gloo/example/projectfn/pet-clinic.yaml
```

Create a route to the pet clinic
```
glooctl route create --path-prefix / --upstream default-petclinic-80
```

Create a route that will add a new vet using the function we just created:
```
glooctl route create --path-exact /vets/new --upstream default-fn-release-fn-api-80  --function myapp:add-vet --http-method POST --sort
```

To complete the picture, you can also update the vets page to 'fix' a bug there:
```
kubectl apply -f $GOPATH/go/src/github.com/solo-io/gloo/example/projectfn/vets.yaml
glooctl route create --sort --path-exact /vets.html --upstream default-petclinic-vets-80
```

