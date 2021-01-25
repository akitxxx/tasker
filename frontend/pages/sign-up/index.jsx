import React, { useState } from 'react';
import {Container, Row, Col, Form, Button} from 'react-bootstrap';
import axios from 'axios';
import Layout from '../../components/layout';
import Router from 'next/router';

const SignUpPage = (props) => {
  const [params, setParams] = useState({});

  const handleChange = e => {
    setParams({ ...params, [e.target.name]: e.target.value });
  };

  const handleKeyPress = e => {
    const ENTER = 13;
    if (e.keyCode === ENTER) {
      submit();
    }
  };

  const submit = async () => {
    const uri = '/api/sign-up';
    try {
      const res = await axios.post(uri, params);
      alert( 'User is registered.\n' + res.data.email);

      Router.push('/task-board');
    } catch(e) {
      alert(e);
    }
  };

  return (
    <Layout>
      <Container className='sign-up'>
        <Row>
          <Col className='text-center my-3'>
            <h2>Sign up</h2>
          </Col>
        </Row>
        <Row>
          <Col>
            <Form className='col-md-6 offset-md-3 col-sm-8 offset-sm-2 col-xs-8 offset-xs-2'>
              <Form.Group>
                <Form.Control onChange={handleChange} onKeyDown={handleKeyPress} type='text' name='email' placeholder='Email' autoComplete='off' />
              </Form.Group>
              <Form.Group>
                <Form.Control onChange={handleChange} onKeyDown={handleKeyPress} type='password' name='password' placeholder='Password' autoComplete='off' />
              </Form.Group>
              <Form.Group>
                <Form.Control onChange={handleChange} onKeyDown={handleKeyPress} type='password' name='confirmPassword' placeholder='Confirm password' autoComplete='off' />
              </Form.Group>
              <div className='text-center'>
                <Button onClick={submit} className='w-100'>Sign up</Button>
              </div>
            </Form>
          </Col>
        </Row>
      </Container>
    </Layout>
  );
};

export default SignUpPage;
