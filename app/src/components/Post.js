import React, { useState, useEffect } from "react";
import Content from "./Content";
import Category from "./Category";
import CommentsList from "./CommentsList"
import Author from "./Author";

import './components.css'

function Post({postId}) {
  const [post, setPost] = useState({});

  useEffect(() => {
    fetch(`http://localhost:8000/posts/${postId}`)
      .then((res) => res.json())
      .then((data) => setPost(data))
      .catch((error) => console.log(error));
  }, [postId]);

  return (
    <div key={post._id?.$oid}>
      <h1>{post.title}</h1>
      <div className="post-date">Opublikowany dn. <strong>{new Date(post.createdAt).toLocaleDateString("pl-Pl")}</strong></div>
      <Author authorId={post.authorId} />
      <p>{post.summary}</p>
      {post.contents?.map((content) => (<Content contentId={content}/>))}

      <div className="categories">Zaszufladkowano do kategorii: 
        <strong>
          {post.categories.map((category) => (
            <Category categoryId={category}/>
          ))}
        </strong>
      </div>

      <CommentsList postId={post._id}/>
      <hr/>          
  </div>
  );
}
      /*

      <div className="categories">Zaszufladkowano do kategorii: 
        <strong>
          {post.categories.map((category) => (
            <Category categoryId={category}/>
          ))}
        </strong>
      </div>

      */

export default Post;