# Schooglink Blogging App

## Overview

This is a blogging application built with Go (Golang) and MongoDB. It provides user authentication, allowing users to register, log in, create, read, update, and delete blog posts.

## Features

- User registration and login
- User profile management
- CRUD operations for blog posts
- Health check endpoint

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: MongoDB (MongoDB Atlas)
- **Web Framework**: Gin
- **Validation**: Go-playground/validator
- **Authentication**: JWT (JSON Web Tokens)

## Getting Started

### Prerequisites

- Go (version 1.22+)
- Docker and Docker Compose
- MongoDB Atlas account

### Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/rohanyh101/schooglink_assignment.git
   cd schooglink_assignment
   ```

2. **Configure environment variables**:
   - If you consider running the application locally, create a `.env` file in the root of your project and add the following variables accordingly:
   ```bash
   PORT = 8080
   MONGO_URI = 
   SECRET_KEY = 
   ENV = development
   ```

3. **Build and run the application**:
   If you consider running the application using docker, follow the commands below,

   1. Pull the image from the docker hub, [image_link](https://hub.docker.com/r/rohanyh/schooglink_blogapp)
   ```bash
   docker pull rohanyh/schooglink_blogapp
   ```
   
   2. Run the below cmd by inserting env variables during runtime
   ```bash
   docker run -e MONGO_URI="<mongodb_atlas_uri>" \
    -e SECRET_KEY="" \
    -e PORT="8080" \
    -e ENV=production \
    -p 8080:8080 \
    rohanyh/schooglink_blogapp
   ```
   
5. **Access the application**:
Your app will be running at `http://localhost:8080`. You can check the health endpoint at `http://localhost:8080/api/v1/health`.

### Endpoints

- Health Check: `GET /api/v1/health`
- User Registration: `POST /api/v1/auth/register`
- User Login: `POST /api/v1/auth/login`
- Get User Profile: `GET /api/v1/users/profile`
- Update User Profile: `PUT /api/v1/users/profile`
- Get All Posts: `GET /api/v1/posts/`
- Get Single Post: `GET /api/v1/posts/{post_id}`
- Create Post: `POST /api/v1/posts`
- Update Post: `PUT /api/v1/posts/{post_id}`
- Delete Post: `DELETE /api/v1/posts/{post_id}`

### Curl Commands
You can find all example curl commands in the `curl.http` file, which demonstrates how to test each endpoint.

### CURL commands And their responses
Don't worry these tokes contain dummy data

