import React, { Component } from 'react';
import { Link } from 'react-router-dom'
import uuid from 'uuid'
import love from '../img/icons/love.svg'


const cardStyles = {
  boxSizing: 'border-box',
  overflow: 'hidden',
  willChange: 'transform, opacity, filter',
  lineHeight: '1.3em',
}


const decorStyle = {
  ...abs,
  top: 0,
  right: 0,
  bottom: 0,
  left: 0,
  padding: '20px 20px 20px 260px',
  overflow: 'hidden',
  ...userNone
}

const abs = {
  position: 'absolute'
}

const rel = {
  position: 'relative'
}

const userNone = {
  userSelect: 'none',
  pointerEvents: 'none'
}

const fill = {
  ...abs,
  left: 0,
  right: 0,
  top: 0,
  bottom: 0
}

const decorHeaderStyle = {
  ...abs,
  transform: 'skewX(25deg) translate(-84px, 20px)',
  height: 195,
  width: '100%'
}

const decorThingStyle = {
  //backgroundImage: 'linear-gradient(to right, transparent, #001d4d)',
  background: '#001d4d',
  ...abs,
  top: 0,
  height: '100%',
  left: 80,
  width: 'calc(100% - 280px)',
  paddingRight: 22.75,
  paddingBottom: 20,
  transform: 'skewX(-25deg) translateX(254.755px)',
  boxSizing: 'border-box',
  color: '#FFF',
  opacity: .9,
  zIndex: 1,
  ...userNone
}

const decorBgStyle = {
  filter: 'blur(3px)',
  ...fill,
  left: '10%',
  backgroundImage: 'linear-gradient(to right, transparent, #660029)',
}

const pStyle = {
  lineHeight: '1.5em',
  padding: '6px 5px 6px 0',
  userSelect: 'none'
}

const hStyle = {
  lineHeight: '1.5em',
  padding: '6px 5px 6px 0',
  userSelect: 'none'
}

const readStyle = {
  ...abs,
  right: 0,
  bottom: 0
}

const decorBlackStripStyle = {
  ...abs,
  width: 60,
  height: 160,
  left: 197,
  top: 0,
  //background: 'black',
  //backgroundImage: 'linear-gradient(30deg, #282828, #121212)',
  transform: 'skewX(-25deg)',
  zIndex: -1,
  ...userNone
}

const styleDelete = {
  ...abs,
  right: 15,
  top: 15,
  width: 15,
  height: 15,
  background: 'red',
  zIndex: 10,
  cursor: 'pointer'
}

class BlogPost extends Component {
  constructor() {
    super()
  }
  render() {
    const { id, userstatus, thumbnail, title: heading, description, tag } = this.props
    const style = {
      backgroundImage: `url(${thumbnail})`
    }
    return (
      <section style={cardStyles} className="blog_post">
        {
          userstatus > 10 ? <div onClick={this.props.deletePost.bind(this, id)} style={styleDelete} /> : ''
        }
        <div className="blog_post_img_holder">
          <aside className="blog_post_img" style={style}></aside>
        </div>
        <div className="blog_post_inner">
          <header className="blog_post_header">
            <Link to={`/blog/posts/${id}`}>
              <h3 style={hStyle} className="heading blog_post_heading">
                {heading}
              </h3>
            </Link>
            <br/>
            <p style={pStyle} className="paragraph blog_post_paragraph blog_post_paragraph-main">
              {description}
            </p>
          </header>
          <div className="blog_post_type">{tag}</div>
        </div>
        <div style={decorStyle}>
          <div style={decorThingStyle}>
            <div style={decorBgStyle}></div>
            <header style={decorHeaderStyle} className="blog_post_header">
              <Link to={`/blog/posts/${id}`}>
                <h3 style={{color: '#006eb3', ...hStyle}} className="heading blog_post_heading blog_post_heading-dark">
                  {heading}
                </h3>
              </Link>
              <br/>
              <p style={pStyle} className="paragraph blog_post_paragraph_dark blog_post_paragraph">
                {description}
              </p>
              <Link to={`/blog/posts/${id}`}>
                <div style={readStyle} className="blog_post_button">
                  <span>
                    читать
                  </span>
                </div>
              </Link>
            </header>
          </div>
        </div>
        <div style={decorBlackStripStyle} />
      </section>
    )
  }
}

export default BlogPost
