import React, { useState, useEffect, useRef } from 'react';
import Card from './card';
import './lane.module.scss';
import { Form, Button } from 'react-bootstrap';
import axios from 'axios';

const Lane = (props) => {

  const [showInput, setShowInput] = useState(false);
  const [taskList, setTaskList] = useState(props.taskList || []);
  const taskInput = useRef(null);

  useEffect(() => {
    // focus input form when it is shown
    showInput && taskInput.current.focus();
  }, [showInput]);

  useEffect(() => {
    setTaskList(props.taskList);
  }, [props.taskList]);

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
      const res = await axios.post(uri, {
        title: taskInput.current.value,
        user_id: props.userId,
        lane_id: props.id,
      }, {
        headers: {
          Authorization: `Bearer ${token}`,
        }
      });

      taskList ? setTaskList([...taskList, res.data]) : setTaskList([res.data]);

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
      {taskList && taskList.map((task) => {
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