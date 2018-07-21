import { connect } from 'react-redux'
import ProfileForm from './ProfileForm'
import { mintBRMTokens, createNewBREMICO } from './ProfileFormActions'

const mapStateToProps = (state, ownProps) => {
  return {
    name: state.user.data.name
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    onMintFormSubmit: (to, amount) => {
      event.preventDefault()

      dispatch(mintBRMTokens(to, amount))
    },
    onCreateBREMICOFormSubmit: (name, symbol, rate, cap, closingTime, description) => {
        event.preventDefault()

        dispatch(createNewBREMICO(name, symbol, rate, cap, closingTime, description))
    }
  }
}

const ProfileFormContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ProfileForm)

export default ProfileFormContainer
