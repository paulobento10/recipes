import React, { Component } from "react";
import { useState, /*useCallback,*/ useEffect } from 'react';
import "antd/dist/antd.css";
import { Pagination } from "antd";
import { Paper, Container } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { Card } from 'antd';
import Typography from '@material-ui/core/Typography';

const { Meta } = Card;

const useStyles = makeStyles(theme => ({
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
}));

function ShowContent(props) {
  const classes = useStyles();
  const [minValue, setMinValue] = useState(0);
  const [maxValue, setMaxValue] = useState(9);
  const [spacing, setSpacing] = React.useState(2);

  const handleChange = value => {
    setMinValue(((value-1)*9));
    setMaxValue(value * 9);
  };

  return (
    <Grid container className={classes.grid} spacing={2}>
    <Grid item xs={12}>
      <Grid container justify="center" spacing={spacing}>
      {props.recipes &&
        props.recipes.length > 0 &&
        props.recipes.slice(minValue, maxValue).map((val, key) => (
          <Card 
            key={key}
            hoverable
            style={{ width: 300, padding: 10}}
            cover={<img alt="example" src={val.picture} height={200} width={250} />}
          >
            <Meta title={val.recipe_name} description={val.category} />
          </Card>
          
        ))}
        
      </Grid>
    </Grid>
    <Container className={classes.container}>
    <Pagination
        defaultCurrent={1}
        defaultPageSize={9}
        onChange={handleChange}
        total={props.recipes.length}
      />
      </Container>
    </Grid>
  );
}

export default ShowContent;
