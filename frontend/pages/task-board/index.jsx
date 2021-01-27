import React, { useState } from 'react';
import Router from 'next/router';
import { Container, Row, Col } from 'react-bootstrap';
import Lane from '../../components/lane';
import axios from 'axios';
import { useEffect } from 'react';
import Layout from '../../components/layout';
import '../../styles/task-board.scss';

const TaskBoard = () => {

  const [laneList, setLaneList] = useState([]);

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
        res.data ? setLaneList(res.data) : setLaneList([]);
      } catch(e) {
        alert(e);
        Router.push('/sign-in')
      }
  };

  return (
    <Layout>
      <Container className='taskBoard'>
        <Row>
          <Col>
            {laneList && laneList.map((lane) => {
              return <Lane key={lane.id} id={lane.id} userId={lane.user_id} name={lane.name} taskList={lane.task_list}
                        fetchTaskList={fetchTaskList}/>
            })}
          </Col>
        </Row>
      </Container>
    </Layout>
  );
};

export default TaskBoard;