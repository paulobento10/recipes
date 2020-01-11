import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
const axios = require('axios');

class Login extends Component {

  constructor(props){
    super(props);
    this.state = {
      toShow: false,
    };
    this.handleLogin=this.handleLogin.bind(this);
  }
  handleLogin(e)
  {
    e.preventDefault();
    if(this.refs.email.value === ""){
      alert("Email is required");
    }
    else if(this.refs.password.value === "") {
      alert("Password is required");
    }
    else{
      var user = {
        email: this.refs.email.value,
        password: this.refs.password.value,
      }
      console.log(user);
      axios.post('http://localhost:8000/api/login', user)
      .then(response => {
        console.log(response.data);
        if (response.data === -1)
        {
          alert("Password or username wrong !");
        }
        else if (response.data !== -1){// if (response.data.access_token){
          sessionStorage.setItem('access_token', response.data.access_token);
          this.setState({toShow: true});
        }
      })
      .catch(error => {
        alert(error);
      });
    }
  }
  render() {
    if (this.state.toShow === true) {
      console.log('entrei')
      return <Redirect to='/show' />
    }

    return (
      <div className="login">
      <h3>Login</h3>
      <form onSubmit={this.handleLogin}>
      <div className="login-form">
      <div className="control-group">
      <input type="text" className="logininput" ref="email" placeholder="Email" />
      <label className="login-field-icon fui-user" for="login-name"></label>
      </div>
      <div className="control-group">
      <input type="password" class="logininput" ref="password" placeholder="password"/>
      <label className="login-field-icon fui-lock" for="login-pass"></label>
      </div>
      <input className="btn btn-primary btn-large btn-block" type="submit" value="Login" />
      </div>
      </form>
      </div>

    );
  }
}

export default Login;