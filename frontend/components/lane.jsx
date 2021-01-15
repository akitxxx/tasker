import React, { useState, useEffect, useRef } from 'react';
import Card from './card';
import './lane.module.scss';
import { Form, Button } from 'react-bootstrap';
import axios from 'axios';

const Lane = (props) => {

  const [showInput, setShowInput] = useState(false);
  const taskInput = useRef(null);

  useEffect(() => {
    // focus input form when it is shown
    showInput && taskInput.current.focus();
  }, [showInput]);

  const handleClickAddCard = (e) => {
    e.target.blur();
    setShowInput(true);
  };

  const handleClickCancel = () => {
    setShowInput(false);
  };

  const handleKeyDown = (e) => {
    if(e.keyCode === 13) {
      handleClickAdd();
    }
  };

  const handleClickAdd = async () => {
    const uri = '/api/create-task';
    const token = localStorage.getItem('tasker_token');

    try {
      const res = await axios.post(uri,{
        title: taskInput.current.value,
      }, {
        headers: {
          Authorization: `Bearer ${token}`,
        }
      });

      props.taskList && props.setTaskList([...props.taskList, res.data]);

      setShowInput(false);
    } catch(e) {
      alert(e);
      return;
    }
  };

  return (
    <div className="lane col-md-4">
      <div className="laneHeader">
        <h5>Tasks</h5>
      </div>
      {props.taskList && props.taskList.map((task) => {
         return <Card key={task.id} id={task.id} title={task.title} fetchTaskList={props.fetchTaskList}/>;
      })}
      <div className="laneFooter">
        {showInput &&
        <Form className="taskInputForm" onSubmit={(e) => {e.preventDefault();}}>
          <Form.Control type="text" ref={taskInput} onKeyDown={handleKeyDown} />
            <Button className="mr-2" onClick={handleClickAdd}>Add</Button>
            <Button variant="default" className="btnCancel" onClick={handleClickCancel}>Cancel</Button>
        </Form>
        }

        {!showInput &&
        <Button variant="default" className="btnAddCard" onClick={handleClickAddCard}>+ Add task</Button>}
      </div>
    </div>
  );
};

export default Lane;