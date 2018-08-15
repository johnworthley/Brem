import React, { Component } from 'react'
import { Link } from 'react-router-dom'
import uuid from 'uuid'
import BlogPost from './BlogPost'

import Loader from '../Components/Loader'

import store from '../lib/store'

class BlogPosts extends Component {
  constructor() {
    super()
    this.state = {
      posts: [],
      loading: true
    }
  }
  componentWillMount() {
    fetch('https://guitclub.ru/api/blogposts')
    .then(data => data.json())
    .then(json => this.setState({
      posts: json,
      loading: false
    }))
  }
  deletePost = id => {
    fetch('https://guitclub.ru/api/blogposts/delete', {
      method: 'POST',
      body: JSON.stringify({id}),
      credentials: 'same-origin'
    })
    .then(data => data.json())
    .then(json => this.setState({
      posts: json,
      loading: false
    }))
  }
  render() {
    const { userstatus } = this.props
    const { posts = [], loading } = this.state
    const output = posts.map(item => (
      <BlogPost
        userstatus={userstatus}
         key={item.id}
         deletePost={this.deletePost}
          {...item}
         />
    ))
    return (
      <div className="blog_posts">
        {

          posts.length ? (
            output
          ) : (
            loading ? (
              <div style={{margin: '40px auto', width: 60}}>
                <Loader />
              </div>
            ) : (
              <div style={{textAlign: 'center', opacity: 0.7, marginTop: 40}}>Извините, здесь пока что нет новостей.</div>
            )

          )
        }
      </div>
    )
  }
}

export default BlogPosts
