from django.shortcuts import render
 
from rest_framework_mongoengine import viewsets as meviewsets
from project.appname.serializers import ToolSerializer
from project.appname.models import Note
 
class NoteViewSet(meviewsets.ModelViewSet):
    lookup_field = 'id'
    queryset = Note.objects.all()
    serializer_class = ToolSerializer
