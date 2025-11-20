📚 Library gRPC Server – Monitoring Deployment
Horizon University – SRE / Observability Assignment

This repository contains the complete monitoring deployment of the Library gRPC Server, instrumented with Prometheus and deployed on the university’s Amazon EKS cluster using Helm.

🚀 Summary of the Work

This project demonstrates:

Building a Docker image for the Library gRPC server

Publishing the image to Docker Hub

Creating a complete Helm chart for deployment

Deploying the service to the Horizon University EKS cluster

Exposing gRPC and Prometheus metrics ports

Verifying and scraping metrics using Prometheus

🛠️ Technologies Used

Go (gRPC server + Prometheus middleware)

Docker (multi-stage build)

Kubernetes (Amazon EKS cluster: horizon2025)

Helm (deployment automation)

Prometheus (metrics scraping)

AWS CLI & kubectl

🐳 Docker Image

Public Docker image available at:
docker.io/ahmedgaida/library-server:monitoring

Commands used to build and push the image:

docker build -t ahmedgaida/library-server:monitoring .

docker push ahmedgaida/library-server:monitoring

📦 Helm Deployment

Helm chart directory: library-server-chart

Deployment steps:

kubectl create namespace sre

helm install library-server library-server-chart/ -n sre

To update: helm upgrade library-server library-server-chart/ -n sre

After deployment, verify with:

kubectl get pods -n sre

kubectl get svc -n sre

Expected result:

Pod in "Running" state

Service exposing ports: 50051 (gRPC) and 2112 (Prometheus metrics)

📊 Accessing Metrics

Forward the Prometheus metrics port locally:

kubectl port-forward -n sre svc/library-server-library-server-chart 30112:2112

Then open metrics in your browser or terminal at:
http://localhost:30112/metrics
You will see metrics such as:

go_gc_duration_seconds

grpc_server_handled_total

process_cpu_seconds_total

etc.

📈 Prometheus Query (for Instructor Verification)

Use the following query to verify the gRPC activity:

sum(rate(grpc_server_handled_total[5m])) by (grpc_service, grpc_method)

This displays the request rate per gRPC service and method.

📂 Project Structure Overview

library/
• api/ — Generated protobuf & gRPC files
• client/ — Minimal gRPC client
• server/ — gRPC server instrumented with Prometheus
• Dockerfile — Multi-stage Docker build
• library-server-chart/ — Helm chart

👤 Author

Ahmed Gaida
Horizon University — SLO / SRE Module
Skills: Docker • Kubernetes • Helm • AWS • Prometheus • gRPC

✅ Deployment Status

This project has been successfully deployed on the Horizon University EKS cluster:

Cluster: horizon2025

Region: eu-south-1

Namespace: sre

Metrics: exposed and scraped by Prometheus
