import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import Paper from '@material-ui/core/Paper';
import DeleteIcon from '@material-ui/icons/Delete';
import { Table } from 'antd';
import { Typography } from '@material-ui/core';
import EditRecipeModal from './EditRecipeModal';
import EditIngredientModal from './EditIngredientModal';
import axios from 'axios'; 

class EditRemoveContent extends Component {

  constructor(props){
    super(props);
    this.state = {
      page: 0,
      rowsPerPage: 10,
      columnsRecipes: [
        {
          title: 'Recipe Name',
          dataIndex: 'recipe_name',
          key: 'name',
        },
        /*{
          title: 'Description',
          dataIndex: 'recipe_description',
          key: 'recipe_description',
        },*/
        {
          title: 'Duration',
          dataIndex: 'duration',
          key: 'duration',
        },
        {
          title: 'Category',
          dataIndex: 'category',
          key: 'category',
        },
        {
          title: 'Edit',
          key: 'edit',
          render: (text, record) => (
                <EditRecipeModal recipe={record}/>
          ),
        },
        {
          title: 'Delete',
          key: 'delete',
          render: (text, record) => (
                <DeleteIcon onClick={() => this.handleDeleteRecipe(record.recipe_id)}/>
          ),
        },
      ],
      dataRecipes: [],
      columnsIngredients: [
          {
            title: 'Ingredient Name',
            dataIndex: 'ingredient_name',
            key: 'ingredient_name',
          },
          {
            title: 'Calories',
            dataIndex: 'kcal',
            key: 'kcal',
          },
          {
            title: 'Edit',
            key: 'edit',
            render: (text, record) => (
              <span>
                <EditIngredientModal ingredient={record}/>
              </span>
            ),
          },
          {
            title: 'Delete',
            key: 'delete',
            render: (text, record) => (
              <span>
                <DeleteIcon onClick={() => this.handleDeleteIngredient(record.ingredient_id)}/>
              </span>
            ),
          },
        ],
        dataIngredients: [],
      };
      this.handleGetIngredients=this.handleGetIngredients.bind(this);
      this.handleGetRecipes=this.handleGetRecipes.bind(this);
      this.handleChangePage=this.handleChangePage.bind(this);
      this.handleChangeRowsPerPage=this.handleChangeRowsPerPage.bind(this);
      this.handleDeleteRecipe=this.handleDeleteRecipe.bind(this);
      this.handleDeleteIngredient=this.handleDeleteIngredient.bind(this);
    }

    componentDidMount() {
      this.handleGetRecipes()
      this.handleGetIngredients()
    }

    handleGetRecipes(){
      axios.get("http://localhost:8000/api/searchUserRecipe/id/"+sessionStorage.getItem('access_token'))
      .then(resulti => {
          if (resulti.status==200) { 
              this.setState({dataRecipes: resulti.data});
          }
      })
    }

    handleGetIngredients(){
      axios.get("http://localhost:8000/api/getIngredientByUserIdRoute/id/"+sessionStorage.getItem('access_token'))
      .then(resulti => {
          if (resulti.status==200) { 
              this.setState({dataIngredients: resulti.data});
          }
      })
    }

    //r.HandleFunc("/api/deleteRecipe/id/{id}", deleteRecipeRoute).Methods("DELETE")
    handleDeleteRecipe = key => {
      axios.delete('http://localhost:8000/api/deleteRecipe/id/'+key)
      .then(resulti => {
        if(resulti.status==200){
          for (let i = 0; i < this.state.dataRecipes.length; i++) {
            if(this.state.dataRecipes[i].recipe_id==key){
              console.log(this.state.dataRecipes[i]);
              var array = [...this.state.dataRecipes]; // make a separate copy of the array
              array.splice(i, 1);
              this.setState({dataRecipes: array});
            }
          }
        }
      })      
    }

    //r.HandleFunc("/api/deleteIngredient/id/{id}", deleteIngredientRoute).Methods("DELETE")
    handleDeleteIngredient = key => {
      axios.delete('http://localhost:8000/api/deleteIngredient/id/'+key)
      .then(resulti => {
        if(resulti.status==200){
          for (let i = 0; i < this.state.dataIngredients.length; i++) {
            if(this.state.dataIngredients[i].ingredient_id==key){
              console.log(this.state.dataIngredients[i]);
              var array = [...this.state.dataIngredients]; // make a separate copy of the array
              array.splice(i, 1);
              this.setState({dataIngredients: array});
            }
          }
        }
      })
    }

    handleChangePage = (event, newPage) => {
        this.setState({page: newPage})
    };

    handleChangeRowsPerPage = event => {
        this.setState({rowsPerPage: +event.target.value})
        this.setState({page: 0})
    };

    render() {
    const { classes } = this.props;

    if (sessionStorage.getItem('access_token') < 0) {
      return <Redirect to='/signin'/>
    }

    return (
        <Paper>
          <Typography style={{textAlign: 'left', fontSize: 40}}>Recipes:</Typography>
          <Table columns={this.state.columnsRecipes} dataSource={this.state.dataRecipes} />

          <Typography style={{textAlign: 'left', fontSize: 40, paddingTop: '20'}}>Ingredients:</Typography>
          <Table columns={this.state.columnsIngredients} dataSource={this.state.dataIngredients} />
        </Paper>
      );
    }
}

export default EditRemoveContent;