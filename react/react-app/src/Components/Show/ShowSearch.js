import React, { Component } from 'react';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';
import MenuIcon from '@material-ui/icons/Menu';
import ListItemText from '@material-ui/core/ListItemText';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import Content from './ShowContent'
import axios from 'axios';

const styles = theme => ({
  paper: {
    maxWidth: 1200,
    margin: 'auto',
    overflow: 'hidden',
  },
  searchBar: {
    borderBottom: '1px solid rgba(0, 0, 0, 0.12)',
  },
  searchInput: {
    fontSize: theme.typography.fontSize,
  },
  block: {
    display: 'block',
  },
  addUser: {
    marginRight: theme.spacing(1),
  },
  contentWrapper: {
    margin: '40px 16px',
  },
});

const StyledMenu = withStyles({
  paper: {
    border: '1px solid #d3d4d5',
  },
})(props => (
  <Menu
    elevation={0}
    getContentAnchorEl={null}
    anchorOrigin={{
      vertical: 'bottom',
      horizontal: 'center',
    }}
    transformOrigin={{
      vertical: 'top',
      horizontal: 'center',
    }}
    {...props}
  />
));

const StyledMenuItem = withStyles(theme => ({
  root: {
    '&:focus': {
      backgroundColor: theme.palette.primary.main,
      '& .MuiListItemIcon-root, & .MuiListItemText-primary': {
        color: theme.palette.common.white,
      },
    },
  },
}))(MenuItem);

class ShowSearch extends Component {

  constructor(props){
    super(props);
    this.state = {
      isError: false,
      anchorEl: null,
      search: "",
      recipes: [],
      ingredient: "",
      dataIngredients: [],
      meal: "",
      dataType: [{
        value: 'Breakfast & Brunch',
      }, {
        value: 'Lunch & Dinner',
      }, {
        value: 'Dessert',
      }, {
        value: 'Appetizers & Snacks',
      },{
        value: 'Drinks',
      },]
    };
    this.searchGet=this.searchGet.bind(this);
    this.getAll=this.getAll.bind(this);
    this.mealGet=this.mealGet.bind(this);
    this.ingredientGet=this.ingredientGet.bind(this);
    this.handleClick=this.handleClick.bind(this);
    this.handleClose=this.handleClose.bind(this);
  }
      
  componentDidMount() {
    this.getAll();

    axios.get("http://localhost:8000/api/searchIngredientAll")
    .then(resulti => {
        if (resulti.status==200) { 
            this.setState({dataIngredients: []});
            for (let ingObject of resulti.data) {
                var ing= ingObject.ingredient_name;
                this.setState(prevState => ({
                  dataIngredients: [...prevState.dataIngredients, {value: ing}]
                }))
            }
        } else {
          this.setState({isError: true})
        }
    }).catch(e => {
      this.setState({isError: true})
    });
  }

  searchGet()
  {
    console.log(this.state.search)
    axios.get("http://localhost:8000/api/searchRecipeName/name/"+this.state.search)
    .then(result => {
        if (result.status==200) { 
          this.setState({recipes: result.data})
        } else {
          this.setState({isError: true})
        }
    }).catch(e => {
      this.setState({isError: true})
    });
  }

  getAll()
  {
    axios.get("http://localhost:8000/api/searchRecipeAll")
    .then(result => {
        if (result.status==200) { 
          this.setState({recipes: result.data})
        } else {
          this.setState({isError: true})
        }
    }).catch(e => {
      this.setState({isError: true})
    });
  }

  mealGet(mealValue)
  {
    console.log(mealValue)
    axios.get("http://localhost:8000/api/searchRecipeCategory/category/"+mealValue)
    .then(result => {
        console.log(result.data);
        if (result.status==200) { 
          this.setState({recipes: result.data})
          this.handleClose()
        } else {
          this.setState({isError: true})
        }
    }).catch(e => {
      this.setState({isError: true})
    });
  }

  ingredientGet(ingredientValue)
  {
    axios.get("http://localhost:8000/api/searchRecipeNameTotal/name/"+ingredientValue) 
    .then(result => {
        console.log(result);
        if (result.status==200) { 
          this.setState({recipes: result.data})
          this.handleClose()
        } else {
          this.setState({isError: true})
        }
    }).catch(e => {
      this.setState({isError: true})
    });
  }

  handleClick = event => {
    this.setState({anchorEl: event.currentTarget})
    };

  handleClose = () => {
    this.setState({anchorEl: null})
  };

  render() {
    const { classes } = this.props;

    return (
      <Paper className={classes.paper}>
        <AppBar className={classes.searchBar} position="static" color="default" elevation={0}>
          <Toolbar>
            <Grid item xs={1}>
              <MenuIcon color="inherit" 
              aria-controls="customized-menu"
              aria-haspopup="true"
              variant="contained" 
              onClick={this.handleClick} 
              />
              <StyledMenu
                  id="customized-menu"
                  anchorEl={this.state.anchorEl}
                  keepMounted
                  open={Boolean(this.state.anchorEl)}
                  onClose={this.handleClose}>
                <Typography color="textSecondary" align="center">Meal Type ></Typography>
                {this.state.dataType.map((val, key)=>(
                  <StyledMenuItem key={key}>
                    <ListItemText primary={val.value} onClick={() => this.mealGet(val.value)}/>
                  </StyledMenuItem>
                ))}
                <Typography color="textSecondary" align="center">Ingredients ></Typography>
                {this.state.dataIngredients.map((val, key)=>(
                  <StyledMenuItem key={key}>
                    <ListItemText primary={val.value} onClick={() => this.ingredientGet(val.value)}/>
                  </StyledMenuItem>
                ))}
              </StyledMenu>
            </Grid>
            <Grid container spacing={1} alignItems="center">
              <Grid item xs={10}>
                <TextField
                  fullWidth
                  placeholder="Search by name"  
                  onKeyPress={ (e) => {
                    if (e.key === 'Enter') {
                      console.log('Enter key pressed');
                      this.searchGet();
                    }
                  }}
                  InputProps={{
                    disableUnderline: true,
                    className: classes.searchInput,
                  }}
                  onChange={e => {
                    this.setState({
                      search: e.target.value
                    });  
                  }}
                />
              </Grid>
              <Grid item >
                <Button variant="outlined" onClick={this.searchGet}>
                  Search
                </Button>
              </Grid>
            </Grid>
          </Toolbar>
        </AppBar>
        <Content recipes={this.state.recipes}/>
      </Paper>
    );
  }
}

export default withStyles(styles)(ShowSearch);