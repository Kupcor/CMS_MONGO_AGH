import React, { useState, useEffect } from "react";

function Category({categoryId}) {
  const [category, setCategory] = useState({});

  useEffect(() => {
    fetch(`http://localhost:8000/categories/${categoryId}`)
      .then((res) => res.json())
      .then((data) => setCategory(data))
      .catch((error) => console.log(error));
  }, [categoryId]);

  return (
    <div>
      {category.categoryName}
    </div>
  );
}

export default Category;