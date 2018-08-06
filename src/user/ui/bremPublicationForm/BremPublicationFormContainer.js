import { connect } from "react-redux";
import BremPublicationForm from "./BremPublicationForm";
import { addNewAuditor, publishProject } from "./BremPublicationFormActions";

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  };
};

const mapDispatchToProps = dispatch => {
  return {
    onAddNewAuditorSubmit: (contractAddress, auditorAddress, form) => {
      event.preventDefault();

      dispatch(addNewAuditor(contractAddress, auditorAddress, form));
    },
    onPublishProjectSubmit: (contractAddress, form) => {
      event.preventDefault();

      dispatch(publishProject(contractAddress, form));
    }
  };
};

const BremPublicationFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremPublicationForm);

export default BremPublicationFormContainer;
