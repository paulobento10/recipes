import React, { Component } from 'react';
import { createMuiTheme, ThemeProvider, Text } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import { Redirect } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import ShowRecipe from '../Components/Show/ShowRecipe';
import Header from '../Components/Header';
import Copyright from "../Components/Copyright";
import axios from 'axios';

const useStyles = theme => ({
  root: {
    display: 'flex',
    minHeight: '100vh',
  },
  drawer: {
    [theme.breakpoints.up('sm')]: {
      width: 256,
      flexShrink: 0,
    },
  },
  app: {
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
  },
  main: {
    flex: 1,
    padding: theme.spacing(6, 4),
    background: '#eaeff1',
  },
  footer: {
    padding: theme.spacing(2),
    background: '#eaeff1',
  },
});

let theme = createMuiTheme({
  palette: {
    primary: {
      light: '#63ccff',
      main: '#009be5',
      dark: '#006db3',
    },
  },
  typography: {
    h5: {
      fontWeight: 500,
      fontSize: 26,
      letterSpacing: 0.5,
    },
  },
  shape: {
    borderRadius: 8,
  },
  props: {
    MuiTab: {
      disableRipple: true,
    },
  },
  mixins: {
    toolbar: {
      minHeight: 48,
    },
  },
});

theme = {
  ...theme,
  overrides: {
    MuiDrawer: {
      paper: {
        backgroundColor: '#18202c',
      },
    },
    MuiButton: {
      label: {
        textTransform: 'none',
      },
      contained: {
        boxShadow: 'none',
        '&:active': {
          boxShadow: 'none',
        },
      },
    },
    MuiTabs: {
      root: {
        marginLeft: theme.spacing(1),
      },
      indicator: {
        height: 3,
        borderTopLeftRadius: 3,
        borderTopRightRadius: 3,
        backgroundColor: theme.palette.common.white,
      },
    },
    MuiTab: {
      root: {
        textTransform: 'none',
        margin: '0 16px',
        minWidth: 0,
        padding: 0,
        [theme.breakpoints.up('md')]: {
          padding: 0,
          minWidth: 0,
        },
      },
    },
    MuiIconButton: {
      root: {
        padding: theme.spacing(1),
      },
    },
    MuiTooltip: {
      tooltip: {
        borderRadius: 4,
      },
    },
    MuiDivider: {
      root: {
        backgroundColor: '#404854',
      },
    },
    MuiListItemText: {
      primary: {
        fontWeight: theme.typography.fontWeightMedium,
      },
    },
    MuiListItemIcon: {
      root: {
        color: 'inherit',
        marginRight: 0,
        '& svg': {
          fontSize: 20,
        },
      },
    },
    MuiAvatar: {
      root: {
        width: 32,
        height: 32,
      },
    },
  },
};

class Recipe extends Component {

  constructor(props){
    super(props);
    this.state = {
      id: this.props.match.params.id,
      isError: false,
      recipe: [],
      ingredients: [],
    };
  }

  componentDidMount() {
    axios.get("http://localhost:8000/api/searchRecipe/id/"+this.state.id)
    .then(resulti => {
      if (resulti.status==200) {
        this.setState({recipe: resulti.data[0]});
        axios.get("http://localhost:8000/api/getIngredientsByRecipeId/id/"+resulti.data[0].recipe_id)
        .then(resultj => {
          console.log(resultj.data)
          if (resultj.status==200) {
            this.setState({ingredients: resultj.data});
          }
        });
      } 
      if(resulti.data.length<1) {
        alert("Error!");
        this.setState({isError: true})
      }
    });
  }

  render() {
    if (this.state.isError === true) {
      return <Redirect to='/show'/>
    }
    const { classes } = this.props;
    console.log(this.state.recipe);
    
    return (
      <ThemeProvider theme={theme}>
        <div className={classes.root}>
          <CssBaseline />
          <div className={classes.app}>
            <Header />
            <main className={classes.main}>
            <ShowRecipe recipe={this.state.recipe} ingredients={this.state.ingredients}/>
            </main>
            <footer className={classes.footer}>
              <Copyright />
            </footer>
          </div>
        </div>
      </ThemeProvider>
    );
  }
}

export default withStyles(useStyles)(Recipe);