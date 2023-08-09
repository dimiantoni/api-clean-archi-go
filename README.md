
## Users API with Clean Architecture Pattern using Go Lang

### Runnig application
```
git clone https://github.com/dimiantoni/api-clean-archi-go.git
cd api-clean-archi-go/

make dependencies
make build-mocks
make test

docker-compose build
docker-compose up -d

cp .env_example .env

make run-api
```

#### Run tests

  make test

#### API requests

#### Add user

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
#### Search user by email

```
curl "http://localhost:8080/v1/user?email=buster@gmail.com" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

#### Show users

```
curl "http://localhost:8080/v1/user" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

#### Delete user by MongoID

```
curl --request DELETE \
  --url http://localhost:8080/v1/user/64d310ac2a5593c7c9e27e44 \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data '
'
```