import { connect } from "react-redux";
import BremICOForm from "./BremICOForm";
import { buyTokens } from "./BremICOFormActions";

const mapStateToProps = (state, ownProps) => {
  return {};
};

const mapDispatchToProps = dispatch => {
  return {
    onBuyTokenSubmit: (contractAddress, weiAmount) => {
      event.preventDefault();

      dispatch(buyTokens(contractAddress, weiAmount));
    }
  };
};

const BremPublicationFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremICOForm);

export default BremPublicationFormContainer;
