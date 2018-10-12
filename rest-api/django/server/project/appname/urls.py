from django.conf.urls import url
from rest_framework_mongoengine import routers as merouters
from project.appname.views import NoteViewSet
 
merouter = merouters.DefaultRouter()
merouter.register("", NoteViewSet)
 
urlpatterns = [
 
]
 
urlpatterns += merouter.urls
