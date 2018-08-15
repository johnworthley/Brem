import { connect } from "react-redux";
import BremAuditForm from "./BremAuditForm";
import { confirmWithdraw } from "./BremAuditFormActions";

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  };
};

const mapDispatchToProps = dispatch => {
  return {
    onConfirmSubmit: (contractAddress, form) => {
      event.preventDefault();

      dispatch(confirmWithdraw(contractAddress, form));
    }
  };
};

const BremAuditFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremAuditForm);

export default BremAuditFormContainer;
