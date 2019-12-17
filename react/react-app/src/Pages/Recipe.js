import React from 'react';

function Recipe(props) {

  const id = props.match.params.id;

  return (
    <div>
      <h3>Show recipe test!</h3>
      <h3>Recipe ID: {id}</h3>
    </div>
  );
}

export default Recipe;