import React from "react";
import { useAuth } from "../context/auth";
import Button from '@material-ui/core/Button';

function Insert(props) {
  const { setAuthTokens } = useAuth();

  function logOut() {
    setAuthTokens();
  }

  return (
    <div>
      <div>Insert Page - Teste</div>
      <Button onClick={logOut}>Log out</Button>
    </div>
  );
}

export default Insert;