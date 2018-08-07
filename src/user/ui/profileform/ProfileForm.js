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
              this.setState({ icoName: "" });
              this.setState({ icoSymbol: "" });
              this.setState({ icoRate: 0 });
              this.setState({ icoCap: 0 });
              this.setState({ icoClosingTime: new Date(Date.now()) });
              this.setState({ icoDescription: "" });

              instance.icoCreationPrice().then(price => {
                this.setState({
                  price: web3.utils.fromWei(price.toNumber(), "ether")
                });
              });

              // Get withdrawFeePercent
              instance.withdrawFeePercent().then(withdrawFeePercent => {
                this.setState({
                  withdrawFeePercent: withdrawFeePercent.toNumber()
                });
              });

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
              this.setState({ withdrawValue: 0 });
              this.setState({ newICOCreationPrice: 0 });
              this.setState({ newWithdrawFeePercent: 0 });
              this.setState({ newAuditorAddress: "" });
              this.setState({ mintAddress: "" });
              this.setState({ mintAmount: 0 });

              // Get BREM contract balance
              web3.eth.getBalance(instance.address).then(bremBalance => {
                this.setState({
                  bremBalance: web3.utils.fromWei(bremBalance, "ether")
                });
              });

              // Get ICO creation price
              instance.icoCreationPrice().then(icoCreationPrice => {
                this.setState({
                  icoCreationPrice: web3.utils.fromWei(
                    icoCreationPrice.toNumber(),
                    "ether"
                  )
                });
              });
              // Get withdrawFeePercent
              instance.withdrawFeePercent().then(withdrawFeePercent => {
                this.setState({
                  withdrawFeePercent: withdrawFeePercent.toNumber()
                });
              });
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

      this.props.onMintFormSubmit(recieverAddress, mintAmount, this);
    }
  }

  handleWithdrawValueChange(event) {
    const withdrawValue = event.target.value;
    if (withdrawValue > this.state.bremBalance) {
      alert(
        "Error, withdraw value must be less than " +
          this.state.bremBalance +
          " Eth"
      );
      return;
    }
    this.setState({ withdrawValue: event.target.value });
  }

  handleWithdraw(event) {
    event.preventDefault();

    const withdrawAmount = this.state.withdrawValue;
    if (withdrawAmount <= 0 || withdrawAmount > this.state.bremBalance) {
      return alert("Mint amount must be bigger than 0");
    }

    this.props.onWithdrawFormSubmit(withdrawAmount, this);
  }

  handleCreationPriceChange(event) {
    this.setState({ newICOCreationPrice: event.target.value });
  }

  handleChangeICOCreationPrice(event) {
    event.preventDefault();

    const creationPrice = this.state.newICOCreationPrice;
    if (creationPrice === this.state.icoCreationPrice) {
      this.setState({ newICOCreationPrice: 0 });
      return;
    }

    this.props.onChangeICOCrationPriceSubmit(creationPrice, this);
  }

  handleFeeValueChange(event) {
    const fee = event.target.value;
    if (fee < 0 || fee > 100) {
      return;
    }
    this.setState({ newWithdrawFeePercent: fee });
  }

  handleWithdrawFeeChange(event) {
    event.preventDefault();

    const fee = this.state.newWithdrawFeePercent;
    if (fee === this.state.withdrawFeePercent) {
      this.setState({ newWithdrawFeePercent: 0 });
      return;
    }

    this.props.onChangeWithdrawFeeSubmit(fee, this);
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

      this.props.onAddNewAuditorSubmit(address, this);
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
          <h3>You don't have ICOs to submit withdraws</h3>
        </div>
      );
    }

    return this.state.auditorICOs.map(ico => (
      <BremAuditFormContainer key={ico.address} value={ico} />
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
    if (date < new Date(Date.now())) {
      return alert("Error, date must be bigger than now");
    }
    if (date) this.setState({ icoClosingTime: date.getTime() / 1000 });
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
    if (cap === 0) {
      alert("Cap must me bigger than 0");
    }

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
      files,
      this
    );
  }

  DevBremList() {
    const amount = this.state.devICOs.length;
    if (amount === 0) {
      return (
        <div>
          <p>You didn't create ICOs</p>
        </div>
      );
    }
    return this.state.devICOs.map(ico => (
      <BremDevFormContainer key={ico.address} value={ico} />
    ));
  }

  render() {
    const SuperuserForm = (
      <div>
        <h3>Superuser</h3>

        <h4>BREM</h4>
        {this.state &&
          this.state.bremBalance && (
            <div>
              <p>BREM balance: {this.state.bremBalance} Eth</p>
              {this.state &&
                this.state.bremBalance > 0 && (
                  <form
                    className="pure-form pure-form-ctacked"
                    onSubmit={this.handleWithdraw.bind(this)}
                  >
                    <fieldset>
                      <legend>Withdraw fees</legend>
                      <input
                        id="withdrawValue"
                        type="number"
                        step="0.000000000000000001"
                        value={this.state.withdrawValue}
                        onChange={this.handleWithdrawValueChange.bind(this)}
                        placeholder="Amount in ETH"
                      />
                      <span> Eth </span>
                      <p>
                        <button
                          type="submit"
                          className="pure-button pure-button-primary"
                        >
                          Withdraw fees
                        </button>
                      </p>
                    </fieldset>
                  </form>
                )}
            </div>
          )}

        {this.state &&
          this.state.icoCreationPrice !== undefined && (
            <div>
              <p>BREM ICO creation price: {this.state.icoCreationPrice} BRM</p>
              {this.state &&
                this.state.icoCreationPrice !== undefined &&
                this.state.newICOCreationPrice !== undefined && (
                  <form
                    className="pure-form pure-form-ctacked"
                    onSubmit={this.handleChangeICOCreationPrice.bind(this)}
                  >
                    <fieldset>
                      <legend>Change ICO creation price</legend>
                      <input
                        id="newCreatonPrice"
                        type="number"
                        step="0.000000000000000001"
                        value={this.state.newICOCreationPrice}
                        onChange={this.handleCreationPriceChange.bind(this)}
                        placeholder="Amount in BRM"
                      />
                      <span> BRM </span>
                      <p>
                        <button
                          type="submit"
                          className="pure-button pure-button-primary"
                        >
                          Change
                        </button>
                      </p>
                    </fieldset>
                  </form>
                )}
            </div>
          )}

        {this.state &&
          this.state.withdrawFeePercent !== undefined && (
            <div>
              <p>BREM ICO withdraw fee: {this.state.withdrawFeePercent} %</p>
              {this.state &&
                this.state.withdrawFeePercent !== undefined &&
                this.state.newWithdrawFeePercent !== undefined && (
                  <form
                    className="pure-form pure-form-ctacked"
                    onSubmit={this.handleWithdrawFeeChange.bind(this)}
                  >
                    <fieldset>
                      <legend>Change ICO creation price</legend>
                      <input
                        id="newFeePercent"
                        type="number"
                        step="1"
                        min="0"
                        max="100"
                        value={this.state.newWithdrawFeePercent}
                        onChange={this.handleFeeValueChange.bind(this)}
                      />
                      <span> % </span>
                      <p>
                        <button
                          type="submit"
                          className="pure-button pure-button-primary"
                        >
                          Change
                        </button>
                      </p>
                    </fieldset>
                  </form>
                )}
            </div>
          )}

        <h4>BRM Token</h4>
        {this.state &&
          this.state.mintAddress !== undefined &&
          this.state.mintAmount !== undefined && (
            <form
              className="pure-form pure-form-ctacked"
              onSubmit={this.handleMint.bind(this)}
            >
              <fieldset>
                <legend>Mint BRM tokens</legend>
                <input
                  id="mintAddress"
                  type="text"
                  value={this.state.mintAddress}
                  onChange={this.handleMintAddressChange.bind(this)}
                  placeholder="Address"
                />
                <input
                  id="mintValue"
                  type="number"
                  value={this.state.mintAmount}
                  onChange={this.handleMintAmountChange.bind(this)}
                  placeholder="Amount"
                />

                <button
                  type="submit"
                  className="pure-button pure-button-primary"
                >
                  Mint
                </button>
                <span className="pure-form-message">
                  To simplify the MVP, only manual BRM mint is available
                </span>
              </fieldset>
            </form>
          )}

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

        <h4>ICOs' withdraws requests</h4>
        {this.state && this.state.auditorICOs && this.AuditorICOList()}
      </div>
    );

    const DeveloperForm = (
      <div>
        <h3>Developer</h3>

        {this.state &&
          this.state.icoName !== undefined &&
          this.state.icoSymbol !== undefined &&
          this.state.icoRate !== undefined &&
          this.state.icoCap !== undefined &&
          this.state.icoClosingTime !== undefined &&
          this.state.icoDescription !== undefined &&
          this.state.price !== undefined && (
            <div>
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
                      value={this.state.icoName}
                      onChange={this.handleICONameChange.bind(this)}
                      placeholder="BREM ICO name"
                    />
                  </p>
                  <p>
                    <input
                      id="symbol"
                      type="text"
                      value={this.state.icoSymbol}
                      onChange={this.handleICOTokenSymbolChange.bind(this)}
                      placeholder="BREM ICO Token Symbol"
                    />
                  </p>
                  <p>
                    <input
                      id="rate"
                      type="number"
                      value={this.state.icoRate}
                      onChange={this.handleICORateChange.bind(this)}
                      placeholder="BREM ICO rate"
                    />
                  </p>
                  <p>
                    <input
                      id="cap"
                      type="number"
                      step="0.000000000000000001"
                      value={this.state.icoCap}
                      onChange={this.handleICOCapChange.bind(this)}
                      placeholder="BREM ICO cap"
                    />
                    <span> Eth </span>
                  </p>
                  <p>
                    <input
                      id="closingTime"
                      type="datetime-local"
                      value={this.state.closingTime}
                      onChange={this.handleICOClosingTimeChange.bind(this)}
                      placeholder="BREM ICO Closing Time"
                    />
                  </p>
                  <p>
                    <textarea
                      id="description"
                      rows="5"
                      cols="70"
                      value={this.state.icoDescription}
                      onChange={this.handleICODescriptionChange.bind(this)}
                      placeholder="BREM ICO Description"
                      wrap="soft"
                    />
                  </p>
                  <p>
                    <input
                      type="file"
                      onChange={this.handleICOFilesChange.bind(this)}
                      ref={node => this._addDirectory(node)}
                    />
                  </p>
                  <button
                    type="submit"
                    className="pure-button pure-button-primary"
                  >
                    Create new BREM ICO
                  </button>
                  <span className="pure-form-message">
                    <p>BREM ICO Creation price: {this.state.price} BRM</p>
                    <p>BREM Fee: {this.state.withdrawFeePercent} %</p>
                  </span>
                </fieldset>
              </form>
            </div>
          )}

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
