import React, { Component } from "react";
import store from "../../../store";
import axios from "axios";
import BREMContract from "../../../../build/contracts/BREM.json";
import BremPublicationFormContainer from "../bremPublicationForm/BremPublicationFormContainer";
import BremAuditFormContainer from "../bremAuditForm/BremAuditFormContrainer";
import BremDevFormContainer from "../bremDevForm/BremDevFormContainer";

const contract = require("truffle-contract");

class ProfileForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      BREMInstance: null,
      name: this.props.name,
      address: null
    };

    this.AuditorsList.bind(this);
    this.BremList.bind(this);
    this.AuditorICOList.bind(this);
    this.DevBremList.bind(this);

    this.init();
  }

  init() {
    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined") {
      const brem = contract(BREMContract);
      brem.setProvider(web3.currentProvider);
      brem.deployed().then(instance => {
        this.setState({ BREMInstance: instance });
        // Get coinbase
        web3.eth.getCoinbase((err, coinbase) => {
          if (err) {
            console.error(err);
          }
          this.setState({ address: coinbase });
          instance.isDeveloper(coinbase).then(res => {
            if (res) {
              this.setState({ role: "developer" });

              // Get developer's ICOs
              axios
                .get("http://127.0.0.1:8080/ico/dev", {
                  params: {
                    address: coinbase
                  }
                })
                .then(res => this.setState({ devICOs: res.data }))
                .catch(err => console.log(err));
            }
          });
          instance.isAuditor(coinbase).then(res => {
            if (res) {
              this.setState({ role: "auditor" });

              // Get current auditor ICOs
              axios
                .get("http://127.0.0.1:8080/audit/ico", {
                  params: {
                    address: coinbase
                  }
                })
                .then(res => {
                  this.setState({ auditorICOs: res.data });
                })
                .catch(err => console.log(err));
            }
          });
          instance.isSuperuser(coinbase).then(res => {
            if (res) {
              this.setState({ role: "superuser" });
              // Get all auditors' addresses
              axios
                .get("http://127.0.0.1:8080/audit")
                .then(res => {
                  this.setState({ auditors: res.data });
                })
                .catch(err => {
                  console.error(err);
                });
              // Get all new icos
              axios
                .get("http://127.0.0.1:8080/ico/created")
                .then(res => {
                  this.setState({ createdICOs: res.data });
                })
                .catch(err => console.error(err));
            }
          });

          instance.projectsAmount().then(projectsAmount => {
            this.setState({ icoAmount: projectsAmount });
          });
        });
      });
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  // Superuser
  handleMintAddressChange(event) {
    this.setState({ mintAddress: event.target.value });
  }

  handleMintAmountChange(event) {
    this.setState({ mintAmount: event.target.value });
  }

  handleMint(event) {
    event.preventDefault();

    const web3 = store.getState().web3.web3Instance;
    if (typeof web3 !== "undefined") {
      const recieverAddress = this.state.mintAddress;
      if (!web3.utils.isAddress(recieverAddress)) {
        return alert(recieverAddress + " is not Ethereum address");
      }

      const mintAmount = this.state.mintAmount;
      if (mintAmount <= 0) {
        return alert("Mint amount must be bigger than 0");
      }

      this.props.onMintFormSubmit(recieverAddress, mintAmount);
    }
  }

  handleNewAuditorChange(e) {
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

      this.props.onAddNewAuditorSubmit(address);
    } else {
      console.error("Web3 is not initialized.");
    }
  }

  AuditorsList() {
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

  // Superuser's project for publication

  BremList() {
    const amount = this.state.createdICOs.length;
    if (amount === 0) {
      return (
        <div>
          <h3>Brem ICOs not created yet...</h3>
        </div>
      );
    }

    return this.state.createdICOs.map(ico => (
      <BremPublicationFormContainer
        key={ico.address.toString()}
        value={ico.address}
      />
    ));
  }

  // Auditor form
  AuditorICOList() {
    const amount = this.state.auditorICOs.length;
    if (amount === 0) {
      return (
        <div>
          <h3>You don't have ICOs</h3>
        </div>
      );
    }

    return this.state.auditorICOs.map(ico => (
      <BremAuditFormContainer key={ico} value={ico} />
    ));
  }

  // Developer form
  handleICONameChange(e) {
    this.setState({ icoName: e.target.value });
  }

  handleICOTokenSymbolChange(e) {
    this.setState({ icoSymbol: e.target.value });
  }

  handleICORateChange(e) {
    this.setState({ icoRate: e.target.value });
  }

  handleICOCapChange(e) {
    this.setState({ icoCap: e.target.value });
  }

  handleICOClosingTimeChange(e) {
    let date = new Date(e.target.value);
    this.setState({ icoClosingTime: date.getTime() });
  }

  handleICODescriptionChange(e) {
    this.setState({ icoDescription: e.target.value });
  }

  handleICOFilesChange(e) {
    e.preventDefault();

    this.setState({ icoFiles: [] });
    for (let i = 0; i < e.target.files.length; i++) {
      let filePath = e.target.files[i].webkitRelativePath;

      let reader = new window.FileReader();
      reader.readAsArrayBuffer(e.target.files[i]);

      reader.onloadend = () => {
        this.setState({
          icoFiles: [
            ...this.state.icoFiles,
            {
              path: filePath,
              content: Buffer(reader.result)
            }
          ]
        });
      };
    }
  }

  _addDirectory(node) {
    if (node) {
      node.webkitdirectory = true;
    }
  }

  handleCreteBREMICO(e) {
    e.preventDefault();

    const name = this.state.icoName;
    if (name == null || name.length === 0) {
      return alert("BREM ICO name cannot be empty");
    }

    const symbol = this.state.icoSymbol;
    if (symbol == null || symbol.length !== 3) {
      return alert("BREM ICO Token symbol length must be 3");
    }

    const rate = this.state.icoRate;
    if (rate == null || rate <= 0) {
      return alert("BREM ICO rate must be more than 0");
    }

    const cap = this.state.icoCap;

    const closingTime = this.state.icoClosingTime;

    const description = this.state.icoDescription;
    if (description == null || description.length === 0) {
      return alert("BREM ICO description");
    }

    const files = this.state.icoFiles;
    this.props.onCreateBREMICOFormSubmit(
      name,
      symbol,
      rate,
      cap,
      closingTime,
      description,
      files
    );
  }

  DevBremList() {
    const amount = this.state.devICOs;
    if (amount === 0) {
      return (
        <div>
          <p>You didn't create ICOs</p>
        </div>
      );
    }
    return this.state.devICOs.map(ico => (
      <BremDevFormContainer key={ico} value={ico} />
    ));
  }

  render() {
    const SuperuserForm = (
      <div>
        <h3>Superuser</h3>

        <h4>BRM Token</h4>
        <form
          className="pure-form pure-form-ctacked"
          onSubmit={this.handleMint.bind(this)}
        >
          <fieldset>
            <legend>Mint BRM tokens</legend>
            <input
              id="mintAddress"
              type="text"
              onChange={this.handleMintAddressChange.bind(this)}
              placeholder="Address"
            />
            <input
              id="mintValue"
              type="number"
              onChange={this.handleMintAmountChange.bind(this)}
              placeholder="Amount"
            />

            <button type="submit" className="pure-button pure-button-primary">
              Mint
            </button>
            <span className="pure-form-message">
              To simplify the MVP, only manual BRM mint is available
            </span>
          </fieldset>
        </form>

        <h4>Auditors Managing</h4>
        <h5>All BREM auditos</h5>
        {this.state && this.state.auditors && this.AuditorsList()}
        <form
          className="pure-form pure-form-ctacked"
          onSubmit={this.handleAddNewAuditor.bind(this)}
        >
          <fieldset>
            <legend>Add auditor</legend>
            <input
              id="newAuditorAddress"
              type="text"
              onChange={this.handleNewAuditorChange.bind(this)}
              placeholder="New auditor address"
            />

            <button type="submit" className="pure-button pure-button-primary">
              Add
            </button>
          </fieldset>
        </form>

        <h4>Projects for publication</h4>
        {this.state && this.state.createdICOs && this.BremList()}
      </div>
    );

    const AuditorForm = (
      <div>
        <h3>Auditor</h3>

        <h4>My ICOs</h4>
        {this.state && this.state.auditorICOs && this.AuditorICOList()}
      </div>
    );

    const DeveloperForm = (
      <div>
        <h3>Developer</h3>

        <h4> BREM </h4>
        <form
          className="pure-form pure-form-ctacked"
          onSubmit={this.handleCreteBREMICO.bind(this)}
        >
          <fieldset>
            <legend>Create BREM ICO</legend>
            <p>
              <input
                id="name"
                type="text"
                onChange={this.handleICONameChange.bind(this)}
                placeholder="BREM ICO name"
              />
            </p>
            <p>
              <input
                id="symbol"
                type="text"
                onChange={this.handleICOTokenSymbolChange.bind(this)}
                placeholder="BREM ICO Token Symbol"
              />
            </p>
            <p>
              <input
                id="rate"
                type="number"
                onChange={this.handleICORateChange.bind(this)}
                placeholder="BREM ICO rate"
              />
            </p>
            <p>
              <input
                id="cap"
                type="number"
                onChange={this.handleICOCapChange.bind(this)}
                placeholder="BREM ICO cap"
              />
            </p>
            <p>
              <input
                id="closingTime"
                type="datetime-local"
                onChange={this.handleICOClosingTimeChange.bind(this)}
                placeholder="BREM ICO Closing Time"
              />
            </p>
            <p>
              <textarea
                id="description"
                rows="5"
                cols="70"
                onChange={this.handleICODescriptionChange.bind(this)}
                placeholder="BREM ICO Description"
                wrap="soft"
              />
            </p>
            <p>
              <label>
                Selection and deploying docs to SWARM will be later...
              </label>
            </p>
            <p>
              <input
                type="file"
                onChange={this.handleICOFilesChange.bind(this)}
                ref={node => this._addDirectory(node)}
              />
            </p>
            <button type="submit" className="pure-button pure-button-primary">
              Create new BREM ICO
            </button>
          </fieldset>
        </form>

        <h4>My ICOs</h4>
        {this.state && this.state.devICOs && this.DevBremList()}
      </div>
    );

    return (
      <div>
        <form className="pure-form pure-form-stacked">
          <fieldset>
            <label htmlFor="name">Username: {this.state.name}</label>
            <label htmlFor="address">Address: {this.state.address}</label>
          </fieldset>
        </form>

        {this.state &&
          this.state.role &&
          this.state.role === "superuser" &&
          SuperuserForm}

        {this.state &&
          this.state.role &&
          this.state.role === "developer" &&
          DeveloperForm}

        {this.state &&
          this.state.role &&
          this.state.role === "auditor" &&
          AuditorForm}
      </div>
    );
  }
}

export default ProfileForm;
