import { connect } from "react-redux";
import BremDevForm from "./BremDevForm";
import { makeWithrawRequest } from "./BremDevFormActions";

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  };
};

const mapDispatchToProps = dispatch => {
  return {
    onSubmitWithdrawRequest: (contractAddress, weiValue, form) => {
      event.preventDefault();

      dispatch(makeWithrawRequest(contractAddress, weiValue, form));
    }
  };
};

const BremDevFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremDevForm);

export default BremDevFormContainer;
