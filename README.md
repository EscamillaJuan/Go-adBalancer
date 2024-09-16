# Go-adBalancer

Load balancer using Golang

This repository contains a simple implementation of a load balancer and a backend server in Go. The backend server can be run on multiple ports to simulate multiple backend servers, while the load balancer distributes incoming requests to these backend servers.

## Setup the project

```bash
git clone https://github.com/EscamillaJuan/Go-adBalancer.git
cd Go-adBalancer
```

## Run backend server

```bash
cd server
go run main.go --port=XXXX
```

You can run multiple instances of the backend server on different ports. Open multiple terminal windows and execute the command in each terminal with a different port for each one:

## Run load balancer

Before running the load balancer, be sure to add the correct url server in the servers variable, then

```bash
cd src
go run main.go
```
