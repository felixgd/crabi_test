# crabi_test
## Description
This is a test for crabi

## How to run the API
```bash
make build
make run
```

## API Endpoints
POST User creation
```bash
curl --location 'localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "test",
    "first_name": "test",
    "last_name": "test"
}'
```
Response
```json
{
    "data": {
        "email": "test@test.com",
        "password": "",
        "first_name": "test",
        "last_name": "test",
        "PLD": {}
    }
}
```

POST User auth
```bash
curl --location 'localhost:8080/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "test@test.com",
  "password": "test"
}'
```
Response

```json
{
    "data": {
        "auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ0ZXN0QHRlc3QuY29tIiwiZXhwIjoxNjg5MzExMjU2fQ.PGcl9576jYBuIhaAF7Xt2oXTTAfEm96dp1dkHNZ8ZuI"
    }
}
```

Token needed in the `auth` header for getting user data

GET User Data
```bash
curl --location --request GET 'localhost:8080/user/:email' \
--header 'auth: <user_token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "joaquin@guzman.com",
    "password": "test",
    "first_name": "Joaquin",
    "last_name": "Guzman"
}'
```
Response
```json
{
    "data": {
        "email": "test@test.com",
        "password": "",
        "first_name": "test",
        "last_name": "test",
        "PLD": {}
    }
}
```
