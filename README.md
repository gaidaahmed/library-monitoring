<div align="center">
📚 Library gRPC Server — Monitoring & Observability
🏫 Horizon University · SRE / Cloud Engineering








</div>
📌 Table of Contents
🎯 Project Overview

🛠️ Stack Used :

🐳 Docker Image
📦 Helm Deployment
📊 Metrics & Observability
📈 Prometheus Query
📂 Project Structure
👤 Author
✅ Deployment Status

🎯 1. Project Overview

This repository contains the monitoring and deployment of the Library gRPC Server using:
Docker (multi-stage image)
Helm (templated deployment)
Amazon EKS (cluster horizon2025)
Prometheus (metrics scraping)

The objective:
✔ Package the gRPC application
✔ Deploy it to Kubernetes
✔ Expose metrics
✔ Verify KPIs via Prometheus

🛠️ 2. Stack Used

Go – gRPC server with Prometheus middleware
Docker – multi-stage builds
Helm – deployment automation
Kubernetes (AWS EKS) – horizon2025 cluster
Prometheus – metrics scraping
AWS CLI + kubectl

🐳 3. Docker Image

Public image available at:
🔗 docker.io/ahmedgaida/library-server:monitoring
Image built & pushed using:
docker build -t ahmedgaida/library-server:monitoring .
docker push ahmedgaida/library-server:monitoring

📦 4. Helm Deployment

Folder: library-server-chart/
Deployment
kubectl create namespace sre
helm install library-server library-server-chart/ -n sre
Upgrade
helm upgrade library-server library-server-chart/ -n sre
Verify
kubectl get pods -n sre
kubectl get svc -n sre
Expected results:

| Component | Status                                      |
| --------- | ------------------------------------------- |
| Pod       | Running                                     |
| Service   | Exposes **50051/gRPC** and **2112/metrics** |

📊 5. Metrics & Observability

Forward metrics locally :
kubectl port-forward -n sre svc/library-server-library-server-chart 30112:2112
Then open:

👉 http://localhost:30112/metrics
Metrics exposed include:
go_gc_duration_seconds
grpc_server_handled_total
process_cpu_seconds_total
and more…

📈 6. Prometheus Query (Instructor Required)

Use this query in Prometheus:
sum(rate(grpc_server_handled_total[5m])) by (grpc_service, grpc_method)
Shows gRPC traffic grouped by service & method.

📂 7. Project Structure

library/
├── api/                     # gRPC generated files
├── client/                  # Demo gRPC client
├── server/                  # Prometheus-instrumented gRPC server
├── Dockerfile               # Multi-stage image
└── library-server-chart/    # Helm chart

👤 8. Author

Ahmed Gaida
Horizon University — Cloud / SRE / DevOps
Skills: Docker · Kubernetes · Prometheus · gRPC · Helm · AWS

✅ 9. Deployment Status

Successfully deployed to Horizon University EKS:
| Property  | Value                 |
| --------- | --------------------- |
| Cluster   | horizon2025           |
| Region    | eu-south-1            |
| Namespace | sre                   |
| Metrics   | Scraped by Prometheus |
