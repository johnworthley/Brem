import React, { Component } from "react";
import store from "../../../store";
import { browserHistory } from "react-router";
import BREMContract from "../../../../build/contracts/BREM.json";
import ICOContract from "../../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremAuditForm extends Component {
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

        icoInstance.wallet().then(wallet => {
          this.setState({ wallet: wallet });

          const brem = contract(BREMContract);
          brem.setProvider(web3.currentProvider);

          brem.deployed().then(bremInstance => {
            bremInstance.login({ from: wallet }).then(username => {
              this.setState({ username: username });
            });
          });
        });

        // TODO: ICO status
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
          this.state.description &&
          this.state.username && (
            <fieldset>
              <legend>{this.state.name}</legend>
              <span className="pure-form-message">
                by {this.state.username}
              </span>
              <p>Address: {this.state.address}</p>
              <p>{this.state.description}</p>
              <p>Wallet: {this.state.wallet}</p>
              {/* TODO: ICO status and approve button */}

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

export default BremAuditForm;
