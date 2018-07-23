import React, { Component } from "react";
import store from "../../store";
import BREMContract from "../../../build/contracts/BREM.json";

const contract = require("truffle-contract");

class MarketplaceForm extends Component {
  constructor(props) {
    super(props);

    const web3 = store.getState().web3.web3Instance;

    if (typeof web3 !== "undefined" && web3 !== null) {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);
      brem.deployed().then(instance => {
        // Get projects amount
        instance.projectsAmount().then(projectsAmount => {
          this.setState({ amount: projectsAmount.toNumber() });
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  render() {
    return (
      <div>
        <span> This is Marketplace form</span>
        {this.state && this.state.amount && <label>TODO: BREM ICO list </label>}
      </div>
    );
  }
}

export default MarketplaceForm;
