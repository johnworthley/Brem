import { connect } from "react-redux";
import BremICOForm from "./BremICOForm";

const mapStateToProps = (state, ownProps) => {
  return {};
};

const mapDispatchToProps = dispatch => {
  return {};
};

const BremPublicationFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremICOForm);

export default BremPublicationFormContainer;
