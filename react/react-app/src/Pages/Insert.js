import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { useAuth } from "../context/auth";
import Button from '@material-ui/core/Button';

class Insert extends Component {

  constructor(props){
    super(props);
    this.state = {
      toShow: false,
    };
    this.handleLogOut=this.handleLogOut.bind(this);
  }

  handleLogOut()
  {
    sessionStorage.clear();
    this.setState({toShow: true});
  }

  render() {
    if (this.state.toShow === true) {
      return <Redirect to='/signin'/>
    }
    
    return (
      <div>
        <div>Insert Page - Teste</div>
        <Button onClick={this.handleLogOut}>Log out</Button>
      </div>
    );
  }
}

export default Insert;