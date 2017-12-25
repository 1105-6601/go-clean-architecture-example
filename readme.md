Clean architecture example in Golang.

# Vendoring

## install

```
dep ensure
```

# How to use

```
cp .env.default .env
```

## HTTP

```
cd src/infrastructure/ui/server/ 
go build -o server && ./server 
```

```
curl localhost:8080/v1/hello/1
```

## Command

```
cd src/infrastructure/ui/command
go build -o command && ./command -say-hello
```
