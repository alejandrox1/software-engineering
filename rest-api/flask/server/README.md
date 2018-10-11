# Building a RESTful CRUD API with Flask and MongoDB


For another api pattern, not using `flask-restful`, see
[mongodb restful api with flask](https://www.bogotobogo.com/python/MongoDB_PyMongo/python_MongoDB_RESTAPI_with_Flask.php)


### Prerequisites

Make sure the you have an instance of MongoDB running in your local machine, if not, use the our
[containerized mongo](../mongodb/README.md)


### Running the api server

This setup is mean to be run interactively.
As such, we will run the server from within a docker container.

First, you'll need to start the docker container.
Assuming, you have make installed in your machine, you can type
```shell
make shell
```

and then start the server with the following command
```
python api.py
```



## Interacting with the Notes API

The following curl commands show how we can interact with our API.

### Creating a note
```
$ curl -H "Content-Type: application/json" -d '{"title":"first note", "content":"content of note"}' -X POST http://localhost:5000/notes/
```

### Updating a note
```
$ curl -H "Content-Type: application/json" -d '{"title":"this is a note", "content":"ayaya"}' -X PUT http://localhost:5000/notes/5bb7885bbacc1f00fb784c97
{
  "content": "ayaya",
  "id": "5bb7885bbacc1f00fb784c97",
  "title": "this is a note"
}
```

### Deleting a note
```
$ curl -X DELETE http://localhost:5000/notes/5bb7885bbacc1f00fb784c97
{
  "message": "5bb7885bbacc1f00fb784c97 successfully deleted!"
}
```
