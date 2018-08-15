import BREMICOContract from "../../../../build/contracts/BREMICO.json";
import store from "../../../store";
import axios from "axios";
import mHost from "../../../config";

const contract = require("truffle-contract");

export function makeWithrawRequest(contractAddress, value, form) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function() {
      const ico = contract(BREMICOContract);
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
                      .withdraw(web3.utils.toWei(value, "ether"), {
                        from: coinbase
                      })
                      .then(resTX => {
                        axios
                          .put("http://" + mHost + "/ico/request", {
                            address: contractAddress
                          })
                          .then(res => {
                            console.log(res);
                            form.setState({ status: "requsted" });
                            form.setState({ withdrawValue: 0 });
                            form.setState({ requestedValue: value });
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
