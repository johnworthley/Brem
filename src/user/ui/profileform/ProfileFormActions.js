import BRMTokenContract from '../../../../build/contracts/BRMToken.json'
import store from '../../../store'

const contract = require('truffle-contract')

export function mintBRMTokens(reciever, amount) {
    let web3 = store.getState().web3.web3Instance

    // Double-check web3's status.
    if (typeof web3 !== 'undefined') {

    return function(dispatch) {
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
