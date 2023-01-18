kubectl create ns ubic-service
kubectl label ns ubic-service ubic.net/name=ubic-service ubic.net=true

kubectl create ns ingress-nginx
kubectl label ns ingress-nginx app.kubernetes.io/name=ingress-nginx app.kubernetes.io/instance=ingress-nginx

kubectl create ns lease
kubectl label ns lease ubic.net=true