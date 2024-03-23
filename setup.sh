#!/usr/bin/env sh

# Create Cluster
kind create cluster --config=cluster.yaml

# Install Ingress Controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml

# Install ArgoCD
kubectl create namespace argocd
# Add ArgoCD Helm repository
helm repo add argo https://argoproj.github.io/argo-helm

kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d | xclip -selection c
