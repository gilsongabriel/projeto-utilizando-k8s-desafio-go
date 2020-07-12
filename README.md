# Desafio Golang - Utilizando K8S
## [Code Education](https://code.education) / [School Of Net](https://schoolofnet.com)

### Docker Hub
[gilsongabriel/projeto-utilizando-k8s-desafio-go](https://hub.docker.com/r/gilsongabriel/projeto-utilizando-k8s-desafio-go)

#### Porta web 8000

#### Docker Run

```
docker run -p 8000:8000 --name app_name gilsongabriel/projeto-utilizando-k8s-desafio-go
```

#### Docker Compose

```
version: "3.8"
services:
  app:
    image: gilsongabriel/projeto-utilizando-k8s-desafio-go
    container_name: app_name
    ports:
    - 8000:8000
    restart: always
    tty: true
```