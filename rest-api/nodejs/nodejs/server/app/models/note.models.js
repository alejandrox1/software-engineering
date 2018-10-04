const mongoose = require('mongoose');

const NoteSchema = mongoose.Schema({
    title: String,
    content: String
}, {
    // Create the fields createdAt and updatedAt.
    timestamps: true
});

module.exports = mongoose.model('Note', NoteSchema);
