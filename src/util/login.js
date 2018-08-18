import getWebThree from './getweb3'
import store from 'Store'
import config from 'Config'
import contract from 'truffle-contract'
import axios from 'axios'
import BREMContract from "../../build/contracts/BREM.json"

export default async () => {
  console.log('login started')
  const { host } = config
  await getWebThree()
  const { web3Instance, web3Account } = store
  const { currentProvider, utils, eth } = web3Instance
  const brem = contract(BREMContract);
  brem.setProvider(currentProvider)
  const coinbase = await eth.getCoinbase()
  store.update({
    web3Coinbase: coinbase
  })

  const name = "usernameZZZ"

  currentProvider.sendAsync({
      method: "personal_sign",
      params: [utils.utf8ToHex(coinbase), coinbase],
      from: coinbase
    }, async (err, res) => {
      if(err) return console.error(err)
      const { result: sign } = res
      const bremInstance = await brem.deployed()

      const signup = async name => {
        // Checking for valid username
        const isValidUsername = await bremInstance.isValidUsername(name)
        if (!isValidUsername) {
          console.log("Not valid username")
          return
        }
        // Sign up and login
        bremInstance.signUp(name, {from: coinbase}).then(txRes => {
          alert(txRes.tx) // Show tx hash
          return txRes.receipt.status === "0x1" // Return tx status
        })
        .catch(err => {
          console.error(err)
        })
      }

      const loginSuperAuditor = async function() {
          return await bremInstance.login({from: coinbase})
      }
        
      const loginDeveloper = async developer => {
        axios.post(host + "login", developer)
        .then()
      }

      const signupDeveloper = async developer => {
        axios.post(host + "signup", developer)
        .then()
      }

      const isSuperUser = await bremInstance.isSuperuser(coinbase)
      const isAuditor = await bremInstance.isAuditor(coinbase)
      if(isSuperUser || isAuditor) {
        // Checking for sign up
        const isSignUp = await bremInstance.isSignUp({from: coinbase})

        if (!isSignUp) {
          const isSuccess = await signup(name)
          if (!isSuccess) {
            alert("Ошибка записи в блокчейн")
            return
          }
        }
       
        const username = await loginSuperAuditor()
        console.log(username)

        // Write cookies
        axios.post(host + "session", {
          address: coinbase,
          sign: sign
        })
        .then(res => console.log(res))
        .catch(err => console.error(err))

      } else {
        let developer = {
          address: coinbase,
          sign: sign
        }
        // Login developer
        axios.post(host + "login", developer)
        .then(res => {
          developer = res.data
          console.log(developer)
        })
        .catch(err => {
          // Sign up developer
          developer.username = name
          axios.post(host + "signup", developer)
          .then(res => {
            console.log(res)
            axios.post(host + "login", developer)
            .then(res => {
              developer = res.data
              console.log(developer)
            })
          })
          .catch(err => console.log(err))
        })
        
      }
       
    })

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
