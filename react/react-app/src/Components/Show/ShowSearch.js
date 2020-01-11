import React from 'react';
import { useState, /*useCallback,*/ useEffect } from 'react';
import PropTypes from 'prop-types';
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

function ShowSearch(props) {

    const { classes } = props;
    const [isError, setIsError] = useState(false);
    const [anchorEl, setAnchorEl] = useState(null);
    const [search, setSearch] = useState("");
    const [recipes, setRecipes] = useState([]);
    const [ingredient, setIngredient] = useState("");
    const [dataIngredients, setDataIngredients] = useState([]);
    const [searchByMeal, setSearchByMeal] = useState(false);
    const [meal, setMeal] = useState("");
    const [dataType, setDataType] = useState([{
        value: 'Breakfast & Brunch',
      }, {
        value: 'Lunch & Dinner',
      }, {
        value: 'Desert',
      }, {
        value: 'Appetizers & Snacks',
      },{
        value: 'Drinks',
      },]);
      

    useEffect(() => {

        getAll();

        //axios.get("http://192.168.1.68:8000/api/searchIngredientAll")
        axios.get("http://localhost:8000/api/searchIngredientAll")
        .then(resulti => {
            if (resulti.status==200) { 
              console.log(resulti.data)
                setDataIngredients([]);
                for (let ingObject of resulti.data) {
                    var ing= ingObject.ingredient_name;
                    setDataIngredients(dataIngredients =>[...dataIngredients, {value: ing}])
                }
            } else {
                setIsError(true);
            }
        }).catch(e => {
            setIsError(true);
        });
      }, []);

      const searchGet = () => {
        console.log(search)
        //axios.get("http://192.168.1.68:8000/api/searchRecipeName/name/"+search.text)
        axios.get("http://localhost:8000/api/searchRecipeName/name/"+search)
        .then(result => {
            console.log(result)
            if (result.status==200) { 
                setRecipes(result.data);
            } else {
                setIsError(true);
            }
        }).catch(e => {
            setIsError(true);
        });
    };

    const getAll = () => {
        //axios.get("http://192.168.1.68:8000/api/searchRecipeAll")
        axios.get("http://localhost:8000/api/searchRecipeAll")
        .then(result => {
            if (result.status==200) { 
                setRecipes(result.data);
                setSearchByMeal(false)
            } else {
                setIsError(true);
            }
        }).catch(e => {
            setIsError(true);
        });
    }

    const mealGet = (mealValue) => {
      console.log(mealValue)
        //axios.get("http://192.168.1.68:8000/api/searchRecipeCategory/category/"+meal.value)
        axios.get("http://localhost:8000/api/searchRecipeCategory/category/"+mealValue)
        .then(result => {
            console.log(result.data);
            if (result.status==200) { 
                setRecipes(result.data);
                setSearchByMeal(false)
            } else {
                setIsError(true);
            }
        }).catch(e => {
            setIsError(true);
        });
    };

    const ingredientGet = (ingredientValue) => {
      console.log(ingredientValue)
        //axios.get("http://192.168.1.68:8000/api/searchIngredientName/name/"+ingredient.value)
        axios.get("http://localhost:8000/api/searchRecipeNameTotal/name/"+ingredientValue) 
        .then(result => {
            console.log(result.data);
            if (result.status==200) { 
                setRecipes(result.data);
                setSearchByMeal(false)
            } else {
                setIsError(true);
            }
        }).catch(e => {
            setIsError(true);
        });
    };

    const handleClick = event => {
      setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
      setAnchorEl(null);
    };

    return (
      <Paper className={classes.paper}>
        <AppBar className={classes.searchBar} position="static" color="default" elevation={0}>
          <Toolbar>
            <Grid item xs={1}>
              <MenuIcon color="inherit" 
              aria-controls="customized-menu"
              aria-haspopup="true"
              variant="contained" 
              onClick={handleClick} 
              />
              <StyledMenu
                  id="customized-menu"
                  anchorEl={anchorEl}
                  keepMounted
                  open={Boolean(anchorEl)}
                  onClose={handleClose}>
                <Typography color="textSecondary" align="center">Meal Type ></Typography>
                {dataType.map((val, key)=>(
                  <StyledMenuItem key={key}>
                    <ListItemText primary={val.value} onClick={mealGet}/>    
                  </StyledMenuItem>
                ))}
                <Typography color="textSecondary" align="center">Ingredients ></Typography>
                {dataIngredients.map((val, key)=>(
                  <StyledMenuItem key={key}>
                    <ListItemText primary={val.value} onClick={ingredientGet}/>    
                  </StyledMenuItem>
                ))}
              </StyledMenu>
            </Grid>
            <Grid container spacing={1} alignItems="center">
              <Grid item xs={10}>
                <TextField
                  fullWidth
                  placeholder="Search by name"  //a pesquisa por ingredientes pode ser como este site usa https://www.allrecipes.com/
                  InputProps={{
                    disableUnderline: true,
                    className: classes.searchInput,
                  }}
                  onChange={e => {
                    setSearch(e.target.value);
                  }}
                />
              </Grid>
              <Grid item >
                <Button variant="outlined" onClick={searchGet} /*onClick(função de get de receitas)*/>
                  Search
                </Button>
              </Grid>
            </Grid>
          </Toolbar>
        </AppBar>
        <Content recipes={recipes}/>
      </Paper>
    );
}

ShowSearch.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ShowSearch);