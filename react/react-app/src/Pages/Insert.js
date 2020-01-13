import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { useAuth } from "../context/auth";
import Button from '@material-ui/core/Button';

class Insert extends Component {

  render() {
    if (sessionStorage.getItem('access_token') === null) {
      return <Redirect to='/signin'/>
    }
    
    return (
      <div>
        <div>Insert Page - Teste</div>
        <Button onClick={localStorage.clear()}>Log out</Button>
      </div>
    );
  }
}

export default Insert;