import ICOContract from "../../../build/contracts/BREMICO.json";
import BREMTokenContract from "../../../build/contracts/BREMToken";
import store from "../../store";

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
                  web3.eth.getBalance(this.state.address).then(balance => {
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
