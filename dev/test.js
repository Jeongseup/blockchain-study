const Blockchain = require("./blockchain");

const bitcoin = new Blockchain();

const previouseBlockHash = '876546132454aDASDAKJADKAJXAKCLAJKL';

const currentBlockData = [{
        amount: 10,
        sender: 'akdajlksjakda',
        recipient: 'ASDAKJLD112'
    },
    {
        amount: 20,
        sender: 'akdajlksjakda',
        recipient: 'ASDAKJLD112'
    },
    {
        amount: 30,
        sender: 'akdajlksjakda',
        recipient: 'ASDAKJLD112'
    }
]

const nonce = 201202;
console.log(bitcoin.hashBlock(previouseBlockHash, currentBlockData, nonce))
// console.log(bitcoin.proofOfWork(previouseBlockHash, currentBlockData));

// bitcoin.createNewBlock(2389, 'OIUASPDASDPO', '78s97dxasdq54');

// bitcoin.createNewTransaction(789457, 'GAJSDPAOPD', 'PASDIP0OA');

// bitcoin.createNewBlock(213, 'JOEOGNAPSID', '78s97dxasdq55');

// bitcoin.createNewTransaction(100, 'GAJSDPAOPD', 'PASDIP0OA');
// bitcoin.createNewTransaction(200, 'GAJSDPAOPD', 'PASDIP0OA');
// bitcoin.createNewTransaction(300, 'GAJSDPAOPD', 'PASDIP0OA');

// bitcoin.createNewBlock(14738, 'POOPOPSDPO', '78s97dxasdq56');

// console.log(bitcoin);
// console.log(bitcoin.chain[1])