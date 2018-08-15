import BREMContract from "../../build/contracts/BREM.json";
import web3Provider from "../util/getweb3";
import axios from "axios";

const contract = require("truffle-contract");

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
            const host = "Здесь должна быть переменная с используменым сервером"
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

  export function createNewBREMICO(
  name,
  symbol,
  rate,
  cap,
  closingTime,
  description,
  files,
  image,
  form
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
                  if (form.state.price > 0) {
                    brmTokenInstance
                      .approve(brem.address, factoryPrice, { from: coinbase })
                      .then(res => {
                        bremInstance
                          .createBREMICO(
                            name,
                            symbol,
                            rate,
                            web3.utils.toWei(cap, "ether"),
                            closingTime.toString(),
                            docHash,
                            { from: coinbase }
                          )
                          .then(TXres => {
                            const ico = {
                              address: TXres.logs[0].args.icoAddress,
                              developer: {
                                address: coinbase
                              }
                            };
                            axios
                              .post("http://" + mHost + "/ico", ico)
                              .then(res => {
                                console.log(res);
                                let formData = new FormData();
                                formData.append(
                                  "address",
                                  TXres.logs[0].args.icoAddress
                                );
                                formData.append("image", image);
                                const config = {
                                  headers: {
                                    "content-type": "multipart/form-data"
                                  }
                                };
                                axios
                                  .post(
                                    "http://" + mHost + "/ico/image",
                                    formData,
                                    config
                                  )
                                  .then(res => {
                                    axios
                                      .get("http://" + mHost + "/ico/dev", {
                                        params: {
                                          address: coinbase
                                        }
                                      })
                                      .then(res => {
                                        console.log(res);

                                        form.setState({ devICOs: res.data });
                                        form.setState({ icoName: "" });
                                        form.setState({ icoSymbol: "" });
                                        form.setState({ icoRate: 0 });
                                        form.setState({ icoCap: 0 });
                                        form.setState({ icoDescription: "" });
                                      });
                                  });
                              })
                              .catch(err => console.log(err));
                            alert(
                              "TX: " +
                                TXres.tx +
                                " ICO: " +
                                TXres.logs[0].args.icoAddress +
                                " Token: " +
                                TXres.logs[0].args.tokenAddress
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
                  } else {
                    bremInstance
                      .createBREMICO(
                        name,
                        symbol,
                        rate,
                        web3.utils.toWei(cap, "ether"),
                        closingTime.toString(),
                        docHash,
                        { from: coinbase }
                      )
                      .then(TXres => {
                        const ico = {
                          address: TXres.logs[0].args.icoAddress,
                          developer: {
                            address: coinbase
                          }
                        };
                        axios
                          .post("http://" + mHost + "/ico", ico)
                          .then(res => {
                            console.log(res);
                            let formData = new FormData();
                            formData.append(
                              "address",
                              TXres.logs[0].args.icoAddress
                            );
                            formData.append("image", image);
                            const config = {
                              headers: {
                                "content-type": "multipart/form-data"
                              }
                            };
                            axios
                              .post(
                                "http://" + mHost + "/ico/image",
                                formData,
                                config
                              )
                              .then(res => {
                                axios
                                  .get("http://" + mHost + "/ico/dev", {
                                    params: {
                                      address: coinbase
                                    }
                                  })
                                  .then(res => {
                                    console.log(res);

                                    form.setState({ devICOs: res.data });
                                    form.setState({ icoName: "" });
                                    form.setState({ icoSymbol: "" });
                                    form.setState({ icoRate: 0 });
                                    form.setState({ icoCap: 0 });
                                    form.setState({ icoDescription: "" });
                                  });
                              });
                          })
                          .catch(err => console.log(err));
                        alert(
                          "TX: " +
                            TXres.tx +
                            " ICO: " +
                            TXres.logs[0].args.icoAddress +
                            " Token: " +
                            TXres.logs[0].args.tokenAddress
                        );
                      });
                  }
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



export function addNewAuditor(address, form) {
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
                .post("http://" + mHost + "/audit", auditor)
                .then(res => {
                  console.log(res);
                  form.setState({ newAuditorAddress: "" });
                  axios
                    .get("http://" + mHost + "/audit")
                    .then(res => {
                      form.setState({ auditors: res.data });
                      form.setState({ newAuditorAddress: "" });
                    })
                    .catch(err => {
                      console.error(err);
                    });
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

export function withdrawFees(withdrawAmount, form) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);

      // Get current ethereum wallet.
      web3.eth.getCoinbase((error, coinbase) => {
        // Log errors, if any.
        if (error) {
          console.error(error);
        }

        brem.deployed().then(function(instance) {
          instance
            .withdrawFees(web3.utils.toWei(withdrawAmount, "ether"), {
              from: coinbase
            })
            .then(res => {
              form.setState({ withdrawValue: 0 });
              web3.eth.getBalance(instance.address).then(balance => {
                form.setState({
                  bremBalance: web3.utils.fromWei(balance, "ether")
                });
              });
              return alert(res.tx);
            });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}

export function changeWithdrawFee(feePercent, form) {
  if (feePercent < 0 || feePercent > 100) {
    alert("WithdrawFeePercent must be between 0 and 100");
    return;
  }
  
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);

      // Get current ethereum wallet.
      web3.eth.getCoinbase((error, coinbase) => {
        // Log errors, if any.
        if (error) {
          console.error(error);
        }

        brem.deployed().then(function(instance) {
          instance.withdrawFeePercent().then(currentWithdrawFeePercent => {
            if (currentWithdrawFeePercent === feePercent) {
              alert("WithdrawFeePercent is already " + feePercent);
              return;
            }

            instance
              .setWithdrawFeePercent(feePercent, { from: coinbase })
              .then(res => {
                form.setState({ newWithdrawFeePercent: 0 });
                instance.withdrawFeePercent().then(fee => {
                  form.setState({ withdrawFeePercent: 0 });
                });
                return alert(res.tx);
              });
            });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}
