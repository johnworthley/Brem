import BREMContract from "../../../../build/contracts/BREM.json";
import { loginUser } from "../loginbutton/LoginButtonActions";
import store from "../../../store";
import axios from "axios";

const contract = require("truffle-contract");

export function signUpUser(name) {
  let web3 = store.getState().web3.web3Instance;

  // Double-check web3's status.
  if (typeof web3 !== "undefined") {
    return function(dispatch) {
      // Using truffle-contract we create the authentication object.
      const authentication = contract(BREMContract);

      authentication.setProvider(web3.currentProvider);

      // Declaring this for later so we can chain functions on Authentication.
      var authenticationInstance;

      // Get current ethereum wallet.
      web3.eth.getCoinbase((error, coinbase) => {
        // Log errors, if any.
        if (error) {
          console.error(error);
        }

        authentication.deployed().then(function(instance) {
          authenticationInstance = instance;

          // Attempt to sign up user.
          authenticationInstance
            .signUp(name, { from: coinbase })
            .then(function(result) {
              authenticationInstance.isSuperuser(coinbase).then(res => {
                if (!res) {
                  authenticationInstance.isAuditor(coinbase).then(res => {
                    if (!res) {
                      const developer = {
                        address: coinbase
                      };
                      axios
                        .post("http://127.0.0.1:8080/dev", developer)
                        .then(res => {
                          console.log(res);
                        })
                        .catch(err => {
                          console.error(err);
                        });
                    }
                  });
                }
              });
              // If no error, login user.
              return dispatch(loginUser());
            })
            .catch(function(result) {
              // If error...
            });
        });
      });
    };
  } else {
    console.error("Web3 is not initialized.");
  }
}
