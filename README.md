
# Users API with Clean Architecture Pattern using Go Lang

## Build

  make

## Run tests

  make test

## API requests 

### Add user

```
curl -X "POST" "http://localhost:8080/v1/user" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "name": "Buster",
  "email": "buster@gmail.com",
  "password": "123456",
  "address": "Lime Street 612",
  "age": 18
}'

```
### Search user

```
curl "http://localhost:8080/v1/user?name=buster" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Show users

```
curl "http://localhost:8080/v1/user" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```