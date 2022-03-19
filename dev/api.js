var express = require('express')
var app = express()
const uuid = require('uuid').v1

// init
const nodeAddress = uuid().split('-').join('')
const Blockchain = require('./blockchain.js')
const bitcoin = new Blockchain()

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.get('/blockchain', function (req, res) {
    res.send(bitcoin)
})

app.get('/transaction', function (req, res) {
    res.send('Here is transaction')
})

app.post('/transaction', function (req, res) {
    const blockIndex = bitcoin.createNewTransaction(
        req.body.amount,
        req.body.sender,
        req.body.recipient
    )

    // res.send(`The amount of the transaction is ${req.body.amount}bitcoin`)
    res.json({
        note: `Transaction will be added in block ${blockIndex}`
    })
})

app.get('/mine', function (req, res) {
    const lastBlock = bitcoin.getLastBlock()
    const previousBlockHash = lastBlock['hash']
    const currentBlockData = {
        transactions: bitcoin.pendingTransactions,
        index: lastBlock['index'] + 1
    }

    const nonce = bitcoin.proofOfWork(previousBlockHash, currentBlockData)

    const blockHash = bitcoin.hashBlock(
        previousBlockHash,
        currentBlockData,
        nonce
    )

    bitcoin.createNewTransaction(12.5, '00', nodeAddress)
    const newBlock = bitcoin.createNewBlock(nonce, previousBlockHash, blockHash)

    res.json({
        note: 'New block mined successfully',
        block: newBlock
    })
})

app.listen(3000, function () {
    console.log('listening on port 3000...')
})
