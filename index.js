const express = require('express')
const app = express()

app.get('/', (req, res) => {
    res.send('Hello World, once again 2!')
})

app.listen(8000)