# [Creating a Rest API/Webservice with Django Rest Framework and MONGO](http://fernandorodrigues.pro/creating-a-rest-apiwebservice-with-django-rest-framework-and-mongo-mongoengine-using-python-3/)


For more backround about the app, see 
[Django + MongoDB = Django REST Framework Mongoengine](https://medium.com/@vasjaforutube/django-mongodb-django-rest-framework-mongoengine-ee4eb5857b9a)


### Prerequisites                                                               
                                                                                
Make sure the you have an instance of MongoDB running in your local machine, if not, use the our
[containerized mongo](../mongodb/README.md)


## Running the app

To run the application, go inside a the container
```shell
make shell
```

then run the app by
```shell
python manage.py runserver
```


and visit [localhost:8000/notes](http://localhost:8000/notes/).

This will be a lot more interactive.
You can use curl to make request or you can do get and post request right from
your browser. 
Make sure to explore all the options!

You can browse notes by using their ids.

If, for example, you want to update a note and its id is `5bc11d00bacc1f00c4038ac0`
then try:
```shell
curl -H "Content-Type: application/json" -d '{"title": "another title", "content": "note"}' -X PUT http://localhost:8000/notes/5bc11d00bacc1f00c4038ac0/
```
