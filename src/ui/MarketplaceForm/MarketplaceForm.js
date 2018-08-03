import React, { Component } from "react";
import axios from "axios";
import store from "../../store";
import BREMContract from "../../../build/contracts/BREM.json";
import BremItem from "../BremItem/BremItem";

const contract = require("truffle-contract");

class MarketplaceForm extends Component {
  constructor(props) {
    super(props);

    console.log(props);
    this.state = {
      status: props.value
    };

    this.BremList.bind(this);

    const web3 = store.getState().web3.web3Instance;

    if (typeof web3 !== "undefined" && web3 !== null) {
      switch (this.state.status) {
        case "opened":
          axios
            .get("http://127.0.0.1:8080/ico/opened")
            .then(res => this.setState({ icos: res.data }));
          break;
        case "success":
          axios
            .get("http://127.0.0.1:8080/ico/success")
            .then(res => this.setState({ icos: res.data }));
          break;
        case "failed":
          axios
            .get("http://127.0.0.1:8080/ico/failed")
            .then(res => this.setState({ icos: res.data }));
          break;
        case "withdrawn":
          axios
            .get("http://127.0.0.1:8080/ico/withdrawn")
            .then(res => this.setState({ icos: res.data }));
          break;
        default:
          this.setState({ invalidStaus: true });
      }
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

  BremList() {
    const amount = this.state.amount;
    if (amount === 0) {
      return (
        <div>
          <h3>Brem ICOs not created yet...</h3>
        </div>
      );
    }

    const indexes = Array.from(
      Array(amount),
      (val, index) => amount - index - 1
    );
    //Make list item components using indexes as keys
    const bremItems = indexes.map(index => (
      <BremItem key={index.toString()} value={index} />
    ));
    return <span>{bremItems}</span>;
  }

  render() {
    return (
      <div>
        <h2> This is Marketplace form</h2>
        {this.state && this.state.amount !== null && this.BremList()}
      </div>
    );
  }
}

export default MarketplaceForm;
