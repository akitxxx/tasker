import React, { useState } from 'react';
import './card.module.scss';

const Card = (props) => {
  return (
    <div className="card" className="card">
      <div className="title">{props.title}</div>
    </div>
  );
};

export default Card;