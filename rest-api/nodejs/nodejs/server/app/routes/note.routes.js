module.exports = (app) => {
    const notes = require('../controllers/note.controllers.js');

    // Create a new note.
    app.post('/notes', notes.create);

    // Retrieve all notes.
    app.get('/notes', notes.findAll);

    // Retrieve a single note wit a noteId.
    app.get('/notes/:noteId', notes.findOne);

    // Update a note with a noteId.
    app.put('/notes/:noteId', notes.update);

    // Delete a note with a noteId.
    app.delete('/notes/:noteId', notes.delete);
}
