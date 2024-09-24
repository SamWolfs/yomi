#!/usr/bin/env sh

# Create Cluster
kind create cluster --config=cluster.yaml

# Install Ingress Controller
helm repo add cilium https://helm.cilium.io/
helm upgrade --install cilium cilium/cilium --version 1.16.1 \
    --namespace kube-system \
    --reuse-values \
    --set ingressController.enabled=true \
    --set ingressController.loadbalancerMode=dedicated
kubectl -n kube-system rollout restart deployment/cilium-operator
kubectl -n kube-system rollout restart ds/cilium

# Install ArgoCD
kubectl create namespace argocd
helm repo add argo https://argoproj.github.io/argo-helm
helm upgrade --install --wait -n argocd -f argocd/values.yaml argocd argo/argo-cd

argocd login --core
kubectl config set-context --current --namespace=argocd

# Install ArgoCD Image Updater
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj-labs/argocd-image-updater/stable/manifests/install.yaml
IMAGE_UPDATER_TOKEN=$(argocd account generate-token --account image-updater --id image-updater)
kubectl create secret generic argocd-image-updater-secret \
  --from-literal argocd.token="$IMAGE_UPDATER_TOKEN" --dry-run -o yaml | kubectl -n argocd apply -f -
kubectl -n argocd rollout restart deployment argocd-image-updater

kubectl port-forward svc/argocd-server -n argocd 8080:443
