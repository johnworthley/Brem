import React, { Component } from "react";
import store from "../../store";
import ICOContract from "../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremICOForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      address: props.value
    };

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);
      ico.at(this.state.address).then(icoInstance => {
        icoInstance.name().then(icoName => {
          this.setState({ icoName: icoName });
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  render() {
    return (
      <div lassName="pure-u-1-1">
        {this.state && this.state.icoName && <p>{this.state.icoName}</p>}
        <p>BREM ICO Address: {this.state.address}</p>
      </div>
    );
  }
}

export default BremICOForm;
