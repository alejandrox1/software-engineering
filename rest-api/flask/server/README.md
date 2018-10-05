For another api pattern, not using `flask-restful`, see
[mongodb restful api with flask](https://www.bogotobogo.com/python/MongoDB_PyMongo/python_MongoDB_RESTAPI_with_Flask.php)

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
