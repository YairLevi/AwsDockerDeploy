const express = require('express')
const app = express()

app.use(express.json())

app.get('/', (req, res) => {
    res.send('Hello World, once again 5!')
})

app.post('/', (req, res) => {
    const { body } = req
    console.log(body)
    res.status(200).json({
        success: true,
        content: body
    })
})

app.listen(8000)