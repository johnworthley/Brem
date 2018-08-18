import React, { Component } from 'react'
import { css, StyleSheet } from 'aphrodite/no-important'


const FormRow = ({ setRef, label, required, placeholder}) => (
  <div>
    <label htmlFor={label}>{ label } {required ? <span>*</span> : ''}</label>
    <input ref={setRef} type="text" placeholder={placeholder} name={label}/>
  </div>
)

class Create extends Component {
  state = {

  }
  setRef = name => elem => this[name] = elem

  handleProjectCreate = e => {
    e.preventDefault()
    const { name, desc, time, street, token, cap, ratio } = this
    const data = {
      name: name.value,
      desc: desc.value,
      time: time.value,
      street: street.value,
      token: token.value,
      cap: cap.value,
      ratio: ratio.value
    }
    console.log(data)
  }

  render() {
    const { setRef } = this
    return (
      <div>
        <h3 className="page-heading">
          New Project
        </h3>
        <form onSubmit={this.handleProjectCreate}>
          <FormRow setRef={setRef('name')} label="Name" placeholder="Object name" required />
          <FormRow setRef={setRef('desc')} label="Desctiption" placeholder="Tell us more about the project" required />
          <FormRow setRef={setRef('time')} label="Closing time" placeholder="Closing time" required />
          <FormRow setRef={setRef('street')} label="Street" placeholder="Street" required />
          <FormRow setRef={setRef('token')} label="Token Name" placeholder="Enter token name(3 letters)" required />
          <FormRow setRef={setRef('cap')} label="Soft-Cap" placeholder="Soft cap for your project" required />
          <FormRow setRef={setRef('ratio')} label="Rate" placeholder="token to WEI ratio" required />
          <button>Submit</button>
        </form>
      </div>
    )
  }
}

export default Create
