import { connect } from "react-redux";
import BremICOForm from "./BremICOForm";
import { buyTokens, refund } from "./BremICOFormActions";

const mapStateToProps = (state, ownProps) => {
  return {};
};

const mapDispatchToProps = dispatch => {
  return {
    onBuyTokenSubmit: (contractAddress, tokenAddress, etherAmount, form) => {
      event.preventDefault();

      dispatch(buyTokens(contractAddress, tokenAddress, etherAmount, form));
    },
    onRefundSubmit: (contractAddress, tokenAddress, form) => {
      event.preventDefault();

      dispatch(refund(contractAddress, tokenAddress, form));
    }
  };
};

const BremPublicationFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremICOForm);

export default BremPublicationFormContainer;
