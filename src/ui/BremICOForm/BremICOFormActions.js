import ICOContract from "../../../build/contracts/BREMICO.json";
import store from "../../store";

const contract = require("truffle-contract");

export function buyTokens(contractAddress, weiAmount, form) {
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
                  // TODO: Set user's amount
                  // TODO: contract balance
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
