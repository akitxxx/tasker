import React, { useState } from 'react';
import Router from 'next/router';
import { Container, Row, Col } from 'react-bootstrap';
import Lane from '../../components/lane';
import axios from 'axios';
import { useEffect } from 'react';
import '../../styles/task-board.scss';

const TaskBoard = () => {

  const [taskList, setTaskList] = useState([]);

  useEffect(() => {
    // componentDidMount
    fetchTaskList();
  },[]);

  const fetchTaskList = async () => {
      const uri = '/api/task';
      // get token from local storage
      const token = localStorage.getItem('tasker_token');

      try {
        // get task list from server
        const res = await axios.get(uri, {
          headers: {
            Authorization: `Bearer ${token}`,
          }
        });
        setTaskList(res.data);
      } catch(e) {
        alert(e);
        Router.push('/sign-in')
      }
  };

  return (
    <Container className='taskBoard'>
      <Row>
        <Col className="taskBoardHeader">
          <h2>Task Board</h2>
        </Col>
      </Row>
      <Row>
        <Col>
          <Lane taskList={taskList} fetchTaskList={fetchTaskList} />
        </Col>
      </Row>
    </Container>
  );
};

export default TaskBoard;