import React, { Component } from 'react';
import Modal from 'react-modal';
import Grid from '@material-ui/core/Grid';
import AddIcon from '@material-ui/icons/Add';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField'
import axios from 'axios'; 

const customStyles = {
    content : {
      top                   : '50%',
      left                  : '50%',
      right                 : 'auto',
      bottom                : 'auto',
      marginRight           : '-50%',
      transform             : 'translate(-50%, -50%)'
    }
};

class InsertModal extends Component {

    constructor(props){
      super(props);
      this.state = {
        ingredient_name: "",
        kcal:"",
        user_id: 0,
        modalIsOpen: false,
        auxIngredients: []
      };
      this.handlePost=this.handlePost.bind(this);
      this.openModal = this.openModal.bind(this);
      this.afterOpenModal = this.afterOpenModal.bind(this);
      this.closeModal = this.closeModal.bind(this);
    }

    openModal() {
        this.setState({modalIsOpen: true});
    }

    afterOpenModal() {
        // references are now sync'd and can be accessed.
        this.subtitle.style.color = '#f00';
    }
    
    closeModal() {
        this.setState({modalIsOpen: false});
    }

    handlePost(){
        var ingredient = {
            ingredient_name: this.state.ingredient_name,
            kcal: this.state.kcal,
            user_id: parseInt(sessionStorage.getItem('access_token')) 
        }
        
        axios.post("http://localhost:8000/api/insertIngredient", ingredient)
        .then(result => {
            if (result.data==true) {
                axios.get("http://localhost:8000/api/searchIngredientName/name/"+ingredient.ingredient_name)
                .then(resulti => {
                    
                    this.setState({auxIngredients: this.props.ingredients}) //Preencher um array local com o array de ingredientes enviado pelo parent component
                    var ing= ingredient.ingredient_name;
                    var ing_id= resulti.data[0].ingredient_id
                    var ing_kcal=ingredient.kcal;
                    this.setState({auxIngredients: [...this.state.auxIngredients, {label: ing, value: ing_id, kcal: ing_kcal}]})    //adicionar ao array local o novo ingrediente 
                    this.props.handleChangeArray(this.state.auxIngredients) //chamar a função do parent component e enviar o array local (no parent component, esta função ira substituir o array de ingredientes por este array que lhe enviamos)
                    console.log(ingredient);
                    this.closeModal();
                })
            }
        })
    }

    render() {
        const { classes } = this.props;
    
        return (
          <Grid container direction="row" alignItems="center">
            <AddIcon onClick={this.openModal}/>
            <Modal
            isOpen={this.state.modalIsOpen}
            onAfterOpen={this.afterOpenModal}
            onRequestClose={this.closeModal}
            style={customStyles}
            >
                <h2 ref={subtitle => this.subtitle = subtitle}>Create Ingredient</h2>
                
                <div>Here you can create ingredients for your recipe</div>
                <form>
                    <TextField
                    ref="ingredient_name"
                    variant="outlined"
                    margin="normal"
                    fullWidth
                    id="ingredient_name"
                    label="Ingredient Name"
                    name="ingredient_name"
                    onChange={e => {
                        this.setState({
                            ingredient_name: e.target.value
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
                    onChange={e => {
                        this.setState({
                            kcal: e.target.value
                        });  
                    }}
                    />
                </form>

                <Grid container direction="row" alignItems="center">
                    <Grid item xs={10}>
                        <Button
                            fullWidth
                            variant="contained"
                            color="primary"
                            onClick={this.handlePost}
                        >
                            Create
                        </Button>
                    </Grid>
                    <Grid item xs={1}>
                        <Button onClick={this.closeModal}>Close</Button>
                    </Grid>
                </Grid>

            </Modal>
          </Grid>
        );
    }
}

export default InsertModal;