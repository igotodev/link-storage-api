# link-storage-api

simple link storage rest api with chi router

## Authorization

All the methods in the API protected by JWT.

We must send `Authorization: Bearer <token>` in Header.

### POST

To create user and get authorization token we should use this request:

`POST /auth/sign-up`

and send username, password and email, like that:

```
{
    "username": "test2",
    "password": "test1234",
    "email": "test2@test.test"
}
```

If all fields are correct we will get 200 OK and response with new token:

```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Vary: Origin
Date: Thu, 24 Feb 2022 13:54:18 GMT
Content-Length: 128

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU3NTUyNDEsImlkIjozfQ.4qrgq_xkq7JsNvU8CdLc2zNTqIgBK5sZnS47-lgktXI"
}
```

If something went wrong we will get 404 Not Found or 400 Bad Request

To get authorization token we should use this request:

`POST /auth/sign-in`

and send username and password, like that:

```
{
    "username": "test2",
    "password": "test1234"
}
```

If user and password are correct we will get 200 OK and response with new token:

```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Vary: Origin
Date: Thu, 24 Feb 2022 13:57:21 GMT
Content-Length: 128

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU3NTUyNDEsImlkIjozfQ.4qrgq_xkq7JsNvU8CdLc2zNTqIgBK5sZnS47-lgktXI"
}
```

If username or password or request is not correct we will get 404 Not Found or 400 Bad Request

## Links

### GET

`GET /api/v1/link/` - get all links

Response:

```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Vary: Origin
Date: Sun, 13 Feb 2022 22:31:52 GMT
Content-Length: 248

[
  {
    "id": 1,
    "name": "Golang",
    "category": "Programming",
    "url": "https://golang.org",
    "date": "2022-02-10T16:33:58.714049Z"
  },
  {
    "id": 9,
    "name": "Microsoft",
    "category": "Computing and software",
    "url": "https://microsoft.com",
    "date": "2022-02-14T01:48:41.900768Z"
  }
]
```

`GET /api/v1/link/{id}` - get a link

Response:

```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Vary: Origin
Date: Sun, 13 Feb 2022 22:32:13 GMT
Content-Length: 114

{
  "id": 1,
  "name": "Golang",
  "category": "Programming",
  "url": "https://golang.org",
  "date": "2022-02-10T16:33:58.714049Z"
}
```

If request is not correct we will get 400 Bad Request or 404 Not Found with the message:

```
{
  "message": "bad request"
}
```

or

```
{
  "message": "not found"
}
```

### POST

`POST /api/v1/link/` - to add new link

We need to send:

```
{
  "name": "Microsoft",
  "category": "Computing and software",
  "url": "https://microsoft.com"
}
```

Response:

```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Vary: Origin
Date: Sun, 13 Feb 2022 22:36:27 GMT
Content-Length: 9

{
  "id": 9
}

```

If the database contains the same data we get a code 409 Conflict with message:

```
{
  "message": "conflict"
}
```

### PUT

`PUT /api/v1/link/{id}` - to update a link

```
{
  "name": "Microsoft",
  "category": "Computing, software and other same things",
  "url": "https://microsoft.com"
}
```

Response:

```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Vary: Origin
Date: Sun, 13 Feb 2022 22:38:01 GMT
Content-Length: 9

{
  "id": 9
}
```

### DELETE

`DELETE /api/v1/link/{id}` - to delete a link

Response:

```
HTTP/1.1 204 No Content
Access-Control-Allow-Origin: *
Content-Type: application/json
Vary: Origin
Date: Sun, 13 Feb 2022 22:40:11 GMT

<Response body is empty>
```