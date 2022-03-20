#!/bin/bash
LOCAL_REPO='^localhost:5000'
CMD="docker"

if command -v podman &> /dev/null 
then
  CMD="podman"
fi

minikube start
minikube addons enable ingress
minikube addons enable registry
tmux new-window -n "minikube-registry"
tmux send-keys "kubectl port-forward --namespace kube-system service/registry 5000:80" "C-m"
"$CMD" images --format "{{.Repository}}:{{.Tag}}" | grep "$LOCAL_REPO" | xargs -I {} docker push {}
