import { connect } from "react-redux";
import BremICOForm from "./BremICOForm";
import { buyTokens } from "./BremICOFormActions";

const mapStateToProps = (state, ownProps) => {
  return {};
};

const mapDispatchToProps = dispatch => {
  return {
    onBuyTokenSubmit: (contractAddress, tokenAddress, etherAmount, form) => {
      event.preventDefault();

      dispatch(buyTokens(contractAddress, tokenAddress, etherAmount, form));
    }
  };
};

const BremPublicationFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremICOForm);

export default BremPublicationFormContainer;