```bash
curl localhost:8080/api/v1/health | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    41  100    41    0     0    221      0 --:--:-- --:--:-- --:--:--   308
{
  "success": "server is up and running..."
}
❯ curl --location --request POST 'http://localhost:8080/api/v1/auth/login' \
 --header 'Content-Type: application/json' \
 --data-raw '{ "email": "rohan@gmail.com", "password": "1212" }' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   537  100   487  100    50     88      9  0:00:05  0:00:05 --:--:--   148
{
  "id": "6721c3a0b3d63417b86b54e4",
  "username": "rohan",
  "email": "rohan@gmail.com",
  "password": "$2a$15$o7ReuxCfjK03wB4LLMSTz.3Lv5u16TG32tS8PbeQa56wmn4ZtaUjq",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo",
  "created_at": "2024-10-30T05:26:56Z",
  "updated_at": "2024-10-30T10:10:16Z",
  "user_id": "6721c3a0b3d63417b86b54e4"
}
❯ curl --location --request PUT 'http://localhost:8080/api/v1/users/profile' \
 --header 'Content-Type: application/json' \
 --data-raw '{ "username": "rohanyh", "password": "1212" }' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    45    0     0  100    45      0    759 --:--:-- --:--:-- --:--:--  3461
100    84  100    39  100    45      7      8  0:00:05  0:00:05 --:--:--    10
{
  "message": "user updated successfully"
}
❯ curl --location --request GET 'http://localhost:8080/api/v1/users/profile' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    57  100    57    0     0   7668      0 --:--:-- --:--:-- --:--:-- 11400
100   492  100   492    0     0   8443      0 --:--:-- --:--:-- --:--:--  8443
{
  "id": "6721c3a0b3d63417b86b54e4",
  "username": "rohanyh",
  "email": "rohan@gmail.com",
  "password": "$2a$15$Sols/yqm79UKCR8W6shllO5ysH186qP.1eYKLwhY.E4LCwh8mlvgO",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo",
  "created_at": "2024-10-30T05:26:56Z",
  "updated_at": "2024-10-30T10:11:41.82Z",
  "user_id": "6721c3a0b3d63417b86b54e4"
}
❯ curl --location --request POST 'http://localhost:8080/api/v1/posts' \
 --header 'Content-Type: application/json' \
 --data-raw '{ "title": "demo", "content": "demo", "description": "demo" }' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    61    0     0  100    61      0    575 --:--:-- --:--:-- --:--:--   586
100   102  100    41  100    61    245    365 --:--:-- --:--:-- --:--:--   610
{
  "InsertedID": "672206f7f8af84bb9a97b921"
}
❯ curl --location --request GET 'http://localhost:8080/api/v1/posts/' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1388  100  1388    0     0  30851      0 --:--:-- --:--:-- --:--:-- 33047
[
  {
    "id": "6721c014e142f671118f94da",
    "user_id": "6721a6016284bc64e8e8dddd",
    "title": "demo",
    "content": "demo content",
    "description": "demo description",
    "author": "rohan",
    "created_at": "2024-10-30T05:11:48.339Z",
    "updated_at": "2024-10-30T05:28:42.27Z",
    "post_id": "6721c014e142f671118f94da"
  },
  {
    "id": "6721c3ceb3d63417b86b54e5",
    "user_id": "6721c3a0b3d63417b86b54e4",
    "title": "demo title",
    "content": "demo content",
    "description": "demo description",
    "author": "rohan",
    "created_at": "2024-10-30T05:27:42.622Z",
    "updated_at": "2024-10-30T05:27:42.622Z",
    "post_id": "6721c3ceb3d63417b86b54e5"
  }
]
❯ curl --location --request GET 'http://localhost:8080/api/v1/posts/672206f7f8af84bb9a97b921' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   257  100   257    0     0   5676      0 --:--:-- --:--:-- --:--:--  5840
{
  "id": "672206f7f8af84bb9a97b921",
  "user_id": "6721c3a0b3d63417b86b54e4",
  "title": "demo",
  "content": "demo",
  "description": "demo",
  "author": "rohan",
  "created_at": "2024-10-30T10:14:15.149Z",
  "updated_at": "2024-10-30T10:14:15.149Z",
  "post_id": "672206f7f8af84bb9a97b921"
}
❯ curl --location --request PUT 'http://localhost:8080/api/v1/posts/672206f7f8af84bb9a97b921' \
 --header 'Content-Type: application/json' \
 --data-raw '{ "title": "demo title" }' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    69  100    44  100    25    512    291 --:--:-- --:--:-- --:--:--   821
{
  "message": "blog post updated successfully"
}
❯ curl --location --request GET 'http://localhost:8080/api/v1/posts/672206f7f8af84bb9a97b921' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   263  100   263    0     0   4927      0 --:--:-- --:--:-- --:--:--  4962
{
  "id": "672206f7f8af84bb9a97b921",
  "user_id": "6721c3a0b3d63417b86b54e4",
  "title": "demo title",
  "content": "demo",
  "description": "demo",
  "author": "rohan",
  "created_at": "2024-10-30T10:14:15.149Z",
  "updated_at": "2024-10-30T10:17:22.812Z",
  "post_id": "672206f7f8af84bb9a97b921"
}
❯ curl --location --request GET 'http://localhost:8080/api/v1/user/posts/' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    53  100    53    0     0    713      0 --:--:-- --:--:-- --:--:-- 10600
100  1117  100  1117    0     0   9548      0 --:--:-- --:--:-- --:--:--  9548
[
  {
    "id": "6721c3ceb3d63417b86b54e5",
    "user_id": "6721c3a0b3d63417b86b54e4",
    "title": "demo title",
    "content": "demo content",
    "description": "demo description",
    "author": "rohan",
    "created_at": "2024-10-30T05:27:42.622Z",
    "updated_at": "2024-10-30T05:27:42.622Z",
    "post_id": "6721c3ceb3d63417b86b54e5"
  },
  {
    "id": "6721c3cfb3d63417b86b54e6",
    "user_id": "6721c3a0b3d63417b86b54e4",
    "title": "demo title",
    "content": "demo content",
    "description": "demo description",
    "author": "rohan",
    "created_at": "2024-10-30T05:27:43.978Z",
    "updated_at": "2024-10-30T05:27:43.978Z",
    "post_id": "6721c3cfb3d63417b86b54e6"
  }
]
❯ curl --location --request DELETE 'http://localhost:8080/api/v1/posts/672206f7f8af84bb9a97b921' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    44  100    44    0     0    512      0 --:--:-- --:--:-- --:--:--   523
{
  "message": "blog post deleted successfully"
}
❯ curl --location --request GET 'http://localhost:8080/api/v1/posts/672206f7f8af84bb9a97b921' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjcyMWMzYTBiM2Q2MzQxN2I4NmI1NGU0IiwidXNlcm5hbWUiOiJyb2hhbiIsImVtYWlsIjoicm9oYW5AZ21haWwuY29tIiwiZXhwIjoxNzMwMzY5NDE2fQ.SL3JaRvQFNJR2ZmZUFoV6vH7_7VUc5nI_YJZzphgUPo' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    41  100    41    0     0    928      0 --:--:-- --:--:-- --:--:--   976
{
  "error": "mongo: no documents in result"
}
```

### Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
