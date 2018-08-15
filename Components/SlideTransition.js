import React, { Component } from 'react'
import { Motion, TransitionMotion, spring } from 'react-motion'

const SpringModel = { stiffness: 250, damping: 25, precision: 0.05 }

class SlideTransition extends Component {
  willLeave() {
    // triggered when c's gone. Keeping c until its width/height reach 0.
    return {opacity: spring(0.5)}
  }
  render() {
    const { urlkey: key, items } = this.props
    console.log(items)
    return (
      <TransitionMotion
            willLeave={this.willLeave}
            styles={items.map(item => ({
              key: item.id,
              data: {
                comp: item.comp
              },
              style: {opacity: 1}
            }))
        }>
            {interpolatedStyles => {
              if(!interpolatedStyles || !interpolatedStyles.data) return <div></div>
              console.log(interpolatedStyles)

              return <div>
                {

                  interpolatedStyles.map(config => {
                  const Form = interpolatedStyles.data.comp
                  console.log(Form)
                  return <div key={config.key}>
                    {<interpolatedStyles.data.comp />}
                  </div>
                })}
              </div>
            }
              // first render: a, b, c. Second: still a, b, c! Only last one's a, b.

            }
          </TransitionMotion>
    )
  }
}

export default SlideTransition
