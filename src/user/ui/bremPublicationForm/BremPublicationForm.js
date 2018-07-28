import React, { Component } from "react";
import store from "../../../store";
import { browserHistory } from "react-router";
import BREMContract from "../../../../build/contracts/BREM.json";
import ICOContract from "../../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremPublicationForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      index: props.value
    };
    // TODO: Checking for closing time
    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);
      brem.deployed().then(bremInstance => {
        bremInstance.getProject(this.state.index).then(icoAddress => {
          this.setState({ icoAddress: icoAddress });
          const ico = contract(ICOContract);
          ico.setProvider(web3.currentProvider);

          ico.at(icoAddress).then(icoInstance => {
            icoInstance.auditSelected().then(res => {
              this.setState({ selected: res });
              if (!res) {
                icoInstance.name().then(icoName => {
                  this.setState({ icoName: icoName });
                });

                icoInstance.description().then(icoDescription => {
                  this.setState({ icoDescription: icoDescription });
                });

                icoInstance.wallet().then(wallet => {
                  this.setState({ wallet: wallet });
                  bremInstance.login({ from: wallet }).then(username => {
                    this.setState({ username: username });
                  });
                });

                // TODO: Docs
                // TODO: Auditors
              }
            });
          });
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleAddNewAuditorChange(e) {
    this.setState({ newAuditorAddress: e.target.value });
  }

  handleAddNewAuditor(e) {
    e.preventDefault();

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const address = this.state.newAuditorAddress;
      if (!web3.utils.isAddress(address)) {
        return alert(address + " is not Ethereum address");
      }

      this.props.onAddNewAuditorSubmit(this.state.icoAddress, address);
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handlePublishProject(e) {
    e.preventDefault();
    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      this.props.onPublishProjectSubmit(this.state.icoAddress);
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleOpenICOPage(e) {
    browserHistory.push({
      pathname: "/ico/" + this.state.icoAddress
    });
  }

  render() {
    return (
      <div>
        {this.state &&
          this.state.selected !== null &&
          !this.state.selected &&
          this.state.name !== null &&
          this.state.icoDescription &&
          this.state.wallet &&
          this.state.username && (
            <fieldset>
              <legend>{this.state.icoName}</legend>
              <p>{this.state.icoDescription}</p>
              <p>Wallet: {this.state.wallet} </p>
              <span className="pure-form-message">
                by {this.state.username}
              </span>
              <form
                className="pure-form pure-form-ctacked"
                onSubmit={this.handleAddNewAuditor.bind(this)}
              >
                <fieldset>
                  <legend>Add new auditor to project</legend>
                  <input
                    type="text"
                    onChange={this.handleAddNewAuditorChange.bind(this)}
                    placeholder="New Auditor address"
                  />

                  <button
                    type="submit"
                    className="pure-button pure-button-primary"
                  >
                    Add auditor to project
                  </button>
                </fieldset>
              </form>

              <button
                className="pure-button pure-button-primary"
                onClick={this.handlePublishProject.bind(this)}
              >
                Publish project
              </button>

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

export default BremPublicationForm;
