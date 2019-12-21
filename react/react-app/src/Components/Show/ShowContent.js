import React, { Component } from "react";
import { useState, /*useCallback,*/ useEffect } from 'react';
import "antd/dist/antd.css";
import { Card, Pagination } from "antd";
import { Paper } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';

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
}));

function ShowContent(props) {
  const classes = useStyles();
  const [minValue, setMinValue] = useState(0);
  const [maxValue, setMaxValue] = useState(5);
  const [spacing, setSpacing] = React.useState(2);

  const handleChange = value => {
    setMinValue(((value-1)*5));
    setMaxValue(value * 5);
  };

  return (
    <Grid container className={classes.grid} spacing={2}>
    <Grid item xs={12}>
      <Grid container justify="center" spacing={spacing}>
      {props.recipes &&
        props.recipes.length > 0 &&
        props.recipes.slice(minValue, maxValue).map(val => (
          <Card 
            title={val.label}
            extra={<a href="#">More</a>}
            style={{ width: 300 }}
          >
            <p>{val.id}</p>
          </Card> 
        ))}
      </Grid>
    </Grid>
    <Pagination
        defaultCurrent={1}
        defaultPageSize={5}
        onChange={handleChange}
        total={props.recipes.length}
      />
    </Grid>
  );
}

export default ShowContent;
