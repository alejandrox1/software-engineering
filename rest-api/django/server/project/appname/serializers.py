from rest_framework_mongoengine import serializers
from project.appname.models import Note
 
class ToolSerializer(serializers.DocumentSerializer):
    class Meta:
        model = Note
        fields = '__all__'
