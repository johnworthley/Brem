import store from 'Store'
import Web3 from "web3"

export const WEB3_INITIALIZED = "WEB3_INITIALIZED"
function web3Initialized(results) {
  return {
    type: WEB3_INITIALIZED,
    payload: results
  }
}

let getWeb3 = () => new Promise(function(resolve, reject) {
  store.update({
    web3Status: {
      logged: 'pending'
    }
  })
  // Wait for loading completion to avoid race conditions with web3 injection timing.
  window.addEventListener('load', async () => {
    var results
    var web3 = window.web3

    // Checking if Web3 has been injected by the browser (Mist/MetaMask)
    if (typeof web3 !== "undefined") {
      // Use Mist/MetaMask's provider.
      web3 = new Web3(web3.currentProvider);

      results = {
        web3Instance: web3
      };

      console.log("Web3 is successfully injected.")
      const init = web3Initialized(results)
      const instance = init.payload.web3Instance
      const account = await instance.eth.getCoinbase()
      if(account) {
        const ethAmount = instance.utils.fromWei(await instance.eth.getBalance(account), "ether")
        resolve(store.update({
          web3Init: init,
          web3Instance: instance,
          web3Account: account,
          web3EthAmount: ethAmount,
          web3Status: {
            logged: true
          }
        }))
      }
      else resolve(store.update({
        web3Status: {
          logged: false
        }
      }))
    } else {
      // Fallback to localhost if no web3 injection. We've configured this to
      // use the development console's port by default.
      var provider = new Web3.providers.HttpProvider("http://127.0.0.1:9545")

      web3 = new Web3(provider)

      results = {
        web3Instance: web3
      }

      const init = web3Initialized(results)
      const instance = init.payload.web3Instance
      const account = await instance.eth.getCoinbase()
      resolve(store.update({
        web3Init: init,
        web3Instance: instance,
        web3Account: account
      }))
    }
  })
})

export default getWeb3
