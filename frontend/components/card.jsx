import React from 'react';
import './card.module.scss';
import {Button} from 'react-bootstrap';
import axios from 'axios';

const Card = (props) => {

  const handleClickRemove = async (e) => {
    e.target.blur();
    e.stopPropagation();

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

  const handleClickCard = () => {
    // show task detail modal
    props.setShowTaskModal(true);
    // set this task to modal
    props.setTargetTask(props.task);
  };

  return (
    <div className="card" className="card" onClick={handleClickCard}>
      <div className="title d-inline-block">{props.title}</div>
      <Button variant="white" onClick={handleClickRemove}>x</Button>
    </div>
  );
};

export default Card;