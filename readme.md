Financial Health Indicator (FHI) Web Application with Nginx Reverse Proxy and TLS

This repository contains the Financial Health Indicator (FHI) Web Application, a Go-based application designed to calculate financial health and savings projections. The application serves HTML pages with results and is containerized alongside a Dockerized Nginx reverse proxy, configured to handle TLS encryption using Let's Encrypt.

Features

Go Application:

Developed using Go to calculate financial health indicators and savings projections.

Renders HTML pages for results instead of JSON API responses.

Containerized with Docker for ease of deployment.

Nginx Reverse Proxy:

Acts as a reverse proxy to forward client requests to the Go application.

Configured for HTTPS using TLS certificates from Let's Encrypt.

TLS/SSL Encryption:

Ensures secure communication with the web server using a Let's Encrypt certificate.

Docker Compose:

Orchestrates the Go application and Nginx reverse proxy using a single command.

Production-Ready Configuration:

Optimized Nginx settings for security and performance.

Repository Structure

FHI
├── etc/                 # Configuration files (e.g., environment variables)
├── internal/            # Internal Go application logic
├── models/              # Go structs and models
├── nginx/               # Nginx configuration files
│   └── nginx.conf       # Nginx reverse proxy and SSL configuration
├── static/              # Static assets (CSS, JS, images)
├── templates/           # HTML templates served by the Go application
├── var/                 # Variable data for the app
├── .gitignore           # Files to ignore in Git
├── docker-compose.dev.yml # Docker Compose configuration for development
├── docker-compose.yml   # Docker Compose configuration for production
├── Dockerfile           # Dockerfile for the Go application
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
└── main.go              # Main application entry point

Prerequisites

Before starting, ensure you have the following installed:

Docker

Docker Compose

Go (for local development)

Setup Instructions

1. Clone the Repository

git clone https://github.com/your-repository-name/fhi.git
cd fhi

2. Replace Configuration Parameters

Update DNS settings in nginx/nginx.conf to match your domain:

server_name calculator.conectapro.tech www.calculator.conectapro.tech;

Replace Docker Hub account name in docker-compose.yml with your account:

app:
  image: your_dockerhub_username/fhi-app:latest

3. Build and Run Locally (Development)

For local development:

docker-compose -f docker-compose.dev.yml up --build

Access the Go application at: http://localhost:8080

Access the Nginx reverse proxy at: http://localhost

4. Production Deployment

Build and start the containers:

docker-compose up --build -d

Ensure the DNS record for your domain (e.g., calculator.conectapro.tech) points to the server's public IP address.

Test HTTPS:

curl -v https://calculator.conectapro.tech

Configuration

Environment Variables

Update environment variables in etc/ for customizing the application.

Nginx Configuration

The nginx/nginx.conf file includes:

Reverse proxy settings.

TLS/SSL configuration.

Example snippet:

server {
    listen 443 ssl;
    server_name calculator.conectapro.tech;

    ssl_certificate /etc/letsencrypt/live/calculator.conectapro.tech/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/calculator.conectapro.tech/privkey.pem;

    location / {
        proxy_pass http://app:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}

Docker Compose

docker-compose.yml orchestrates both the Go app and Nginx reverse proxy.

Example snippet:

services:
  app:
    image: your_dockerhub_username/fhi-app:latest
    expose:
      - "8080"
    networks:
      - app_network

  proxy:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    depends_on:
      - app
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf
      - ./etc/letsencrypt:/etc/letsencrypt
      - ./var/lib/letsencrypt:/var/lib/letsencrypt
    networks:
      - app_network

networks:
  app_network:
    driver: bridge

Testing

Test the Web Application

Access the homepage:

Visit http://localhost (development) or https://calculator.conectapro.tech (production).

Test the financial health and savings projection by submitting form data on the homepage. The application will render HTML pages with the results.

Troubleshooting

Nginx Not Forwarding Requests:

Check the nginx.conf file for correct proxy settings.

Ensure the app service is running and reachable within the Docker network.

TLS Certificate Issues:

Verify that the Let's Encrypt certificate files exist in /etc/letsencrypt.

Test renewal with:

certbot renew --dry-run

DNS Issues:

Verify the DNS A record points to the server's public IP.

Check DNS propagation using tools like whatsmydns.net.

