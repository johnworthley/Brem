import { connect } from "react-redux";
import BremAuditForm from "./BremAuditForm";
import { approveWithdraw } from "./BremAuditFormActions";

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  };
};

const mapDispatchToProps = dispatch => {
  return {
    onApproveSubmit: contractAddress => {
      event.preventDefault();

      dispatch(approveWithdraw(contractAddress));
    }
  };
};

const BremAuditFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremAuditForm);

export default BremAuditFormContainer;
