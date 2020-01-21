import React, { Component } from "react";
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import axios from 'axios';
import Copyright from "../Components/Copyright";
import { Redirect } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';

const useStyles = theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(3),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
});

class SignUp extends Component {

  constructor(props){
    super(props);
    this.state = {
      registered: false,
      name: "",
      password: "",
      email:"",
    };
    this.handleRegister=this.handleRegister.bind(this);
  }

  handleRegister(e)
  {
    e.preventDefault();
    if(this.state.email === ""){
      alert("Email is required");
    }
    else if(this.state.password === "") {
      alert("Password is required");
    }
    else if(this.state.name === "") {
      alert("Username is required");
    }
    else{
      var user = {
        user_name: this.state.name,
        email: this.state.email,
        password: this.state.password,
      }
      console.log(user);
      axios.post('http://localhost:8000/api/insertUser', user)
      .then(response => {
        console.log(response.data);
        if (response.data === false)
        {
          alert("Something went wrong!");
        }
        else if (response.data !== false){
          this.setState({registered: true});
        }
      })
      .catch(error => {
        alert(error);
      });
    }
  }

  render() {
    const { classes } = this.props;
    if (this.state.registered === true) {
      console.log('registou')
      return <Redirect to='/signin' />
    }

    return (
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <div className={classes.paper}>
          <Avatar className={classes.avatar}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign up
          </Typography>
          <form className={classes.form} noValidate>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  name="name"
                  variant="outlined"
                  required
                  fullWidth
                  id="name"
                  label="Name"
                  onChange={e => {
                    this.setState({
                      name: e.target.value
                    });  
                  }}
                  autoFocus
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  id="email"
                  label="Email Address"
                  name="email"
                  onChange={e => {
                    this.setState({
                      email: e.target.value
                    });  
                  }}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  onChange={e => {
                    this.setState({
                      password: e.target.value
                    });  
                  }}
                />
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              className={classes.submit}
              onClick={this.handleRegister}
            >
              Sign Up
            </Button>
            <Grid container justify="flex-end">
              <Grid item>
                <Link href="/signin" variant="body2">Already have an account? Sign In</Link>
              </Grid>
            </Grid>
          </form>
        </div>
        <Box mt={5}>
          <Copyright />
        </Box>
      </Container>
    );
  }
}

export default withStyles(useStyles)(SignUp);