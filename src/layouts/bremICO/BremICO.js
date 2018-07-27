import React, { Component } from "react";
import BremICOFormContainer from "../../ui/BremICOForm/BremICOFormContainer";

class BremICO extends Component {
  render() {
    return (
      <main className="container">
        <div className="pure-g">
          <div className="pure-u-1-1">
            <h1>BREM ICO</h1>
            <BremICOFormContainer value={this.props.params.address} />
          </div>
        </div>
      </main>
    );
  }
}

export default BremICO;
