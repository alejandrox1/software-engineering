from django.conf.urls import url
from rest_framework_mongoengine import routers as merouters
from project.appname.views import ToolViewSet
 
merouter = merouters.DefaultRouter()
merouter.register('mongo', ToolViewSet)
 
urlpatterns = [
 
]
 
urlpatterns += merouter.urls
