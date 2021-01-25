import React, { useState } from 'react';
import Router from 'next/router';
import '../../styles/common.scss';
import {Container, Row, Col, Form, Button} from 'react-bootstrap';
import axios from 'axios'
import { useEffect } from 'react';
import Layout from '../../components/layout';

const SignInPage = () => {

  const [needSignIn, setNeedsSignIn] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  useEffect(() => {
    const token = localStorage.getItem('tasker_token');
    if(token) {
      Router.push('/task-board');
    } else {
      setNeedsSignIn(true);
    }
  }, []);

  const handleEmailChange = e => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = e => {
    setPassword(e.target.value);
  };

  const handleKeyPress = e => {
    const ENTER = 13;
    if (e.keyCode === ENTER) {
      submit();
    }
  };

  const submit = async () => {
    const uri = '/api/sign-in';
    try {
      const res = await axios.post(uri, {
        email: email,
        password: password
      })

      // Set token to local storage
      localStorage.setItem('tasker_token', res.data)

      Router.push('/task-board')
    } catch (err) {
      alert('ログイン失敗\n' + err)
    }
  };

  return needSignIn ? (
    <Layout>
      <Container className='login'>
        <Row>
          <Col className='text-center my-3'>
            <h2>Sign in</h2>
          </Col>
        </Row>
        <Row>
          <Col>
            <Form className='col-md-6 offset-md-3 col-sm-8 offset-sm-2 col-xs-8 offset-xs-2'>
              <Form.Group>
                <Form.Control onChange={handleEmailChange} onKeyDown={handleKeyPress} type='text' name='id' placeholder='e-mail' />
              </Form.Group>
              <Form.Group>
                <Form.Control onChange={handlePasswordChange} onKeyDown={handleKeyPress} type='password' name='password' placeholder='password' />
              </Form.Group>
              <div className='text-center'>
                <Button onClick={submit} className='w-100'>Sign in</Button>
              </div>
            </Form>
          </Col>
        </Row>
      </Container>
    </Layout>
  ) : <></>;
};

export default SignInPage;
