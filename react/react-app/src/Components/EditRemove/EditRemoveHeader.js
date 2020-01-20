import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import Toolbar from '@material-ui/core/Toolbar';
import { withStyles } from '@material-ui/core/styles';
import logo from '../images/image.png';

const lightColor = 'rgba(255, 255, 255, 0.7)';

const styles = theme => ({
  secondaryBar: {
    zIndex: 0,
  },
  menuButton: {
    marginLeft: -theme.spacing(1),
  },
  iconButtonAvatar: {
    padding: 4,
  },
  link: {
    textDecoration: 'none',
    color: lightColor,
    '&:hover': {
      color: theme.palette.common.white,
    },
  },
  button: {
    borderColor: lightColor,
  },
});

class EditRemoveHeader extends Component {
    constructor(props){
        super(props);
        this.state = {
            toSignIn: false,
        };
        this.handleLogOut=this.handleLogOut.bind(this);
    }

    handleLogOut()
    {
        sessionStorage.clear();
        this.setState({toSignIn: true});
    }

    render() {
    if (this.state.toSignIn === true) {
        return <Redirect to='/signin'/>
    }

    const { classes } = this.props;

    return (
      <React.Fragment>
        <AppBar
          component="div"
          className={classes.secondaryBar}
          position="static"
          elevation={0}
          style={{ backgroundColor: '#F06923' }}
        >
          <Toolbar>
            <Grid container alignItems="center">
              <Grid item xs>
              <a href="http://localhost:3000/show/">
                <img src={logo} alt="Logo"/>
              </a>
              </Grid>
              <Grid item>
                <Button className={classes.button} variant="outlined" color="inherit" size="medium" onClick={this.handleLogOut}> 
                  Log Out
                </Button>
              </Grid>
            </Grid>
          </Toolbar>
        </AppBar>
      </React.Fragment>
    );
  }
}

export default withStyles(styles)(EditRemoveHeader);