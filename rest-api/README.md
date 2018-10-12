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

In flask, the data model is implicitly defined throughout the different request
handlers (the `NotesAPI` methods `get`, `put`, `delete`, etc.).

The routes are defined by a class, `NotesAPI` which handles all user requests.
SO this class (or `flask_restful.Resouce`) corresponds to the view.

The controller part of the api is also explicitly defined withint the request
hanlders through the `mongo.db.notes` object.


## Django

With Django the Model-View-Controller pattern is a lot more ovbious.
Django allows you to mess around with that part and does everything else ebhind
the scenes for you.

For this app you will need to create a model (describes what the data will look like),
a serializer (tells the app how to convert data to a python object and vice
versa), and a view (tells the app how to use the model and the serializer).

Once you have this parts you are done.

The thing with django is that it is extremely easy to set up, which is great.
But it is orders of magnitude slower than other alternatives.
As a matter of fact, django is pretty much the slowest thing you can build to
serve request but there are a lot of ways to get around this - service
replication, request caching, to name a few.
