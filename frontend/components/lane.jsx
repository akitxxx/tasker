import React, { useState, useEffect, useRef } from 'react';
import Card from './card';
import './lane.module.scss';
import { Form, Button } from 'react-bootstrap';
import axios from 'axios';

const Lane = (props) => {

  const [showInputTask, setShowInputTask] = useState(false);
  const [showInputLaneName, setShowInputLaneName] = useState(false);
  const [taskList, setTaskList] = useState(props.taskList || []);
  const taskInput = useRef(null);
  const laneNameInput = useRef(null);

  useEffect(() => {
    // focus new task input form when it is shown
    showInputTask && taskInput.current.focus();
  }, [showInputTask]);

  useEffect(() => {
    // focus lane name input form when it is shown
    showInputLaneName && laneNameInput.current.focus();
  }, [showInputLaneName]);

  useEffect(() => {
    setTaskList(props.taskList);
  }, [props.taskList]);

  const handleClickAddCard = (e) => {
    e.target.blur();
    setShowInputTask(true);
  };

  const handleClickCancel = () => {
    setShowInputTask(false);
  };

  const handleKeyDownTaskInput = (e) => {
    if(e.keyCode === 13) {
      handleClickAdd();
    }
  };

  /**
   * Create task
   */
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

      setShowInputTask(false);
    } catch(e) {
      alert(e);
      return;
    }
  };

  /**
   * Delete lane
   */
  const handleClickRemove = async () => {
    const uri = `/api/delete-lane/${props.id}`;
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

  const handleClickLaneName = () => {
    setShowInputLaneName(true);
  };

  const handleKeyDownLaneNameInput = async (e) => {
    if(e.keyCode === 13) {
      updateLane(e);
    }
  };

  const handleBlurLaneNameInput = () => {
    setShowInputLaneName(false);
  };

  const updateLane = async (e) => {
    const uri = "/api/update-lane";
    const token = localStorage.getItem('tasker_token');

    try {
      const res = await axios.put(uri, {
        id: props.id,
        name: e.target.value,
      }, {
        headers: {
          Authorization: `Bearer ${token}`,
        }
      });

      res.data && await props.fetchTaskList();

      setShowInputLaneName(false);
    } catch(e) {
      alert(e);
      return;
    }
  };

  return (
    <div className="lane col-md-4">
      <div className="laneHeader">
        {showInputLaneName &&
          <Form.Control type="text" className="laneNameInputForm" defaultValue={props.name} ref={laneNameInput} onKeyDown={handleKeyDownLaneNameInput} onBlur={handleBlurLaneNameInput} />
        }
        {!showInputLaneName &&
          <Form.Control className="laneName" plaintext readOnly defaultValue={props.name} onClick={handleClickLaneName} />
        }
        <Button variant="white" size="sm" className="btnRemove float-right" onClick={handleClickRemove}>x</Button>
      </div>
      {taskList && taskList.map((task) => {
         return <Card key={task.id} id={task.id} title={task.title} fetchTaskList={props.fetchTaskList}/>;
      })}
      <div className="laneFooter">
        {showInputTask &&
        <Form className="taskInputForm" onSubmit={(e) => {e.preventDefault();}}>
          <Form.Control type="text" ref={taskInput} onKeyDown={handleKeyDownTaskInput} />
          <Button className="mr-2" onClick={handleClickAdd}>Add</Button>
          <Button variant="default" className="btnCancel" onClick={handleClickCancel}>Cancel</Button>
        </Form>
        }

        {!showInputTask &&
        <Button variant="default" className="btnAddCard" onClick={handleClickAddCard}>+ Add task</Button>}
      </div>
    </div>
  );
};

export default Lane;