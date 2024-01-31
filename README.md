# fiber-hexagonal

## config Environment

```bash
cp .env.example .env
```

- then update environments in .env

## Run Project With Docker

```bash
docker create network Heepoke
```

```bash
docker compose up -d && docker compose logs api --follow
```

## Clear

```bash
go mod tidy
```

## Generate Swagger

```bash
swag init -g cmd/main.go --output=internals/app/docs
```

## Local Swagger Doc Api

```bash
http://localhost:6476/apis/docs/index.html
```
