import { connect } from "react-redux";
import BremAuditForm from "./BremAuditForm";

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  };
};

const mapDispatchToProps = dispatch => {
  return {};
};

const BremAuditFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremAuditForm);

export default BremAuditFormContainer;
