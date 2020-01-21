import React, { Component } from 'react';
import { withStyles } from '@material-ui/core/styles';
import { Image } from 'react-bootstrap';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import QueryBuilderIcon from '@material-ui/icons/QueryBuilder';
import EqualizerIcon from '@material-ui/icons/Equalizer';
import Divider from '@material-ui/core/Divider';
import KeyboardArrowRightOutlinedIcon from '@material-ui/icons/KeyboardArrowRightOutlined';

const useStyles = theme => ({
  grid: {
    flexGrow: 1,
    padding: theme.spacing(3, 2),
    verticalAlign: 'middle',
    margin: 'auto',
  },
  gridText: {
    flexGrow: 1,
    padding: theme.spacing(6, 6),
  },
  paper: {
    maxWidth: 1200,
    margin: 'auto',
    overflow: 'hidden',
  },
  root: {
    width: 'fit-content',
    borderRadius: theme.shape.borderRadius,
    backgroundColor: theme.palette.background.paper,
    color: theme.palette.text.secondary,
    '& svg': {
      margin: theme.spacing(2),
    },
    '& hr': {
      margin: theme.spacing(0, 0.5),
    },
  },
  rootIng: {
    width: 'fit-content',
    '& svg': {
      margin: theme.spacing(2),
    },
    '& hr': {
      margin: theme.spacing(0, 0.5),
    },
  },
});

class Recipe extends Component {

  render() {
    const { classes } = this.props;
    
    return (
      <Paper className={classes.paper}>
        <Grid container justify="center">
            <Grid item xs={4} className={classes.gridText}>
                <Typography style={{textAlign: 'left', fontSize: 40, fontWeight: 'bold'}}>{this.props.recipe.recipe_name}</Typography>
                <Typography style={{textAlign: 'left', fontSize: 25, color: '#787878'}}>{this.props.recipe.category}</Typography>
            </Grid>
            <Grid item xs={8} className={classes.grid}>
                <Image src={this.props.recipe.picture} fluid />
            </Grid>
            <Grid container justify="center" direction="row" alignItems="center" className={classes.root}>
                <Grid item>
                    <Typography> {this.props.recipe.duration} m</Typography>
                </Grid>
                <Grid item>
                    <QueryBuilderIcon />
                </Grid>
                <Divider light orientation="vertical" />
                <Grid item>
                    <EqualizerIcon />
                </Grid>
                <Grid item>
                    <Typography> {this.props.recipe.kcal} kcal{"\n"}</Typography>
                </Grid>
            </Grid>

            <Grid item xs={10}>
                <Typography style={{textAlign: 'left', fontSize: 30}}>Ingredients:</Typography>
                <Divider />
                {this.props.ingredients.map(val => (
                    <Grid container direction="row" alignItems="center" className={classes.rootIng}>
                        <Grid item>
                            <KeyboardArrowRightOutlinedIcon />
                        </Grid>
                        <Grid item>
                            <Typography style={{textAlign: 'left', fontSize: 18}}>{val.ingredient_name}</Typography>
                        </Grid>
                    </Grid>
                ))}
            </Grid>

            <Grid item xs={10} style={{paddingTop: 15}}>
                <Typography style={{textAlign: 'left', fontSize: 30}}>Description:</Typography>
                    <Divider />
                <Typography style={{textAlign: 'left', fontSize: 20, padding: 6}}> {this.props.recipe.recipe_description}</Typography>
            </Grid>
        </Grid>
      </Paper>
    );
  }
}

export default withStyles(useStyles)(Recipe);