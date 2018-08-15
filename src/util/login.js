import getWebThree from './getweb3'
import store from 'Store'
import config from 'Config'
import contract from 'truffle-contract'
import axios from 'axios'
import BREMContract from "../../build/contracts/BREM.json"

export default async () => {
  const { host } = config
  await getWebThree()
  const { web3Instance, web3Account } = store
  const { currentProvider, utils, eth } = web3Instance
  const brem = contract(BREMContract);
  brem.setProvider(currentProvider)
  const coinbase = await eth.getCoinbase()
  web3Instance.defaultAccount = web3Account
  // eth.sign('DATE', web3Account)
  console.log(web3Instance)


  currentProvider.sendAsync({
      method: "personal_sign",
      params: [utils.utf8ToHex(coinbase), coinbase],
      from: coinbase
    }, async (err, res) => {
      if(err) return console.log(err)
      const { result: sign } = res
      const bremInstance = await brem.deployed()


      const signup = async name => {
        bremInstance.signUp(name, {from: coinbase}).then(name => {
          console.log(name)
          return
        })
        .catch(err => {
          console.error(err)
        })
      }

      const login = async name => {

      }

        // Checking for superuser address
        const isSuperUser = await bremInstance.isSuperuser(coinbase)
        if(isSuperUser) {
          let name
          bremInstance.signUp(name, {from: coinbase}).then(name => {
            console.log(name)
            return
          })
          .catch(err => {
            console.error(err)
            signup(name)
          })
          return
        }


        const isAuditor = await bremInstance.isAuditor(coinbase)
        console.log(isSuperUser, isAuditor)
        return
        const developer = {
          address: coinbase,
          username: name,
          sign: sign
        }
        const signupRes = await axios.post(host + "/signup", developer)
    })
  // const req = await fetch(`http://${host}/login`)

}

/*
currentProvider.sendAsync(
  {
    method: "personal_sign",
    params: [utils.utf8ToHex(coinbase), coinbase],
    from: coinbase
  },
  (err, res) => {
    if (err) {
      console.error(err);
      return
    }
    const sign = res.result;

    brem.deployed().then(bremIntance => {
      // Checking fro superuser address
      bremIntance.isSuperuser(coinbase).then(isSuperuser => {
        if (isSuperuser) {
          // Try to login
          bremIntance.login({from: coinbase}).then(name => {
            // Name to login
            return
          })
          .catch(err => {
            // Sign up
            bremIntance.signUp(name, {from: coinbase}).then(name => {
              console.log(name)
              return
            })
            .catch(err => {
              console.error(err)
            })
          })
        } else {
          // Checking for auditor address
          bremIntance.isAuditor(coinbase).then(isAuditor => {
            if (isAuditor) {
              bremIntance.login({from: coinbase}.then(name => {
                console.log(name)
                return
              }))
              .catch(err => {
                bremIntance.signUp(name, {from: coinbase}).then(name => {
                  console.log(name)
                  return
                })
                .catch(err => {
                  console.error(err)
                })
              })
            } else {
             // Register developer
             const host = "Здесь должна быть переменная с используменым сервером"
             const developer = {
               address: coinbase,
               username: name,
               sign: sign
             }
             axios.post(host + "/signup", developer)
              .then(res => {
                console.log(res);
              })
              .catch(err => {
                console.error(err);
              });
            }
          })
        }
      })
    })
  })
*/
