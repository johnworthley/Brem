import BRMTokenContract from '../../../../build/contracts/BRMToken.json'
import BREMContract from '../../../../build/contracts/BREM.json';
import store from '../../../store'

const contract = require('truffle-contract')

export function mintBRMTokens(reciever, amount) {
    let web3 = store.getState().web3.web3Instance

    // Double-check web3's status.
    if (typeof web3 !== 'undefined') {

    return function() {
        const brmToken = contract(BRMTokenContract)
        brmToken.setProvider(web3.currentProvider)

        // Declaring this for later so we can chain functions on Authentication.
        var brmTokenInstance

        // Get current ethereum wallet.
        web3.eth.getCoinbase((error, coinbase) => {
            // Log errors, if any.
            if (error) {
                console.error(error);
            }

            brmToken.deployed().then(function(instance) {
                brmTokenInstance = instance

                brmTokenInstance.mint(reciever, amount, {from: coinbase})
                .then(function(result) {
                    return alert('Success mint! TX: ' + result.tx);
                })
            .catch(function(result) {
                console.error(result);
            })
        })
        })
    }
    } else {
        console.error('Web3 is not initialized.');
    }
}

export function createNewBREMICO(name, symbol, rate, cap, closingTime, description) {
    let web3 = store.getState().web3.web3Instance

    // Double-check web3's status.
    if (typeof web3 !== 'undefined') {
        return function() {
            const brem = contract(BREMContract)
            brem.setProvider(web3.currentProvider)

            const brmToken = contract(BRMTokenContract)
            brmToken.setProvider(web3.currentProvider)

            var bremInstance, brmTokenInstance

            web3.eth.getCoinbase((error, coinbase) => {
                if (error) {
                    console.error(error)
                }

                brmToken.deployed().then((instance) => {
                    brmTokenInstance = instance

                    brem.deployed().then(function(instance) {
                        bremInstance = instance
    
                        // Get ICO creation price
                        bremInstance.icoCreationPrice({from: coinbase}).then((res) => {
                            const factoryPrice = res;
                            // Approve BRM token spend to BREM contract address
                            brmTokenInstance.approve(factoryPrice, brem.address, {from: coinbase}).then((res) => {
                                console.log(res);

                                bremInstance.createBREMICO(name, symbol, rate, cap, closingTime, description, [], {from: coinbase})
                                .then((res) => {
                                    // TODO: Sync with tokens approve
                                    console.log(res);
                                })
                                .catch((err) => {
                                    console.error(err);
                                    // Reject approving of transfering of tokens
                                    brmTokenInstance.decreaseApproval(brem.address, factoryPrice, {from: coinbase}).then((res) => {
                                        console.log(res);
                                    })
                                })
                            })
                        })
                    })
                })
                .catch((err) => {
                    console.error(err);
                })
            })
        }
    } else {
        console.error('Web3 is not initialized.');
    }
}
