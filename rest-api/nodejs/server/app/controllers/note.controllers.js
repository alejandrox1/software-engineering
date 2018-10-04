const Note = require('../models/note.models.js');


// Create and save a new note.
exports.create = (req, res) => {
    // Validate request.
    if (!req.body.content) {
        return res.status(400).send({
            message: "Note content cannot be empty"
        });
    }

    // Create a note.
    const note = new Note({
        title: req.body.title || "Untitle note",
        content: req.body.content
    });

    // Save note in database.
    note.save()
    .then(data => {
        res.send(data);
    }).catch(err => {
        res.status(500).send({
            message: err.message || "Some error occurred while creating the note"
        });
    });
};


// Retrieve and return all notes from the database.
exports.findAll = (req, res) => {
    Note.find()
    .then(notes => {
        res.send(notes);
    }).catch(err => {
        res.status(500).send({
            message: err.message || "Some error occurred while retrieving the notes"
        });
    });
};


// Find a note given its noteId.
exports.findOne = (req, res) => {
    Note.findById(req.params.noteId)
    .then(note => {
        if (!note) {
            return res.status(400).send({
                message: "Note could not be found with id " + req.params.noteId
            });
        }

        res.send(note);
    }).catch(err => {
        if (err.kid === 'ObjectId') {
            return res.status(404).send({
                message: "Note not found with id " + req.params.noteId
            });
        }

        return res.status(500).send({
            message: "Error retrieving note with id " + req.params.noteId
        });
    });
};


// Update a note given its noteId.
exports.update = (req, res) => {
    // Validate request.
    if (!req.body.content) {
        return res.status(400).send({
            message: "Note content cannot be empty"
        });
    }

    // Find the note and update it with the request body.
    Note.findByIdAndUpdate(req.params.noteId, {
        title: req.body.title || "Untitled note",
        content: req.body.content
    }, {new: true})
    .then(note => {
        if (!note) {
            return res.status(404).send({
                message: "Note not found with id " + req.params.noteId
            });
        }

        res.send(note);
    }).catch(err => {
        if (err.kind === 'ObjectId') {
            return res.status(404).send({
                message: "Note not found with id " + req.params.noteId
            });
        }

        return res.status(500).send({
            message: "Error updating note with id " + req.params.noteId
        });
    })
};


// Delete a note given its noteId.
exports.delete = (req, res) => {
    Note.findByIdAndRemove(req.params.noteId)
    .then(note => {
        if (!note) {
            return res.status(404).send({
                message: "Note not found with id " + req.params.noteId
            });
        }

        res.send({message: "Note deleted successfully!"});
    }).catch(err => {
        if (err.kind === 'ObjectId' || err.name === 'NotFound') {
            return res.status(404).send({
                message: "Note not found with id " + req.params.noteId
            });
        }

        return res.status(500).send({
            message: "Could not delete note with id " + req.params.noteId
        });
    });
};
