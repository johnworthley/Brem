import React, { Component } from "react";
import axios from "axios";
import BremItem from "../BremItem/BremItem";
import mHost from "../../config";

class MarketplaceForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      status: props.value
    };

    this.BremList.bind(this);
    this.init();
  }

  init() {
    switch (this.state.status) {
      case "opened":
        axios
          .get("http://" + mHost + "/ico/opened")
          .then(res => {
            this.setState({ icos: res.data });
          })
          .catch(err => console.error(err));
        break;
      case "success":
        axios
          .get("http://" + mHost + "/ico/success")
          .then(res => this.setState({ icos: res.data }));
        break;
      case "failed":
        axios
          .get("http://" + mHost + "/ico/failed")
          .then(res => this.setState({ icos: res.data }));
        break;
      case "withdrawn":
        axios
          .get("http://" + mHost + "/ico/withdrawn")
          .then(res => this.setState({ icos: res.data }));
        break;
      default:
        this.setState({ invalidStaus: true });
        break;
    }
  }

  BremList() {
    const amount = this.state.icos.length;
    if (amount === 0) {
      return (
        <div>
          <h3>Brem ICOs not created yet...</h3>
        </div>
      );
    }

    return this.state.icos.map(ico => (
      <BremItem key={ico.address} value={ico.address} />
    ));
  }

  render() {
    return <div>{this.state && this.state.icos && this.BremList()}</div>;
  }
}

export default MarketplaceForm;
