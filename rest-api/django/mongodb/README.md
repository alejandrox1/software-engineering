# Running an instance of MongoDB in your local machine

To run an instance of MongoDB in your machine run the following commands
```
docker build -t mongo . 

docker run --rm --net host --name mongodb mongo mongod
```

Alternatively, you can spin up the container as such
```
docker run --rm --name mongodb --hostname mongodb -p 27017:27017 mongo mongod
```

If you use this last method you will have to change the 
[app configurations](../server/projecet/settings.py) to 
```python
mongoengine.connect(                                                            
    db="tools",                                                                 
    host="mongodb"                                                            
)
```
