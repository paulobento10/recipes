import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import axios from 'axios';

class Recipe extends Component {

  constructor(props){
    super(props);
    this.state = {
      id: this.props.match.params.id,
      isError: false,
      recipe: [],
    };
  }

  componentDidMount() {
    axios.get("http://localhost:8000/api/searchRecipe/id/"+this.state.id)
    .then(resulti => {
      if (resulti.status==200) { 
        this.setState({recipe: resulti.data[0]});
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
    
    
    return (
      <div>
        <h3>Show recipe test!</h3>
        <h3>Recipe ID: {this.state.recipe.recipe_id}</h3>
        <h3>Recipe Name: {this.state.recipe.recipe_name}</h3>
      </div>
    );
  }
}

export default Recipe;