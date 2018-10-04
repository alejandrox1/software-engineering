from django.shortcuts import render
 
from rest_framework_mongoengine import viewsets as meviewsets
from project.appname.serializers import ToolSerializer
from project.appname.models import Tool
 
class ToolViewSet(meviewsets.ModelViewSet):
    lookup_field = 'id'
    queryset = Tool.objects.all()
    serializer_class = ToolSerializer
