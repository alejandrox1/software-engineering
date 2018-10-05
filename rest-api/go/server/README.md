# [Build RESTful API in Go and MongoDB](https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94)


```
$ curl -H "Content-Type: application/json" -d '{"title":"first note", "content":"content of note"}' -X POST http://localhost:5000/notes/
```


```
$ curl -H "Content-Type: application/json" -d '{"title":"this is a note", "content":"ayaya"}' -X PUT http://localhost:5000/notes/5bb7885bbacc1f00fb784c97
{
  "content": "ayaya",
  "id": "5bb7885bbacc1f00fb784c97",
  "title": "this is a note"
}
```


```
$ curl -X DELETE http://localhost:5000/notes/5bb7885bbacc1f00fb784c97
{
  "message": "5bb7885bbacc1f00fb784c97 successfully deleted!"
}
```
