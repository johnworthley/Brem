import BREMICOContract from "../../../../build/contracts/BREMICO.json";
import store from "../../../store";
import axios from "axios";

const contract = require("truffle-contract");

export function approveWithdraw(contractAddress) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const ico = contract(BREMICOContract);
      ico.setProvider(web3.currentProvider);

      web3.eth.getCoinbase((error, coinbase) => {
        if (error) {
          console.error(error);
        }

        ico.at(contractAddress).then(instance => {
          ico.isAuditor(coinbase).then(res => {
            if (!res) {
              return alert(
                "Error, coinbase is not current ICO auditor address"
              );
            }

            ico.confirmWithdraw({ from: coinbase }).then(res => {
              // web3.eth.balance(res => {
              //   if (res) {
              //     axios
              //       .post("/ico/withdrawn", {
              //         address: contractAddress
              //       })
              //       .then(res => console.log(res))
              //       .catch(err => console.log(err));
              //   }
              // });
              return alert("TX: " + res.tx);
            });
          });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}
