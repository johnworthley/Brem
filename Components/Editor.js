import React, { Component } from 'react'

const editorStyle = {
  marginTop: 15,
  border: 'solid 1px rgba(0,0,0,.2)',
  minHeight: 200,
  padding: 10,
  overflow: 'auto',
  background: '#fff'
}




class Editor extends Component {
  constructor() {
    super()
    this.state = {
      editorState: {
        html: ''
      }
    }
  }
  componentDidUpdate() {
    this.refs.editor.focus()
  }
  addSquare() {
    const { editorState } = this.state
    editorState.html += ' <div style="background: red; width: 20px; height: 20px"></div>'
    this.setState({ editorState })
  }

  handleStateChange(e) {
    const elem = e.target
    const { editorState } = this.state
    editorState.html = elem.innerHTML
    this.setState({ editorState })
    this.refs.editor.innerHTML = elem.innerHTML

  }

  handleElemFocus(e) {
    const elem = e.target
    elem.focus()
  }

  render() {
    const { editorState } = this.state
    const { html } = editorState
    const markupObject = {__html: html}
    return (
      <div>
        <div ref="editor" onInput={this.handleStateChange.bind(this)} contentEditable style={editorStyle}>
        </div>
        <button onClick={this.addSquare.bind(this)}>Add</button>
      </div>
    )
  }
}

export default Editor
