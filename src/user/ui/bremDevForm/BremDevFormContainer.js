import { connect } from "react-redux";
import BremDevForm from "./BremDevForm";

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  };
};

const mapDispatchToProps = dispatch => {
  return {};
};

const BremDevFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremDevForm);

export default BremDevFormContainer;
