import React from 'react';
import './card.module.scss';
import {Button} from 'react-bootstrap';
import axios from 'axios';

const Card = (props) => {

  const handleClickRemove = async (e) => {
    e.target.blur();

    const uri = `/api/delete-task/${props.id}`;
    const token = localStorage.getItem('tasker_token');

    try {
      await axios.delete(uri, {
        headers: {
          Authorization: `Bearer ${token}`,
        }
      });
    } catch(e) {
      alert(e);
      return;
    }

    props.fetchTaskList();
  };

  return (
    <div className="card" className="card">
      <div className="title d-inline-block">{props.title}</div>
      <Button variant="white" onClick={handleClickRemove}>x</Button>
    </div>
  );
};

export default Card;