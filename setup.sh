#!/usr/bin/env sh

# Create Cluster
kind create cluster --config=cluster.yaml

# Install Ingress Controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml

# Install ArgoCD
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Setup ingress: https://argo-cd.readthedocs.io/en/stable/operator-manual/ingress/#kubernetesingress-nginx
kubectl apply -n argocd -f argocd/ingress.yaml

kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d && echo | xclip -selection c
