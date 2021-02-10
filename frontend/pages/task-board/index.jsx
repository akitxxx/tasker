import React, { useState } from 'react';
import Router from 'next/router';
import { Container, Row, Col, Button, Form } from 'react-bootstrap';
import ModalDialog from '../../components/modal';
import Lane from '../../components/lane';
import axios from 'axios';
import { useEffect } from 'react';
import Layout from '../../components/layout';
import '../../styles/task-board.scss';

const TaskBoard = () => {

  const [laneList, setLaneList] = useState([]);
  const [showAddLaneModal, setShowAddLaneModal] = useState(false);
  const [laneName, setLaneName] = useState(null);

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

  const handleClickBtnAddLane = (e) => {
    e.preventDefault();
    setShowAddLaneModal(true);
  };

  const createLane = async (name) => {
    const uri = '/api/create-lane';
    // get token from local storage
    const token = localStorage.getItem('tasker_token');
    try {
      const res = await axios.post(uri, {
        name: name,
      }, {
        headers: {
          Authorization: `Bearer ${token}`,
        }
      });
      res.data && setLaneList([...laneList, res.data]);
      fetchTaskList();
    } catch(e) {
      alert(e);
      return;
    }

    setShowAddLaneModal(false);
  };

  return (
    <Layout>
      <Container className='taskBoard'>
        <Row>
          {laneList && laneList.map((lane) => {
            return <Lane key={lane.id} id={lane.id} userId={lane.user_id} name={lane.name} taskList={lane.task_list || []}
                      fetchTaskList={fetchTaskList}/>
          })}
          <Col md={2}><Button variant="default" className="btnAddLane" onClick={handleClickBtnAddLane}>+ Add lane</Button></Col>
        </Row>
      </Container>
      <ModalDialog
        title="Add lane"
        btnExecName="Add lane"
        show={showAddLaneModal}
        onHide={() => setShowAddLaneModal(false)}
        onCancel={() => setShowAddLaneModal(false)}
        onExec={() => createLane(laneName)}
      >
        <Form.Label>Lane name</Form.Label>
        <Form.Control type="text" onChange={(e) => {setLaneName(e.target.value)}} />
      </ModalDialog>
    </Layout>
  );
};

export default TaskBoard;