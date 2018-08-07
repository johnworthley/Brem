import BREMICOContract from "../../../../build/contracts/BREMICO.json";
import store from "../../../store";
import axios from "axios";

const contract = require("truffle-contract");

export function addNewAuditor(contractAddress, auditorAddress, form) {
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
                .post("http://127.0.0.1:8080/ico/audit", data)
                .then(res => {
                  console.log(res);
                  form.setState({ newAuditorAddress: "" });
                  axios
                    .get("http://127.0.0.1:8080/ico/audit", {
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
    const ico = contract(BREMICOContract);
    ico.setProvider(web3.currentProvider);

    web3.eth.getCoinbase((err, coinbase) => {
      if (err) {
        console.error(err);
      }

      ico.at(contractAddress).then(instance => {
        instance.finishAuditorSelection({ from: coinbase }).then(txRes => {
          axios
            .put("http://127.0.0.1:8080/ico/open", {
              address: contractAddress
            })
            .then(res => {
              console.log(res);
              form.setState({ visible: false });
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
