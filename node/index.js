const app = require('express')()

app.get('/', (req, res) => {
    res.send('Welcome to application of Node.js! ' + (new Date).toString())
})

app.listen(3000)