const express = require('express');
const app = express();
const port = 3000;

// Middleware to parse JSON bodies
app.use(express.json());

// GET route
app.get('/get', (req, res) => {
    res.send('This is a GET route');
});

// POST route 1
app.post('/post', (req, res) => {
    const data = req.body;
    res.send(`This is POST route 1 with data: ${JSON.stringify(data)}`);
});

// POST route 2
app.post('/postForm', (req, res) => {
    const data = req.body;
    console.log(req)
    res.send(JSON.stringify(req.body));
});

// Start the server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});