import React, { Component } from "react";

class BremItem extends Component {
  constructor(props) {
    super(props);

    this.state = {
      index: props.value
    };
  }

  render() {
    return (
      <div>
        <span>This is BREM item №{this.state.index}</span>
      </div>
    );
  }
}

export default BremItem;
