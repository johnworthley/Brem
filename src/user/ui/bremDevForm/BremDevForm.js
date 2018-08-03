import React, { Component } from "react";
import store from "../../../store";
import { browserHistory } from "react-router";
import ICOContract from "../../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremDevForm extends Component {
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
        icoInstance.name().then(name => {
          this.setState({ name: name });
        });

        icoInstance.description().then(description => {
          this.setState({ description: description });
        });

        // TODO: ICO state
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleOpenICOPage() {
    browserHistory.push({
      pathname: "/ico/" + this.state.address
    });
  }

  render() {
    return (
      <div>
        {this.state &&
          this.state.name &&
          this.state.description && (
            <fieldset>
              <legend>{this.state.name}</legend>
              <p>Address: {this.state.addres}</p>
              <p>{this.state.description}</p>
              {/* TODO: ICO status and request button */}

              <p>
                <button
                  className="pure-button pure-button-primary"
                  onClick={this.handleOpenICOPage.bind(this)}
                >
                  Open ICO Page
                </button>
              </p>
            </fieldset>
          )}
      </div>
    );
  }
}

export default BremDevForm;
