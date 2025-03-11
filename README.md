# Financial Health Indicator (FHI) 

The Financial Health Indicator (FHI) is a web application built with Go and Docker, designed as a marketing hook for a fintech company. It calculates financial health and savings projections based on user inputs, offering a clear and intuitive overview of their financial situation.

https://github.com/user-attachments/assets/cf3ef6fb-1274-4cf1-9f6f-545f17514e2e


## Features

1. **Go Application**:
   - Developed using Go to calculate financial health indicators and savings projections.
   - Renders HTML pages for results according to outcome.
   - Containerized with Docker for ease of deployment.

2. **Nginx Reverse Proxy**:
   - Acts as a reverse proxy to forward client requests to the Go application.
   - Configured for HTTPS using TLS certificates from Let's Encrypt.

3. **TLS/SSL Encryption**:
   - Ensures secure communication with the web server using a Let's Encrypt certificate.

4. **Docker Compose**:
   - Orchestrates the Go application and Nginx reverse proxy using a single command.

5. **Production-Ready Configuration**:
   - Optimized Nginx settings for security and performance.


## Prerequisites

Before starting, ensure you have the following installed:

1. [Docker](https://www.docker.com/)
2. [Docker Compose](https://docs.docker.com/compose/)
3. [Go](https://golang.org/) (for local development)


## Contributing

### 1. Clone the Repository

```bash
git clone https://github.com/Juan-Alves-1/financial-health-calculator.git
cd fhi 
```


### 2. Replace Configuration Parameters

- Update **DNS** settings in `nginx/nginx.conf` to match your domain:

  ```nginx
  server_name yourwebsite.com www.yourwebsite.com;
  ```

### 3. Build and Run Locally (Development)

- For local development:

```bash
docker-compose -f docker-compose.dev.yml up --build
```

Access the Go application at: http://localhost:8080

Access the Nginx reverse proxy at: http://localhost

### 4. Production Deployment

- Build and start the containers:

```bash
docker-compose up --build -d
```

Ensure the DNS record for your domain (e.g., calculator.conectapro.tech) points to the server's public IP address.

Test HTTPS:

```bash
curl -v https://yourwebsite.com
```

## Configuration

### Environment Variables

Update environment variables in etc/ for customizing the application.

### Nginx Configuration

The nginx/nginx.conf file includes:

- Reverse proxy settings.

- TLS/SSL configuration.

Example snippet:

```bash
server {
    listen 443 ssl;
    server_name yourwebsite.com;

    ssl_certificate /etc/letsencrypt/live/yourwebsite.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/yourwebsite.com/privkey.pem;

    location / {
        proxy_pass http://app:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        }
    }
```

### Docker Compose

- docker-compose.yml orchestrates both the Go app and Nginx reverse proxy.

Example snippet:

```bash
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
```

### Testing

- Test the Web Application by accessing the homepage:

Visit http://localhost (development) or https://subdomain.yourwebsite.com (production).

Test the financial health and savings projection by submitting form data on the homepage. The application will render HTML pages with the results.

### Troubleshooting

- Nginx Not Forwarding Requests:

Check the nginx.conf file for correct proxy settings.

Ensure the app service is running and reachable within the Docker network.

- TLS Certificate Issues:

Verify that the Let's Encrypt certificate files exist in /etc/letsencrypt.

Test renewal with:
``` bash 
certbot renew --dry-run
```

- DNS Issues:

Verify the DNS A record points to the server's public IP.

Check DNS propagation using tools like whatsmydns.net.

