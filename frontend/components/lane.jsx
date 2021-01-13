import React, { useState, useEffect, useRef } from 'react';
import Card from './card';
import './lane.module.scss';
import { Form, Button } from 'react-bootstrap';

const Lane = (props) => {

  const [showInput, setShowInput] = useState(false);
  const taskInput = useRef(null);

  useEffect(() => {
    showInput && taskInput.current.focus();
  }, [showInput]);

  const handleClickAddCard = (e) => {
    e.target.blur();
    setShowInput(true);
  };

  const handleClickCancel = () => {
    setShowInput(false);
  };

  const handleClickAdd = () => {
    setShowInput(false);
  };

  return (
    <div className="lane col-md-4">
      <div className="laneHeader">
        <h5>Tasks</h5>
      </div>
      {props.taskList.map((task) => {
        return <Card key={task.id} title={task.title} />;
      })}
      <div className="laneFooter">
        {showInput &&
        <Form className="taskInputForm">
          <Form.Control type="text" ref={taskInput} />
            <Button className="mr-2" onClick={handleClickAdd}>Add</Button>
            <Button variant="" onClick={handleClickCancel}>Cancel</Button>
        </Form>
        }

        {!showInput &&
        <button type="button" className="btn btnAddCard" onClick={handleClickAddCard}>+ Add card</button>}
      </div>
    </div>
  );
};

export default Lane;