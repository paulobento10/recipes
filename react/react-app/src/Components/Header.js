import React, { Component } from 'react';
import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import { Link } from 'react-router-dom';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';
import logo from '../images/image.png';
//import { withRouter } from 'react-router-dom';

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

class Header extends Component {

  render() {
    const { classes } = this.props;
    const imageClick = () => {
      window.location.reload();
    } 

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
              <a href={window.location.href}>
                <img src={logo} alt="Logo"/>
              </a>
              </Grid>
              <Grid item>
                <Button className={classes.button} variant="outlined" color="inherit" size="medium" href="/insert">
                  Create a recipe
                </Button>
              </Grid>
            </Grid>
          </Toolbar>
        </AppBar>
      </React.Fragment>
    );
  }
}

export default withStyles(styles)(Header);