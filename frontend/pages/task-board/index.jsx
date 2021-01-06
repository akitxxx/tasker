import React, { useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import Card from '../../components/card';
import axios from 'axios';
import { useEffect } from 'react';

const TaskBoard = () => {

  const [taskList, setTaskList] = useState([]);

  useEffect(() => {
    // componentDidMount
    (async () => {
      const uri = '/api/task';
      // get token from local storage
      const token = localStorage.getItem('tasker_token');
      // get task list from server
      const res = await axios.get(uri, {
        headers: {
          Authorization: `Bearer ${token}`,
        }
      });
      setTaskList(res.data)
    }) ();
  });

  return (
    <Container className='task-board'>
      <Row>
        <Col className='text-center'>
          <h2>Task Board</h2>
        </Col>
      </Row>
      <Row>
        <Col>
          {taskList.map((task) => {
            return <Card key={task.id} title={task.title} />;
          })}
        </Col>
      </Row>
    </Container>
  );
};

export default TaskBoard;