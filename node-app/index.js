const app = require('express')()

app.get('/', (req, res) => {
    res.send('Welcome to application of Node.js!')
})

app.listen(3000)