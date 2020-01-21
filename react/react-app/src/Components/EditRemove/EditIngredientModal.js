import React, { Component } from 'react';
import Modal from 'react-modal';
import Grid from '@material-ui/core/Grid';
import EditIcon from '@material-ui/icons/Edit';
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

class EditIngredientModal extends Component {

    constructor(props){
      super(props);
      this.state = {
        ingredient_name: "",
        modalIsOpen: false
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
            ingredient_id: this.props.ingredient.ingredient_id,
            ingredient_name: this.state.ingredient_name,
        }
        
        axios.post("http://localhost:8000/api/editIngredientName", ingredient)
        .then(result => {
            console.log(result);
            if (result.data==true) {
                window.location.reload();
            }
        })
    }

    render() {
        const { classes } = this.props;
    
        return (
          <Grid container direction="row" alignItems="center">
            <EditIcon onClick={this.openModal}/>
            <Modal
            isOpen={this.state.modalIsOpen}
            onAfterOpen={this.afterOpenModal}
            onRequestClose={this.closeModal}
            style={customStyles}
            >
            <h2 ref={subtitle => this.subtitle = subtitle}>Editing Ingredient {this.props.ingredient.ingredient_name}</h2>
                
                <div>Here you can edit your ingredient</div>
                <form>
                    <TextField
                    ref="ingredient_name"
                    variant="outlined"
                    margin="normal"
                    fullWidth
                    id="ingredient_name"
                    label="Ingredient Name"
                    name="ingredient_name"
                    defaultValue={this.props.ingredient.ingredient_name}
                    onChange={e => {
                        this.setState({
                            ingredient_name: e.target.value
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
                            Edit
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

export default EditIngredientModal;