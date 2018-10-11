# REST API

This subdirectory contains different versions of the same Notes API.

This is intended to serve as a crash course in web development.

CRUD (create, read, update, delete) operations are the basis for web
applications and as such serve as a perfect example to see how these work.

We will build upon these principles to make more complex, resilient, and
interesting applications.

The applications follow the 
[model-view-controller design pattern](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller),
each in its on way.

The model defines the type of data we will be working with. 
In our case, we will create notes which contain title and content areas.

The view part of the design is what the client sees 
(the response created by our application).

The controller part is what will take our model to create, update, retrieve, or
delete a note.


## NodeJS

We use [express js](https://github.com/expressjs/express) and
[Mongoose](https://github.com/Automattic/mongoose).

The [entrypoint to the server](./nodejs/server/server.js) connects our
application to an instance of mongodb, it defines the routes that it will
serve, and it starts and binds the server to a given port.

The application is organized as follows:
* [routes](./nodejs/server/app/routes/note.routes.js) specifies how requests will be
handled. If I make a POST request to the url "http://domainname.com/notes" what
function will handle the request.

* [models](./nodejs/server/app/models/note.models.js) here, we defined our data
model to tell the api what kind of data we will be working with; what types of
notes we wwill be storing and retrieving from our database.

* [controllers](./nodejs/server/app/controllers/note.controllers.js) the
controllers will utilize our data model and take the request passed on from the 
routes to generate a response for our client.


## Flask

The flask application uses the `flask-restful` library.
Alternatively, the readme has a reference blog post on how one can use pure
flask to build a restful api.

The flask app is contained all in one file.

The [api script](./flask/server/api.py) will run our app and register the routes
to be handled by the object `NotesAPI`, which in turn specifies the
"controllers" by memeber functions.
