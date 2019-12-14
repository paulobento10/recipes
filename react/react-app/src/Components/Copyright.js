import React, { useState } from "react";
import Typography from '@material-ui/core/Typography';

export default function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
        Universidade Fernando Pessoa - PAWB{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}