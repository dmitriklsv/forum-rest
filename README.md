# FORUM REST API

This project consists in creating an API for web forum that allows :

communication between users. associating categories to posts. liking and disliking posts and comments. filtering posts.

## Usage/Examples
In this example *forum:1.0* is the name of Docker image and *forum-ap* is the name of Docker container
Build Docker image based on Dockerfile with Makefile:

```bash
$ make build
```
Create Docker container:

```bash
$ make run
```

server started at http://localhost:8080/