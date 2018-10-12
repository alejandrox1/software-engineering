from mongoengine import Document, EmbeddedDocument, fields
 
 
class Note(Document):
    title = fields.StringField(required=True)
    content = fields.StringField(required=True, null=True)
