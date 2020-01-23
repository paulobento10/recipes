import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { Redirect } from 'react-router-dom';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField'
import SelectIngredients from 'react-select';
import axios from 'axios'; 
import makeAnimated from 'react-select/animated';
import 'bootstrap/dist/css/bootstrap.min.css';
import InsertModal from './InsertModal';

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
  rootIng: {
    '& svg': {
      margin: theme.spacing(1),
    },
    '& hr': {
      margin: theme.spacing(1, 2.5),
    },
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
      kcal: "",
      category:"",
      ingredients: [],
      selectedOptions: [],
    };
    this.handleChange=this.handleChange.bind(this);
    this.handleCreate=this.handleCreate.bind(this);
    this.handleUpdateIngredients=this.handleUpdateIngredients.bind(this);
  }

  componentDidMount() {
    axios.get("http://localhost:8000/api/searchIngredientAll")
    .then(resulti => {
        if (resulti.status==200) { 
            resulti.data.forEach(element => {
                var ing= element.ingredient_name;
                var ing_id= element.ingredient_id;
                var ing_kcal=element.kcal;
                this.setState({ingredients: [...this.state.ingredients, {label: ing, value: ing_id, kcal: ing_kcal}]})
            });
        }
    })
  }

  handleUpdateIngredients(array){
    console.log("Handle Update Ingredients (Parant Component)");
    this.setState({ingredients: array});
  }

  handleChange = (selectedOptions) => {
    this.setState({ selectedOptions });
    const calorieStringTotal = selectedOptions.reduce((totalCalories, meal) => totalCalories + parseInt(meal.kcal, 10), 0);
    this.setState({kcal: calorieStringTotal.toString()})
  }

  handleCreate () {
      var recipe = {
        recipe_name: this.state.recipe_name,
        recipe_description: this.state.recipe_description,
        duration: this.state.duration,
        picture: this.state.picture,
        category: this.state.category,
        kcal: this.state.kcal,
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
  }

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
            onChange={e => {
              this.setState({
                picture: e.target.value
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

          <Grid container direction="row" alignItems="center" className={classes.rootIng}>
            <Grid item xs={11}>
              <SelectIngredients className={classes.select} onChange={this.handleChange} placeholder="Select Ingredients" isMulti options={this.state.ingredients} components={animatedComponents} />
            </Grid>
            <Grid item xs={1}>
              <InsertModal handleChangeArray={this.handleUpdateIngredients} ingredients={this.state.ingredients}/>
            </Grid>
          </Grid>
          
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