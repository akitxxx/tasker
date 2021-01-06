import React, { useState } from 'react';

const Card = (props) => {
  return (
    <div className="cord">
      <div className="title">{props.title}</div>
    </div>
  );
};

export default Card;