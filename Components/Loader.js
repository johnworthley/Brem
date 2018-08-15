import React, { Component } from 'react'

const styleLoaderWrapper = {
  width: 60,
  height: 60,
  borderRadius: '50%'
}

const styleLoader = {
  width: 60,
  height: 60,
  background: 'linear-gradient(35deg, rgb(115, 27, 65), rgb(26, 52, 95))',
  borderRadius: '50%',
  position: 'relative',
}

const styleLoaderInner = {
  width: 50,
  height: 50,
  background: '#fff',
  borderRadius: '50%',
  position: 'absolute',
  left: '50%',
  top: '50%',
  transform: 'translate(-50%, -50%)',
  boxSizing: 'border-box',
}

const styleLoaderBg = {
    animation: 'loader_bg 4s ease-in-out infinite',
    position: 'absolute',
    width: '145%',
    height: '145%',
    left: '50%',
    top: '50%',
    background: 'radial-gradient(ellipse at top, transparent 0%,transparent 40%,#fff 40%,#fff 100%)',
    borderRadius: '50%',
    zIndex: 1,
}

const styleLoaderOuter = {
  animation: 'guitarclub_spin 4s linear infinite',
  width: 60,
  height: 60,
  position: 'relative',
  overflow: 'hidden',
  borderRadius: '50%',
}



export default props => {
  return(
    <div style={styleLoaderOuter}>
      <div style={styleLoaderBg} />
      <div style={styleLoaderWrapper}>
        <div style={styleLoader}>
          <div style={styleLoaderInner}></div>
        </div>
      </div>
    </div>
  )
}
