const express = require('express')
const app = express()

app.get('/', (req, res) => {
    res.send('Hello World, once again 3!')
})

app.listen(8000)