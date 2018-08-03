import BRMTokenContract from "../../../../build/contracts/BRMToken.json";
import BREMContract from "../../../../build/contracts/BREM.json";
import store from "../../../store";
import axios from "axios";

import ipfs from "../../../ipfs";

const contract = require("truffle-contract");

export function mintBRMTokens(reciever, amount) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const brmToken = contract(BRMTokenContract);
      brmToken.setProvider(web3.currentProvider);

      // Declaring this for later so we can chain functions on Authentication.
      var brmTokenInstance;

      // Get current ethereum wallet.
      web3.eth.getCoinbase((error, coinbase) => {
        // Log errors, if any.
        if (error) {
          console.error(error);
        }

        brmToken.deployed().then(function(instance) {
          brmTokenInstance = instance;

          brmTokenInstance
            .mint(reciever, amount, { from: coinbase })
            .then(function(result) {
              return alert("Success mint! TX: " + result.tx);
            })
            .catch(function(result) {
              console.error(result);
            });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}

export function createNewBREMICO(
  name,
  symbol,
  rate,
  cap,
  closingTime,
  description,
  files
) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);

      const brmToken = contract(BRMTokenContract);
      brmToken.setProvider(web3.currentProvider);

      var bremInstance, brmTokenInstance, docHash;

      web3.eth.getCoinbase((error, coinbase) => {
        if (error) {
          console.error(error);
        }

        brmToken
          .deployed()
          .then(instance => {
            brmTokenInstance = instance;

            brem.deployed().then(function(instance) {
              bremInstance = instance;

              ipfs.files.add(files, (error, result) => {
                if (error) {
                  console.log(error);
                  return;
                }
                docHash = result[result.length - 1].hash;

                // Get ICO creation price
                bremInstance.icoCreationPrice({ from: coinbase }).then(res => {
                  const factoryPrice = res;

                  // Approve BRM token spend to BREM contract address
                  brmTokenInstance
                    .approve(brem.address, factoryPrice, { from: coinbase })
                    .then(res => {
                      bremInstance
                        .createBREMICO(
                          name,
                          symbol,
                          rate,
                          cap,
                          closingTime.toString(),
                          description,
                          docHash,
                          { from: coinbase }
                        )
                        .then(res => {
                          const ico = {
                            address: res.logs[0].args.icoAddress,
                            developer: {
                              address: coinbase
                            }
                          };
                          axios
                            .post("http://127.0.0.1:8080/ico", ico)
                            .then(res => console.log(res))
                            .catch(err => console.log(err));
                          return alert(
                            "TX: " +
                              res.tx +
                              " ICO: " +
                              res.logs[0].args.icoAddress +
                              " Token: " +
                              res.logs[0].args.tokenAddress
                          );
                        })
                        .catch(err => {
                          console.error(err);
                          // Reject approving of transfering of tokens
                          brmTokenInstance
                            .decreaseApproval(brem.address, factoryPrice, {
                              from: coinbase
                            })
                            .then(res => {
                              console.log(res);
                            });
                        });
                    });
                });
              });
            });
          })
          .catch(err => {
            console.error(err);
          });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}

export function addNewAuditor(address) {
  let web3 = store.getState().web3.web3Instance;

  if (typeof web3 !== "undefined") {
    return function() {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);

      web3.eth.getCoinbase((err, coinbase) => {
        if (err) {
          console.error(err);
          return;
        }

        brem.deployed().then(bremInstance => {
          // Check for existing of auditor
          bremInstance.isAuditor(address).then(res => {
            if (res) {
              alert(address + " is already auditor");
              return;
            }

            bremInstance.addAuditor(address, { from: coinbase }).then(txRes => {
              const auditor = {
                address: address
              };
              axios
                .post("http://127.0.0.1:8080/audit", auditor)
                .then(res => {
                  console.log(res);
                })
                .catch(err => {
                  console.error(err);
                });
              return alert("Success. TX: " + txRes.tx);
            });
          });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}
