import ICOContract from "../../../build/contracts/BREMICO.json";
import BREMTokenContract from "../../../build/contracts/BREMToken";
import store from "../../store";
import axios from "axios";
import mHost from "../../../config";

const contract = require("truffle-contract");

export function buyTokens(contractAddress, tokenAddress, weiAmount, form) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      web3.eth.getCoinbase((err, coinbase) => {
        if (err) {
          console.error(err);
        }

        const ico = contract(ICOContract);
        ico.setProvider(web3.currentProvider);
        ico.at(contractAddress).then(instance => {
          instance.hasClosed().then(hasClosed => {
            if (hasClosed) {
              return alert("Error, tokensale closed");
            }

            instance.auditSelected().then(auditSelected => {
              if (!auditSelected) {
                return alert("Error, tokensale didn't start");
              }

              // Buy tokens
              web3.eth
                .sendTransaction({
                  from: coinbase,
                  to: contractAddress,
                  value: web3.utils.toWei(weiAmount, "ether")
                })
                .then((res, err) => {
                  if (err) {
                    console.error(err);
                    return alert("Transaction error");
                  }
                  form.setState({ etherAmount: 0 });
                  form.setState({ tokensAmount: 0 });
                  instance.weiRaised().then(weiRaised => {
                    form.setState({ weiRaised: weiRaised.toNumber() });
                  });
                  web3.eth.getBalance(form.state.address).then(balance => {
                    form.setState({
                      balance: web3.utils.fromWei(balance, "ether")
                    });
                  });

                  instance.balances(coinbase).then(userBalance => {
                    form.setState({ userBalance: userBalance.toNumber() });
                  });

                  instance.balancesInToken(coinbase).then(userTotalTokens => {
                    form.setState({
                      userTotalTokens: userTotalTokens.toNumber()
                    });
                  });

                  const bremToken = contract(BREMTokenContract);
                  bremToken.setProvider(web3.currentProvider);
                  bremToken.at(tokenAddress).then(tokenInstance => {
                    tokenInstance
                      .balanceOf(coinbase)
                      .then(userCurrentBalance => {
                        form.setState({
                          userCurrentBalance: userCurrentBalance.toNumber()
                        });
                      });
                  });

                  return alert("Success! TX: " + res.transactionHash);
                });
            });
          });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}

export function refund(contractAddress, tokenAddress, form) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      web3.eth.getCoinbase((err, coinbase) => {
        if (err) {
          console.error(err);
        }

        const ico = contract(ICOContract);
        ico.setProvider(web3.currentProvider);
        ico.at(contractAddress).then(instance => {
          instance.hasClosed().then(hasClosed => {
            if (!hasClosed) {
              return alert("Error, tokensale closed");
            }

            instance.auditSelected().then(auditSelected => {
              if (!auditSelected) {
                return alert("Error, tokensale didn't start");
              }

              instance.capReached().then(capReached => {
                if (capReached) {
                  return alert("Error, cannot refund, ICO cap reached");
                }

                // Refund
                instance.refund({ from: coinbase }).then(res => {
                  web3.eth.getBalance(form.state.address).then(balance => {
                    form.setState({
                      balance: web3.utils.fromWei(balance, "ether")
                    });
                  });

                  instance.balances(coinbase).then(userBalance => {
                    form.setState({ userBalance: userBalance.toNumber() });
                  });

                  instance.balancesInToken(coinbase).then(userTotalTokens => {
                    form.setState({
                      userTotalTokens: userTotalTokens.toNumber()
                    });
                  });

                  const bremToken = contract(BREMTokenContract);
                  bremToken.setProvider(web3.currentProvider);
                  bremToken.at(tokenAddress).then(tokenInstance => {
                    tokenInstance
                      .balanceOf(coinbase)
                      .then(userCurrentBalance => {
                        form.setState({
                          userCurrentBalance: userCurrentBalance.toNumber()
                        });
                      });
                  });

                  return alert("Success! TX: " + res.tx);
                });
              });
            });
          });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}

// Superuser form
export function addNewAuditor(contractAddress, auditorAddress, form) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);

      web3.eth.getCoinbase((error, coinbase) => {
        if (error) {
          console.error(error);
        }

        ico.at(contractAddress).then(instance => {
          instance
            .addAuditor(auditorAddress, { from: coinbase })
            .then(txRes => {
              const data = {
                ico: {
                  address: contractAddress
                },
                auditor: {
                  address: auditorAddress
                }
              };
              axios
                .post("http://" + mHost + "/ico/audit", data)
                .then(res => {
                  console.log(res);
                  form.setState({ newAuditorAddress: "" });
                  axios
                    .get("http://" + mHost + "/ico/audit", {
                      params: {
                        address: contractAddress
                      }
                    })
                    .then(res => {
                      form.setState({ auditors: res.data });
                    })
                    .catch(err => console.log(err));
                })
                .catch(err => console.log(err));
              return alert("Success! TX: " + txRes.tx);
            });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}

export function publishProject(contractAddress, form) {
  let web3 = store.getState().web3.web3Instance;
  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    const ico = contract(ICOContract);
    ico.setProvider(web3.currentProvider);

    web3.eth.getCoinbase((err, coinbase) => {
      if (err) {
        console.error(err);
      }

      ico.at(contractAddress).then(instance => {
        instance.finishAuditorSelection({ from: coinbase }).then(txRes => {
          axios
            .put("http://" + mHost + "/ico/open", {
              address: contractAddress
            })
            .then(res => {
              console.log(res);
              form.setState({ auditSelected: true });
            })
            .catch(err => console.error(err));
          return alert("Success! TX: " + txRes.tx);
        });
      });
    });
  } else {
    console.error("Web3 is not initialized.");
  }
}

// Developer form
export function makeWithrawRequest(contractAddress, weiValue, form) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);

      web3.eth.getCoinbase((error, coinbase) => {
        if (error) {
          console.error(error);
          return;
        }

        ico.at(contractAddress).then(instance => {
          instance.wallet().then(wallet => {
            if (wallet !== coinbase) {
              return alert("Error, your address is not address of ICO wallet");
            }
            instance.isWithdrawn().then(isWithdrawn => {
              if (isWithdrawn) {
                return alert("Error, ICO is finished");
              }

              instance.hasClosed().then(hasClosed => {
                if (!hasClosed) {
                  return alert("Error, ICO didn't close");
                }

                instance.capReached().then(capReached => {
                  if (!capReached) {
                    return alert("Error, ICO failed");
                  }

                  instance.isRequested().then(isRequested => {
                    if (isRequested) {
                      return alert("You made requst already");
                    }

                    instance
                      .withdraw(weiValue, { from: coinbase })
                      .then(resTX => {
                        axios
                          .put("http://" + mHost + "/ico/request", {
                            address: contractAddress
                          })
                          .then(res => {
                            console.log(res);
                            form.setState({ isRequested: true });
                            instance.request().then(request => {
                              form.setState({
                                requestedValue: request[0].toNumber()
                              });
                            });
                          });
                        return alert("TX: " + resTX.tx);
                      });
                  });
                });
              });
            });
          });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}
