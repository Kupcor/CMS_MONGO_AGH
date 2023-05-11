import React, { useState, useEffect } from "react";
import "./components.css";

function CommentsList({postId}) {
  const [comments, setComments] = useState([]);
  
  useEffect(() => {
    fetch(`http://localhost:8000/comments/post/${postId}`)
      .then((res) => res.json())
      .then((data) => setComments(data))
      .catch((err) => console.log(err));
  }, [postId]);
   
  if (!comments) {
    return null
  }

  return ( 
   <div className="comment-list">
      <h3>Komentarze</h3>
      {comments.map((comment) => (
        <div key={comment._id} className="comment">
          <div>{comment.content}</div>
        </div>
      ))}
    </div>
  );
}

export default CommentsList;