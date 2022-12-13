# Example-RestApi using Docker

A complete example service (NoteService) built with Golang.

In this example:

- How to create CRUD endpoint.
- How to use - Postgresql example.
- How to document API with Swagger-UI and OpenApi.

### Endpoints

#### HTML

| HTTP Method | URL                                    | Description     |
|-------------|----------------------------------------|-----------------|
| `GET`       | http://localhost:8080/docs/index.html/ | Swagger UI page |

#### Note Service

| HTTP Method | URL                            | Description       |
|-------------|--------------------------------|-------------------|
| `POST`      | http://localhost:8080/notes    | Create new Note   |
| `PUT`       | http://localhost:8080/note/:id | Update Note by ID |
| `GET`       | http://localhost:8080/note/:id | Get Note by ID    |
| `DELETE`    | http://localhost:8080/note/:id | Delete Note by ID |
| `GET`       | http://localhost:8080/notes    | Get All Notes     |



## Overview

This project is using the following go modules:

- [echo-swagger](https://github.com/swaggo/echo-swagger)

#### Build and run in Docker:

```
$ docker-compose up --build
```