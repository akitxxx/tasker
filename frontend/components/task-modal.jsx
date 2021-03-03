import React, { useState, useEffect } from 'react';
import { Form, Modal } from 'react-bootstrap';

const TaskModal = (props) => {

  return props.task ?
    (<Modal show={props.show} onHide={props.onHide}>
      <Modal.Header closeButton>
        <Modal.Title>{props.task.title}</Modal.Title>
      </Modal.Header>

      <Modal.Body>
        <Form.Label>Task deital</Form.Label>
        <Form.Control as="textarea" rows={3} />
      </Modal.Body>
    </Modal>)
    : <></>;
};

export default TaskModal;