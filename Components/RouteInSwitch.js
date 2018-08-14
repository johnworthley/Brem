import React, { Component } from 'react';
import { Route } from 'react-router-dom'

class RouteInSwitch extends Component {
  handleUpdateHeader() {

  }
  render() {
    const location = this.props.location;
    const path = this.props.path
    return (
      <Route location={location} key={location.key} exact path={path} component={this.props.component} props={{updateHeader: this.handleUpdateHeader.bind(this)}} />
    );
  }
}

export default RouteInSwitch;
