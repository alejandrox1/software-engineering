from flask import Flask, abort, jsonify, request
from flask_restful import Api, Resource
from flask_pymongo import PyMongo
from bson.objectid import ObjectId


app = Flask(__name__)

app.config["MONGO_DBNAME"] = "notesdb"
app.config["MONGO_URI"] = "mongodb://{}:{}/{}".format(
    #"username",
    #"password",
    "localhost",
    "27017",
    "notesdb"
)
mongo = PyMongo(app)

api = Api(app)


class NotesAPI(Resource):
    def get(self, noteId=None):
        # Chose the notes collection from the database.
        notes = mongo.db.notes

        response = []
        # Retrieve all notes.
        if noteId is None:
            for n in notes.find():
                response.append({
                        "id": str(n["_id"]), 
                        "title": n["title"], 
                        "content": n["content"]})
            return jsonify(response)

        # Retrieve the note with id noteId.
        else:
            note = notes.find_one({"_id": ObjectId(noteId)})
            if note is not None:
                return jsonify({
                        "id": str(note["_id"]),
                        "title": note["title"],
                        "content": note["content"]
                        })
            else:
                abort(404)


    def post(self):
        # Chose the notes collection from the database.
        notes = mongo.db.notes

        title   = request.json["title"]
        content = request.json["content"]
        if title is None or content is None:
            abort(404)

        noteId = notes.insert({"title": title, "content": content})
        new_note = notes.find_one({"_id": ObjectId(noteId)})
        return jsonify({
                "id": str(new_note["_id"]),
                "title": new_note["title"],
                "content": new_note["content"]
                })


    def put(self, noteId=None):
        # Chose the notes collection from the database.
        notes = mongo.db.notes

        title   = request.json["title"]
        content = request.json["content"]
        if title is None and content is None:
            abort(404)

        note = notes.find_one({"_id": ObjectId(noteId)})
        if title is None:
            title = note["title"]
        if content is None:
            content = node["content"]

        notes.update_one(
            {"_id": ObjectId(noteId)},
            {
                "$set": {
                    "title": title,
                    "content": content
                }
            }
        )

        updated_note = notes.find_one({"_id": ObjectId(noteId)})
        return jsonify({
                "id": str(updated_note["_id"]),
                "title": updated_note["title"],
                "content": updated_note["content"]
                })


    def delete(self, noteId=None):
        # Chose the notes collection from the database.
        notes = mongo.db.notes

        deleted = notes.delete_one({"_id": ObjectId(noteId)})

        if deleted:
            return jsonify({"message": "{} successfully deleted!".format(noteId)})
        else:
            abort(404)


api.add_resource(NotesAPI, "/notes/", "/notes/<string:noteId>")


if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0")
