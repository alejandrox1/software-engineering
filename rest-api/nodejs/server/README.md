# [Building a Restful CRUD API with Node.js, Express and MongoDB](https://www.callicoder.com/node-js-express-mongodb-restful-crud-api-tutorial/)

### Creating a note
```
$ curl -H "Content-Type: application/json" -d '{"title": "My first note", "content": "This is my first note in the EasyNotes app"}' http://localhost:3000/notes

{"_id":"5bb67eb6dabfe4004e3d174d","title":"My first note","content":"This is my first note in the EasyNotes app","createdAt":"2018-10-04T20:57:26.408Z","updatedAt":"2018-10-04T20:57:26.408Z","__v":0}
```

### Retrieving all notes
```
$ curl http://localhost:3000/notes

[{"_id":"5bb67eb6dabfe4004e3d174d","title":"My first note","content":"This is my first note in the EasyNotes app","createdAt":"2018-10-04T20:57:26.408Z","updatedAt":"2018-10-04T20:57:26.408Z","__v":0}]
```

### Retrieving a single note
```
$ curl http://localhost:3000/notes/5bb67eb6dabfe4004e3d174d

{"_id":"5bb67eb6dabfe4004e3d174d","title":"My first note","content":"This is my first note in the EasyNotes app","createdAt":"2018-10-04T20:57:26.408Z","updatedAt":"2018-10-04T20:57:26.408Z","__v":0}
```


### Updating a note
```
$ curl -H "Content-Type: application/json" -d '{"title": "My first note", "content": "Edited the content of this note"}' -X PUT http://localhost:3000/notes/5bb67eb6dabfe4004e3d174d

{"_id":"5bb67eb6dabfe4004e3d174d","title":"My first note","content":"Edited the content of this note","createdAt":"2018-10-04T20:57:26.408Z","updatedAt":"2018-10-04T21:02:51.365Z","__v":0}
```

### Delete a note
```
$ curl -X DELETE http://localhost:3000/notes/5bb67eb6dabfe4004e3d174d
{"message":"Note deleted successfully!"}
```
