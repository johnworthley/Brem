import { connect } from "react-redux";
import BremICOForm from "./BremICOForm";
import {
  buyTokens,
  refund,
  addNewAuditor,
  publishProject,
  makeWithrawRequest,
  confirmWithdraw
} from "./BremICOFormActions";

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
    },
    onAddNewAuditorSubmit: (contractAddress, auditorAddress, form) => {
      event.preventDefault();

      dispatch(addNewAuditor(contractAddress, auditorAddress, form));
    },
    onPublishProjectSubmit: (contractAddress, form) => {
      event.preventDefault();

      dispatch(publishProject(contractAddress, form));
    },
    onSubmitWithdrawRequest: (contractAddress, value, form) => {
      event.preventDefault();

      dispatch(makeWithrawRequest(contractAddress, value, form));
    },
    onConfirmSubmit: (contractAddress, form) => {
      event.preventDefault();

      dispatch(confirmWithdraw(contractAddress, form));
    }
  };
};

const BremPublicationFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremICOForm);

export default BremPublicationFormContainer;
