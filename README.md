# link-storage-api
simple link storage rest api with chi router

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