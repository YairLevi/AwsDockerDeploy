const express = require('express')
const app = express()

app.get('/', (req, res) => {
    res.send('Hello World, once again!')
})

app.listen(8000)