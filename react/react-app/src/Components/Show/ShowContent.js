import React, { Component } from "react";
import { withStyles } from '@material-ui/core/styles';
import "antd/dist/antd.css";
import { Pagination } from "antd";
import { Container } from "@material-ui/core";
import Grid from '@material-ui/core/Grid';
import { Card } from 'antd';

const { Meta } = Card;

const useStyles = theme => ({
  paper: {
    padding: theme.spacing(3, 2),
    verticalAlign: 'middle',
    margin: 'auto',
  },
  grid: {
    flexGrow: 1,
    padding: theme.spacing(3, 2),
    verticalAlign: 'middle',
    margin: 'auto',
  },
  container: {
    align: 'center',
    paddingTop: 15
  }
});

class ShowContent extends Component {

  constructor(props){
    super(props);
    this.state = {
      minValue: 0,
      maxValue: 9,
      redirectId: -1,
      spacing: 2,
    };
    this.handleChange=this.handleChange.bind(this);
  }

  handleChange = value => {
    this.setState({minValue: (value-1)*9})
    this.setState({maxValue: value * 9})
  }
 
  render() {
    const { classes } = this.props;

    return (
      <Grid container className={classes.grid} spacing={2}>
      <Grid item xs={12}>
        <Grid container justify="center" spacing={this.state.spacing}>
        {this.props.recipes &&
          this.props.recipes.length > 0 &&
          this.props.recipes.slice(this.state.minValue, this.state.maxValue).map((val, key) => (
            <a key={key} href={"/show/recipe/"+val.recipe_id}>
              <Card 
                hoverable
                style={{ width: 300, padding: 10}}
                cover={<img alt="example" src={val.picture} height={200} width={250}/>}
              >
                <Meta title={val.recipe_name} description={val.category}/>
              </Card>
            </a>
          ))}
        </Grid>
      </Grid>
      <Container className={classes.container}>
      <Pagination
          defaultCurrent={1}
          defaultPageSize={9}
          onChange={this.handleChange}
          total={this.props.recipes.length}
        />
        </Container>
      </Grid>
    );
  }
}

export default withStyles(useStyles)(ShowContent);