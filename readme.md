Clean architecture example in Golang.

# Vendoring

## install

```
dep ensure
```

# Migration

## Execute

```
cd src/infrastructure/ui/command
go build -o command && ./command -migrate -step=1
```

## Sample data

```
insert into users (first_name, last_name, created_at) values ('Human', 'Ordinary', now());
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

# TODO

- Add mechanism of DI container
