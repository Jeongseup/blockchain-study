const port = process.argv[2]
const currentNodeUrl = process.argv[3]

var express = require('express')
var app = express()
const uuid = require('uuid').v1
const rp = require('request-promise')

// init
const nodeAddress = uuid().split('-').join('')
const Blockchain = require('./blockchain.js')
const bitcoin = new Blockchain(currentNodeUrl)

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

// register a node and broadcast it the network
app.post('/register-and-broadcast-node', function (req, res) {
    const newNodeUrl = req.body.newNodeUrl

    if (bitcoin.networkNodes.indexOf(newNodeUrl) == -1) {
        bitcoin.networkNodes.push(newNodeUrl)
    }

    const regNodesPromises = []
    bitcoin.networkNodes.forEach((networkNodeUrl) => {
        const requestOptions = {
            uri: networkNodeUrl + '/register-node',
            method: 'POST',
            body: { newNodeUrl: newNodeUrl },
            json: true
        }
        regNodesPromises.push(rp(requestOptions))
    })

    Promise.all(regNodesPromises)
        .then((data) => {
            // use the data ...
            // broadcast가 끝났으면 bulk
            const bulkRegisterOptions = {
                uri: newNodeUrl + '/register-nodes-bulk',
                method: 'POST',
                body: {
                    allNetworkNodes: [
                        ...bitcoin.networkNodes,
                        bitcoin.currentNodeUrl
                    ]
                },
                json: true
            }
            return rp(bulkRegisterOptions)
        })
        .then((data) => {
            res.json({ note: 'New node registered with network successfully.' })
        })
})

// register a node with the network
app.post('/register-node', function (req, res) {
    const newNodeUrl = req.body.newNodeUrl
    // const nodeNotAlreadyPresent = bitcoin.networkNodes.indexOf(newNodeUrl) == -1
    // const notCurrentNode = bitcoin.currentNodeUrl !== newNodeUrl
    // if (nodeNotAlreadyPresent && notCurrentNode)
    //     bitcoin.networkNodes.push(newNodeUrl)
    // res.json({ note: 'New node registered successfully.' })
})

app.post('/register-nodes-bulk', function (req, res) {})

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

app.listen(port, function () {
    console.log(`Listening on port ${port}...`)
})
