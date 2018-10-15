# [curl http://localhost:3000/notesBuild RESTful API in Go and MongoDB](https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94)


## Running the application                                                      
                                                                                
### Prerequisites                                                               
                                                                                
Make sure the you have an instance of MongoDB running in your local machine, if not, use the our
[containerized mongo](../mongodb/README.md)                                     
                                                                                
                                                                                
### Running the app

To run the application, you simply need to compile it.
Run
```
make shell
```

This will place you inside of a container and mount your current working
directory.

To compile the appication, do
```
go build
```

This will create an executable named as the directory you are in, `server`.
So to run it
```
./server
```


### Create a note
```
$ curl -H "Content-Type: application/json" -d '{"title":"first note", "content":"content of note"}' -X POST http://localhost:3000/notes
{"id":"5bc4f9a4a1ef9a4ddc98aaf1","title":"first note","content":"content of note"}
```


### List all notes
```
$ curl http://localhost:3000/notes
[{"id":"5bc4f9a4a1ef9a4ddc98aaf1","title":"first note","content":"content of note"}]
```

### List a specific note
```
$ curl http://localhost:3000/notes/5bc4ff24a1ef9a597606f246
{"id":"5bc4ff24a1ef9a597606f246","title":"first note","content":"content of note"}
```


### Update a note
```
$ curl -H "Content-Type: application/json" -d '{"title":"this is a note", "content":"hi"}' -X PUT http://localhost:3000/notes/5bc4f9a4a1ef9a4ddc98aaf1
{"result":"success"}
```

### Delete a note
```
$ curl -X DELETE http://localhost:3000/notes/5bc4f9a4a1ef9a4ddc98aaf1
{"result":"success"}
```
