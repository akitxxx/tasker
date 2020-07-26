import React, { useState } from 'react';

import {Container, Row, Col} from 'react-bootstrap';

const TaskBoard = () => {

  return (
    <Container className='task-board'>
      <Row>
        <Col className='text-center '>
          <h2>Task Board</h2>
        </Col>
      </Row>
    </Container>
  );
};

export default TaskBoard;