const sha256 = require('sha256')

// 함수형 구현
function Blockchain() {
    this.chain = []
    this.pendingTransactions = []
    this.currentNodeUrl = currentNode
    this.networkNodes = []
    this.createNewBlock(0, '0', '0')
}

Blockchain.prototype.createNewBlock = function (
    nonce,
    previouseBlockHash,
    hash
) {
    const newBlock = {
        index: this.chain.length + 1,
        timestamp: Date.now(),
        transactions: this.pendingTransactions,
        nonce: nonce,
        hash: hash,
        previouseBlockHash: previouseBlockHash
    }

    // 초기화 : 새로운 블록을 만들 떄 우리는 새로운 트랜잭션들을 해당 블록에 추가해서
    this.pendingTransactions = []
    this.chain.push(newBlock)

    return newBlock
}

Blockchain.prototype.getLastBlock = function () {
    return this.chain[this.chain.length - 1]
}

Blockchain.prototype.createNewTransaction = function (
    amount,
    sender,
    recipient
) {
    const newTransaction = {
        amount: amount,
        sender: sender,
        recipient: recipient
    }

    this.pendingTransactions.push(newTransaction)
    return this.getLastBlock()['index'] + 1
}

Blockchain.prototype.hashBlock = function (
    previouseBlockHash,
    currentBlockData,
    nonce
) {
    const dataAsString =
        previouseBlockHash + nonce.toString() + JSON.stringify(currentBlockData)
    const hash = sha256(dataAsString)

    return hash
}

Blockchain.prototype.proofOfWork = function (
    previouseBlockHash,
    currentBlockData
) {
    // nonce값과 hash는 계속 바뀌어서 let 키워드 사용
    let nonce = 0
    let hash = this.hashBlock(previouseBlockHash, currentBlockData, nonce)

    while (hash.substring(0, 4) !== '0000') {
        nonce++
        hash = this.hashBlock(previouseBlockHash, currentBlockData, nonce)
        // console.log(hash);
    }

    return nonce
}

module.exports = Blockchain
