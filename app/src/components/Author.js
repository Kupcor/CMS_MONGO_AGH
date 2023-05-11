import React, { useState, useEffect } from "react";

function Author({authorId}) {
  const [author, setAuthor] = useState({});

  useEffect(() => {
    fetch(`http://localhost:8000/users/${authorId}`)
      .then((res) => res.json())
      .then((data) => setAuthor(data))
      .catch((error) => console.log(error));
  }, [authorId]);

  return (
    <div>
      <p>{author.username}</p>
    </div>
  );
}

export default Author;