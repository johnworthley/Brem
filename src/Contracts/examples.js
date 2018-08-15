import BREMContract from "../../build/contracts/BREM.json";
import web3Provider from "../util/getweb3";
import axios from "axios";

const contract = require("truffle-contract");

const host = "Здесь нужна константа для хоста"

export function signUpUser(name) {
  const web3 = web3Provider.getWeb3();
  // Using truffle-contract we create the authentication object.
  const brem = contract(BREMContract);
  brem.setProvider(web3.currentProvider);

  // Get current ethereum wallet.
  web3.eth.getCoinbase((error, coinbase) => {
    // Log errors, if any.
    if (error) {
      console.error(error);
      return
    }

    web3.currentProvider.sendAsync(
      {
        method: "personal_sign",
        params: [web3.utils.utf8ToHex(coinbase), coinbase],
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
    })
  }

  // Login
  export function loginUser() {
    const web3 = web3Provider.getWeb3();
    // Using truffle-contract we create the authentication object.
    const brem = contract(BREMContract);
    brem.setProvider(web3.currentProvider);
  
    // Get current ethereum wallet.
    web3.eth.getCoinbase((error, coinbase) => {
      if (error) {
        console.error(error);
        return
      }

      web3.currentProvider.sendAsync(
        {
          method: "personal_sign",
          params: [web3.utils.utf8ToHex(coinbase), coinbase],
          from: coinbase
        },
        (err, res) => {
          if (err) {
            console.error(err);
            return
          }
          const sign = res.result;

          brem.deployed().then(bremInstance => {
            // Attempt to login user.
            bremInstance.login({ from: coinbase })
            .then(function(userName) {
            // If no error, login user. (Superuser of auditor)
            // Write to cokkie address and sign
          })
          .catch(res => {
            const developer = {
              address: coinbase,
              sign: sign
            }
            axios.post(host + "/signup", developer)
             .then(res => {
               console.log(res); // В респонсе будет структура застройщика
             })
             .catch(err => {
               console.error(err);
             });
          })
        })
      });
    });
  };