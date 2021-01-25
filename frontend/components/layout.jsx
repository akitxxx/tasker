import Router from 'next/router';
import { Button } from 'react-bootstrap';

const header = {
  height: '3rem',
  background: '#007bff',
  padding: '.2rem 1rem',
  color: 'white',
  marginBottom: '1rem',
};

const title = {
  fontSize: '1.8rem',
  fontWeight: 'bold',
};

const Layout = (props) => {

  const handleClickSignOut = (e) => {
    e.target.blur();
    localStorage.removeItem('tasker_token');
    Router.push('/sign-in');
  };

  return (
    <div>
      <div style={header}>
        <div className="d-inline" style={title}>Tasker</div>
        <Button className="float-right" onClick={handleClickSignOut}>Sign out</Button>
      </div>
      {props.children}
    </div>
  )
};


export default Layout;