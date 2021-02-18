import React from 'react';
import { Modal, Button } from 'react-bootstrap';

const ModalDialog = (props) => {
  
  return (
    <Modal show={props.show} onShow={props.onShow} onHide={props.onHide}>
      <Modal.Header closeButton>
        <Modal.Title>{props.title}</Modal.Title>
      </Modal.Header>

      <Modal.Body>
        {props.children}
      </Modal.Body>

      <Modal.Footer>
        <Button variant="secondary" onClick={props.onCancel}>Cancel</Button>
        <Button variant="primary" onClick={props.onExec}>{props.btnExecName}</Button>
      </Modal.Footer>
    </Modal>
  );
};

export default ModalDialog;