import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField'
import Paper from '@material-ui/core/Paper';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Link from '@material-ui/core/Link';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import axios from 'axios';
import Copyright from "../Components/Copyright";
import { withStyles } from '@material-ui/core/styles';

const useStyles = theme => ({
  root: {
    height: '100vh',
  },
  image: {
    backgroundImage: 'url(https://www.evasoes.pt/files/2017/07/Moura-1200x800.jpg)',
    backgroundRepeat: 'no-repeat',
    backgroundColor:
      theme.palette.type === 'dark' ? theme.palette.grey[900] : theme.palette.grey[50],
    backgroundSize: 'cover',
    backgroundPosition: 'center',
  },
  paper: {
    margin: theme.spacing(8, 4),
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
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
});

class SignIn extends Component {

  constructor(props){
    super(props);
    this.state = {
      toShow: false,
      password: "",
      email:"",
    };
    this.handleLogin=this.handleLogin.bind(this);
  }

  handleLogin(e)
  {
    e.preventDefault();
    if(this.state.email === ""){
      alert("Email is required");
    }
    else if(this.state.password === "") {
      alert("Password is required");
    }
    else{
      var user = {
        email: this.state.email,
        password: this.state.password,
      }
      console.log(user);
      axios.post('http://localhost:8000/api/login', user)
      .then(response => {
        console.log(response.data);
        if (response.data === -1)
        {
          alert("Password or username wrong!");
        }
        else if (response.data !== -1){
          sessionStorage.setItem('access_token', response.data);
          this.setState({toShow: true});
          console.log(sessionStorage)
        }
      })
      .catch(error => {
        alert(error);
      });
    }
  }
  
  render() {
    const { classes } = this.props;
    if (this.state.toShow === true) {
      console.log('entrei')
      return <Redirect to='/show' />
    }

    return (
      <Grid container component="main" className={classes.root}>
        <Grid item xs={false} sm={4} md={7} className={classes.image} />
        <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
          <div className={classes.paper}>
            <Avatar className={classes.avatar}>
              <LockOutlinedIcon />
            </Avatar>
            <Typography component="h1" variant="h5">
              Sign in
            </Typography>
            <form className={classes.form} noValidate>
              <TextField
                ref="email"
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                autoFocus
                onChange={e => {
                  this.setState({
                    email: e.target.value
                  });  
                }}
              />
              <TextField
                ref="password"
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
                onChange={e => {
                  this.setState({
                    password: e.target.value
                  });  
                }}
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
                onClick={this.handleLogin}
              >
                Sign In
              </Button>
              <Grid container>
                <Grid item>
                  <Link href="/signup">Don't have an account? Sign Up</Link>
                </Grid>
              </Grid>
              <Box mt={5}>
                <Copyright />
              </Box>
            </form>
          </div>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(useStyles)(SignIn);