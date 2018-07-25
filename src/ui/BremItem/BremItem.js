import React, { Component } from "react";
import store from "../../store";
import BREMContract from "../../../build/contracts/BREM.json";
import ICOContract from "../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremItem extends Component {
  constructor(props) {
    super(props);

    this.state = {
      index: props.value
    };

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);
      brem.deployed().then(bremInstance => {
        bremInstance.getProject(this.state.index).then(icoAddress => {
          this.setState({ address: icoAddress });

          const ico = contract(ICOContract);
          ico.setProvider(web3.currentProvider);
          ico.at(this.state.address).then(icoInstance => {
            icoInstance.name().then(name => {
              this.setState({ name: name });
            });

            icoInstance.description().then(descr => {
              this.setState({ descr: descr });
            });

            icoInstance.wallet().then(developer => {
              bremInstance.login({ from: developer }).then(username => {
                this.setState({ user: username });
              });
            });
          });
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  render() {
    return (
      <div>
        <form>
          <fieldset>
            {this.state &&
              this.state.name && <legend>{this.state.name}</legend>}
            {this.state &&
              this.state.descr && <label> {this.state.descr}</label>}
            {this.state &&
              this.state.name && (
                <span className="pure-form-message">by {this.state.user} </span>
              )}
          </fieldset>
        </form>
      </div>
    );
  }
}

export default BremItem;
