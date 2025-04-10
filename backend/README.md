# computer-network-g28/backend

backend repository for Pic-me-pls Project

## Perequisite

- Golang ver go1.23.5 or newer https://go.dev/doc/install
- Docker https://docs.docker.com/desktop/setup/install/windows-install/
- air https://github.com/air-verse/air
- golangci-lint https://golangci-lint.run/welcome/install/
- pre-commit https://pre-commit.com/

## Run local server

1. clone repository

```
git clone https://github.com/CP-RektMart/computer-network-g28/backend
cd computer-network-g28/backend
```

2. run docker compose

```
docker-compose up -d
```

3. start server
   there are 2 ways:

- normal

```
make start
```

- with hot reload

```
air
```

4. install pre commit

```
pre-commit install
```
