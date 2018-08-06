import store from "../../store";

export function buyTokens(contractAddress, weiAmount, form) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      web3.eth.getCoinbase((err, coinbase) => {
        if (err) {
          console.error(err);
        }

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
            return alert("Success! TX: " + res.transactionHash);
          });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}
