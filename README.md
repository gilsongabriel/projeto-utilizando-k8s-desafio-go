# Desafio Golang - Utilizando K8S
## [Code Education](https://code.education) / [School Of Net](https://schoolofnet.com)

### Docker Hub
[gilsongabriel/desafio-golang-esquenta-maratona-fullcycle](https://hub.docker.com/r/gilsongabriel/desafio-golang-esquenta-maratona-fullcycle)

#### Porta web 8000

#### Docker Run

```
docker run -p 8000:8000 --name app_name gilsongabriel/desafio-golang-esquenta-maratona-fullcycle
```

#### Docker Compose

```
version: "3.8"
services:
  app:
    image: gilsongabriel/desafio-golang-esquenta-maratona-fullcycle
    container_name: app_name
    ports:
    - 8000:8000
    restart: always
    tty: true
```