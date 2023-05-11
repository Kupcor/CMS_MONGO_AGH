import React, { useState, useEffect } from "react";
import Post from "./Post";

function PostList() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8000/posts")
      .then((res) => res.json())
      .then((data) => setPosts(data))
      .catch((error) => console.log(error));
  }, []);

  return (
    <div>
      <h1>Post List</h1>
      <ul>
        {posts.map((post) => (
          <Post postId={post._id}/>
        ))}
      </ul>
    </div>
  );
}

export default PostList;