import React, { useState, useEffect } from "react";
import Author from "./Author";

function Content({contentId }) {
  const [content, setContent] = useState({});

  useEffect(() => {
    fetch(`http://localhost:8000/contents/id/${contentId}`)
      .then((res) => res.json())
      .then((data) => setContent(data))
      .catch((error) => console.log(error));
  }, [contentId]);

  return (
    <div>
      <h3>{content.title}</h3>
      <p>{content.content}</p>
      <p><Author authorId={content.authorId}/></p>
    </div>
  );
}

export default Content;