import BREMICOContract from "../../../../build/contracts/BREMICO.json";
import store from "../../../store";
import axios from "axios";
import mHost from "../../../config";

const contract = require("truffle-contract");

export function confirmWithdraw(contractAddress, form) {
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
          instance.isAuditor(coinbase).then(res => {
            if (!res) {
              return alert(
                "Error, coinbase is not current ICO auditor address"
              );
            }

            instance.confirmWithdraw({ from: coinbase }).then(res => {
              form.setState({ isConfirmed: true });
              instance.isRequested().then(isRequested => {
                if (!isRequested) {
                  instance.isWithdrawn().then(isWithdrawn => {
                    if (isWithdrawn) {
                      axios
                        .put("http://" + mHost + "/ico/withdrawn", {
                          address: contractAddress
                        })
                        .then(res => {
                          console.log(res);
                        })
                        .catch(err => console.error(err));
                    } else {
                      axios
                        .put("http://" + mHost + "/ico/success", {
                          address: contractAddress
                        })
                        .then(res => {
                          console.log(res);
                        })
                        .catch(err => console.error(err));
                    }
                  });
                }
              });
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
