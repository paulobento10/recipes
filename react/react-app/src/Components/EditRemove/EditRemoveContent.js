import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import { Table, Divider, Tag } from 'antd';
import { Typography } from '@material-ui/core';
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
          title: 'Actions',
          key: 'actions',
          render: (text, record) => (
            <span>
              <a onClick={() => this.handleEditRecipeName(record.recipe_id)}>Edit</a>
              <Divider type="vertical" />
              <a onClick={() => this.handleDeleteRecipe(record.recipe_id)}>Delete</a> 
            </span>
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
            title: 'Actions',
            key: 'actions',
            render: (text, record) => (
              <span>
                <a onClick={() => this.handleEditIngredientName(record.ingredient_id)}>Edit</a>
                <Divider type="vertical" />
                <a onClick={() => this.handleDeleteIngredient(record.ingredient_id)}>Delete</a> 
              </span>
            ),
          },
        ],
        dataIngredients: [],
      };
      this.handleChangePage=this.handleChangePage.bind(this);
      this.handleChangeRowsPerPage=this.handleChangeRowsPerPage.bind(this);
      this.handleDeleteRecipe=this.handleDeleteRecipe.bind(this);
      this.handleDeleteIngredient=this.handleDeleteIngredient.bind(this);
      this.handleEditRecipeName=this.handleEditRecipeName.bind(this);
      this.handleEditIngredientName=this.handleEditIngredientName.bind(this);
    }

    componentDidMount() {
      axios.get("http://localhost:8000/api/getIngredientByUserIdRoute/id/"+sessionStorage.getItem('access_token'))
      .then(resulti => {
          if (resulti.status==200) { 
              this.setState({dataIngredients: resulti.data});
          }
      })

      axios.get("http://localhost:8000/api/searchUserRecipe/id/"+sessionStorage.getItem('access_token'))
      .then(resulti => {
          if (resulti.status==200) { 
              this.setState({dataRecipes: resulti.data});
          }
      })
    }

    //r.HandleFunc("/api/deleteRecipe/id/{id}", deleteRecipeRoute).Methods("DELETE")
    handleDeleteRecipe = key => {
      axios.delete('http://localhost:8000/api/deleteRecipe/id/'+key)
      .then(resulti => {
          console.log(resulti);
          const dataRecipes = [...this.state.dataRecipes];
          this.setState({ dataRecipes: dataRecipes.filter(item => item.key !== key) });
      })
    }

    //r.HandleFunc("/api/editRecipeName", editRecipeNameRoute).Methods("POST")
    handleEditRecipeName = key => {
      alert("Edit Recipe not done yet!");
    }

    //r.HandleFunc("/api/deleteIngredient/id/{id}", deleteIngredientRoute).Methods("DELETE")
    handleDeleteIngredient = key => {
      axios.delete('http://localhost:8000/api/deleteIngredient/id/'+key)
      .then(resulti => {
          console.log(resulti);
          const dataIngredients = [...this.state.dataIngredients];
          this.setState({ dataIngredients: dataIngredients.filter(item => item.key !== key) });
      })
      
    }

    //r.HandleFunc("/api/editIngredientName", editIngredientNameRoute).Methods("POST")
    handleEditIngredientName = key => {
      alert("Edit Ingredient not done yet");
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