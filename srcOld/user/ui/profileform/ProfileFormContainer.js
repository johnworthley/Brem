import { connect } from "react-redux";
import ProfileForm from "./ProfileForm";
import {
  mintBRMTokens,
  createNewBREMICO,
  addNewAuditor,
  withdrawFees,
  changeICOCreationPrice,
  changeWithdrawFee
} from "./ProfileFormActions";

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  };
};

const mapDispatchToProps = dispatch => {
  return {
    onMintFormSubmit: (to, amount, form) => {
      event.preventDefault();

      dispatch(mintBRMTokens(to, amount, form));
    },
    onCreateBREMICOFormSubmit: (
      name,
      symbol,
      rate,
      cap,
      closingTime,
      description,
      files,
      image,
      form
    ) => {
      event.preventDefault();

      dispatch(
        createNewBREMICO(
          name,
          symbol,
          rate,
          cap,
          closingTime,
          description,
          files,
          image,
          form
        )
      );
    },
    onAddNewAuditorSubmit: (address, form) => {
      event.preventDefault();

      dispatch(addNewAuditor(address, form));
    },
    onWithdrawFormSubmit: (value, form) => {
      event.preventDefault();

      dispatch(withdrawFees(value, form));
    },
    onChangeICOCrationPriceSubmit: (icoCreationPrice, form) => {
      event.preventDefault();

      dispatch(changeICOCreationPrice(icoCreationPrice, form));
    },
    onChangeWithdrawFeeSubmit: (value, form) => {
      event.preventDefault();

      dispatch(changeWithdrawFee(value, form));
    }
  };
};

const ProfileFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ProfileForm);

export default ProfileFormContainer;
