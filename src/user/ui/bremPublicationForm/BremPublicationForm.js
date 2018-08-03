import React, { Component } from "react";
import store from "../../../store";
import axios from "axios";
import { browserHistory } from "react-router";
import BREMContract from "../../../../build/contracts/BREM.json";
import ICOContract from "../../../../build/contracts/BREMICO.json";

const contract = require("truffle-contract");

class BremPublicationForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      address: props.value
    };

    this.icoAuditors.bind(this);

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const ico = contract(ICOContract);
      ico.setProvider(web3.currentProvider);

      ico.at(this.state.address).then(icoInstance => {
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

              const brem = contract(BREMContract);
              brem.setProvider(web3.currentProvider);

              brem.deployed().then(bremInstance => {
                bremInstance.login({ from: wallet }).then(username => {
                  this.setState({ username: username });
                });
              });
            });

            axios
              .get("http://127.0.0.1:8080/ico/audit", {
                params: {
                  address: this.state.address
                }
              })
              .then(res => {
                this.setState({ auditors: res.data });
              })
              .catch(err => console.log(err));
          }
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  icoAuditors() {
    const amount = this.state.auditors.length;
    if (amount === 0) {
      return (
        <div>
          <p>Auditors didn't add</p>
        </div>
      );
    }
    return this.state.auditors.map(auditor => (
      <p key={auditor.address}>{auditor.address}</p>
    ));
  }

  handleAddNewAuditorChange(e) {
    this.setState({ newAuditorAddress: e.target.value });
  }

  handleAddNewAuditor(e) {
    e.preventDefault();

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      const auditorAddress = this.state.newAuditorAddress;
      if (!web3.utils.isAddress(auditorAddress)) {
        return alert(auditorAddress + " is not Ethereum address");
      }

      this.props.onAddNewAuditorSubmit(this.state.address, auditorAddress);
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handlePublishProject(e) {
    e.preventDefault();
    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined" && web3 !== null) {
      this.props.onPublishProjectSubmit(this.state.address);
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  handleOpenICOPage(e) {
    browserHistory.push({
      pathname: "/ico/" + this.state.address
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
          this.state.username &&
          this.state.auditors && (
            <fieldset>
              <legend>{this.state.icoName}</legend>
              <p>{this.state.icoDescription}</p>
              <p>Wallet: {this.state.wallet} </p>
              <span className="pure-form-message">
                by {this.state.username}
              </span>
              <h5>Current Auditors</h5>
              {this.icoAuditors()}
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
