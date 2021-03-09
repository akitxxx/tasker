import React from 'react';
import './card.module.scss';
import {Button} from 'react-bootstrap';
import axios from 'axios';
import ClearIcon from '@material-ui/icons/Clear';
import { Draggable } from 'react-beautiful-dnd';

const Card = (props) => {

  const handleClickRemove = async (e) => {
    e.target.blur();
    e.stopPropagation();

    if(!confirm('Are you sure you want to delete this task?')) {
      return;
    }

    const uri = `/api/delete-task/${props.task.id}`;
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
    <Draggable draggableId={props.task.id.toString()} index={props.task.index_num}>
      {(provided) => (
        <div
          ref={provided.innerRef}
          {...provided.draggableProps}
          {...provided.dragHandleProps}
          className="card" className="card" onClick={handleClickCard}
        >
          <div className="title d-inline-block">{props.task.title}</div>
          <Button variant="white" onClick={handleClickRemove}><ClearIcon style={{ fontSize: 15 }} /></Button>
        </div>
      )}
    </Draggable>
  );
};

export default Card;