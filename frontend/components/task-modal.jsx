import React, { useState, useEffect, useRef } from 'react';
import Router from 'next/router';
import { Form, Modal } from 'react-bootstrap';
import axios from 'axios';
import './task-modal.module.scss';

const TaskModal = (props) => {

  const [showTitleInput, setShowTitleInput] = useState(false);
  const [showContentInput, setShowContentInput] = useState(false);
  const titleInput = useRef(null);
  const contentInput = useRef(null);

  useEffect(() => {
    showTitleInput && titleInput.current.focus();
  }, [showTitleInput]);

  useEffect(() => {
    showContentInput && contentInput.current.focus();
  }, [showContentInput]);

  const handleClickTitle = () => {
    setShowTitleInput(true);
  };

  const handleKeyDownTitleInput = async (e) => {
    if(e.keyCode === 13) {
      updateTask({title: e.target.value});
      setShowTitleInput(false);
    }
  };

  const handleBlurTitleInput = (e) => {
      updateTask({title: e.target.value});
    setShowTitleInput(false);
  };

  const handleClickContent = () => {
    setShowContentInput(true);
  };

  const handleKeyDownContentInput = async (e) => {
    if(e.keyCode === 13) {
      updateTask({content: e.target.value});
      setShowContentInput(false);
    }
  };

  const handleBlurContentInput = (e) => {
    updateTask({content: e.target.value});
    setShowContentInput(false);
  };

  const updateTask = async (task) => {
      const uri = '/api/update-task';
      // get token from local storage
      const token = localStorage.getItem('tasker_token');
      try {
        // get task list from server
        const res = await axios.patch(uri, {
          id: props.task.id,
          title: task.title,
          content: task.content,
        }, {
          headers: {
            Authorization: `Bearer ${token}`,
          }
        });
      } catch(e) {
        alert(e);
        Router.push('/sign-in')
      }

      props.fetchTaskList();
  };

  return props.task ?
    (<Modal show={props.show} onHide={props.onHide} animation={false}>
      <Modal.Header closeButton>
        <Modal.Title>
          {showTitleInput ?
            <Form.Control defaultValue={props.task.title} onKeyDown={handleKeyDownTitleInput} onBlur={handleBlurTitleInput} ref={titleInput} /> :
            <Form.Control defaultValue={props.task.title} plaintext readOnly onClick={handleClickTitle} className="cursor-pointer padding-input" />
          }
        </Modal.Title>
      </Modal.Header>

      <Modal.Body>
        {showContentInput ?
          <Form.Control
            as="textarea"
            rows={3}
            defaultValue={props.task.content}
            onKeyDown={handleKeyDownContentInput}
            onBlur={handleBlurContentInput}
            ref={contentInput}
          /> : <Form.Control
            as="textarea"
            className="cursor-pointer padding-input"
            defaultValue={props.task.content}
            plaintext
            readOnly
            onClick={handleClickContent}
          />
        }
      </Modal.Body>
    </Modal>)
    : <></>;
};

export default TaskModal;