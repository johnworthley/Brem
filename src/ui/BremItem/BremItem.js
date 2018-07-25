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
            this.setState({ contract: icoInstance });

            icoInstance.name().then(name => {
              this.setState({ name: name });
            });

            icoInstance.description().then(descr => {
              this.setState({ descr: descr });
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
        {this.state &&
          this.state.name && <label>Project: {this.state.name}</label>}
        <br />
        {this.state &&
          this.state.descr && <label> Description: {this.state.descr}</label>}
      </div>
    );
  }
}

export default BremItem;
