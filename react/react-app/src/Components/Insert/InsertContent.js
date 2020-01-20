import React, { Component } from 'react';
import { createMuiTheme, ThemeProvider, withStyles } from '@material-ui/core/styles';
import { Redirect } from 'react-router-dom';
import CssBaseline from '@material-ui/core/CssBaseline';
import Avatar from '@material-ui/core/Avatar';
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField'
import SelectIngredients from 'react-select';
import Paper from '@material-ui/core/Paper';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Link from '@material-ui/core/Link';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import axios from 'axios'; 
import makeAnimated from 'react-select/animated';
import 'bootstrap/dist/css/bootstrap.min.css';

const animatedComponents = makeAnimated();

const useStyles = theme => ({
  root: {
    display: 'flex',
    minHeight: '100vh',
  },
  form: {
    flexGrow: 1,
    padding: theme.spacing(3, 2),
    verticalAlign: 'middle',
    margin: 'auto',
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
  formControl: {
    width: '100%',
    padding: theme.spacing(2,0),
  },
  select: {
    width: '100%',
    padding: theme.spacing(1,0),
  },
});

class InsertContent extends Component {

  constructor(props){
    super(props);
    this.state = {
      toShow: false,
      recipe_name: "",
      recipe_description:"",
      duration:"",
      picture:"",
      category:"",
      ingredients: [],
      selectedOptions: [],
      Countries: [
        { label: "Albania", value: 355 },
        { label: "Argentina", value: 54 },
        { label: "Austria", value: 43 },
        { label: "Cocos Islands", value: 61 },
        { label: "Kuwait", value: 965 },
        { label: "Sweden", value: 46 },
        { label: "Venezuela", value: 58 }
      ],
    };
    this.handleChange=this.handleChange.bind(this);
    this.handleCreate=this.handleCreate.bind(this);
  }

  componentDidMount() {
    axios.get("http://localhost:8000/api/searchIngredientAll")
    .then(resulti => {
        if (resulti.status==200) { 
            resulti.data.forEach(element => {
                var ing= element.ingredient_name;
                var ing_id= element.ingredient_id;
                this.setState({ingredients: [...this.state.ingredients, {label: ing, value: ing_id}]})
            });
        }
    })
  }

  handleChange = (selectedOptions) => {
    this.setState({ selectedOptions });
  }

  handleCreate () {
      var recipe = {
        recipe_name: this.state.recipe_name,
        recipe_description: this.state.recipe_description,
        duration: this.state.duration,
        picture: this.state.picture,
        category: this.state.category,
        kcal: "0",
        user_id: parseInt(sessionStorage.getItem('access_token')) 
      }
      
      console.log(recipe);
      
      axios.post("http://localhost:8000/api/insertRecipe", recipe)
      .then(result => {
        if (result.data==true) {
            axios.get("http://localhost:8000/api/searchRecipeExactName/name/" + this.state.recipe_name)
            .then(resulti => {
              if (resulti.status==200) { 
  
                this.state.selectedOptions.forEach(element => {
                    const elementIng = {
                      ingredient_id: element.value,
                      recipe_id: resulti.data[0].recipe_id,
                    }
                    axios.post("http://localhost:8000/api/insertRecipeIngredients", elementIng)
                    .then(resultj => {  
                        
                      if (resultj.data==true) {
                        console.log('Success inserting ingredients');
                      } 
                    })
                });
                this.setState({toShow: true})
              }
            })
          }
      })
  };

  render() {
    const { classes } = this.props;

    if (this.state.toShow === true) {
        return <Redirect to='/show' />
    }

    return (
        <form className={classes.form}>
          <TextField
            ref="recipe_name"
            variant="outlined"
            margin="normal"
            fullWidth
            id="recipe_name"
            label="Recipe Name"
            name="recipe_name"
            autoFocus
            onChange={e => {
              this.setState({
                recipe_name: e.target.value
              });  
            }}
          />
          <TextField
            ref="recipe_description"
            variant="outlined"
            margin="normal"
            fullWidth
            multiline
            rows="4"
            id="recipe_description"
            label="Description"
            name="recipe_description"
            autoFocus
            onChange={e => {
              this.setState({
                recipe_description: e.target.value
              });  
            }}
          />
          <TextField
            ref="duration"
            variant="outlined"
            margin="normal"
            fullWidth
            id="duration"
            label="Duration"
            name="duration"
            type="number"
            autoFocus
            onChange={e => {
              this.setState({
                duration: e.target.value
              });  
            }}
          />
          <TextField
            ref="picture"
            variant="outlined"
            margin="normal"
            fullWidth
            id="picture"
            label="Picture URL"
            name="picture"
            autoFocus
            onChange={e => {
              this.setState({
                picture: e.target.value
              });  
            }}
          />
          <TextField
            ref="kcal"
            variant="outlined"
            margin="normal"
            fullWidth
            id="kcal"
            label="Kcal"
            name="kcal"
            type="number"
            autoFocus
            onChange={e => {
              this.setState({
                kcal: e.target.value
              });  
            }}
          />  
          <FormControl className={classes.formControl}>
            <InputLabel id="demo-simple-select-label">Category</InputLabel>
            <Select
            variant="outlined"
            labelId="demo-simple-select-helper-label"
            id="demo-simple-select-helper"
            value={this.state.category}
            onChange={e => {
                this.setState({
                  category: e.target.value
                });  
              }}
            >
                <MenuItem value={'Breakfast & Brunch'}>{'Breakfast & Brunch'}</MenuItem>
                <MenuItem value={'Lunch & Dinner'}>{'Lunch & Dinner'}</MenuItem>
                <MenuItem value={'Dessert'}>Dessert</MenuItem>
                <MenuItem value={'Appetizers & Snacks'}>{'Appetizers & Snacks'}</MenuItem>
                <MenuItem value={'Drinks'}>Drinks</MenuItem>
            </Select>
          </FormControl>
          
          <SelectIngredients className={classes.select} onChange={this.handleChange} placeholder="Select Ingredients" isMulti options={this.state.ingredients} components={animatedComponents} />
        
          <Button
            fullWidth
            variant="contained"
            color="primary"
            onClick={this.handleCreate}
          >
            Create
          </Button>
        </form>
    );
  }
}

export default withStyles(useStyles)(InsertContent);