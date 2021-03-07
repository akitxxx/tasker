import React, { useState, useEffect, useRef } from 'react';
import Router from 'next/router';
import { Container, Row, Col, Button, Form } from 'react-bootstrap';
import ModalDialog from '../../components/modal';
import TaskModal from '../../components/task-modal';
import Lane from '../../components/lane';
import axios from 'axios';
import Layout from '../../components/layout';
import '../../styles/task-board.scss';

const TaskBoard = () => {

  const [laneList, setLaneList] = useState([]);
  const [showAddLaneModal, setShowAddLaneModal] = useState(false);
  const [showTaskModal, setShowTaskModal] = useState(false);
  const [targetTask, setTargetTask] = useState(null);
  const [laneName, setLaneName] = useState(null);
  const laneNameInput = useRef(null);

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
    e.target.blur();
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
                      fetchTaskList={fetchTaskList}
                      setTargetTask={setTargetTask}
                      setShowTaskModal={setShowTaskModal}
                      />
          })}
          <div><Button variant="default" className="btnAddLane" onClick={handleClickBtnAddLane}>+ Add lane</Button></div>
        </Row>
      </Container>
      <ModalDialog
        title="Add lane"
        btnExecName="Add lane"
        show={showAddLaneModal}
        onShow={() => laneNameInput.current.focus()}
        onHide={() => setShowAddLaneModal(false)}
        onCancel={() => setShowAddLaneModal(false)}
        onExec={() => createLane(laneName)}
      >
        <Form.Label>Lane name</Form.Label>
        <Form.Control type="text"
          ref={laneNameInput}
          onChange={(e) => {setLaneName(e.target.value)}}
          onKeyDown={(e) => e.keyCode === 13 && createLane(laneName)}/>
      </ModalDialog>
      <TaskModal
        show={showTaskModal}
        onHide={() => setShowTaskModal(false)}
        task={targetTask}
      >
      </TaskModal>
    </Layout>
  );
};

export default TaskBoard;