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
    onAddNewAuditorSubmit: (contractAddress, auditorAddress) => {
      event.preventDefault();

      dispatch(addNewAuditor(contractAddress, auditorAddress));
    },
    onPublishProjectSubmit: contractAddress => {
      event.preventDefault();

      dispatch(publishProject(contractAddress));
    }
  };
};

const BremPublicationFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(BremPublicationForm);

export default BremPublicationFormContainer;
