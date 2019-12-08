import React, { /*Component,*/ useState } from 'react';
import './App.css';
import {BrowserRouter as Router, Switch,Route} from "react-router-dom";
import { AuthContext } from "./context/auth";
import PrivateRoute from './PrivateRoute';
import Show from './Components/Show';
import Insert from './Components/Insert';
import SignIn from './Components/SignIn';
import SignUp from './Components/SignUp';

function App(props) {
  const [authTokens, setAuthTokens] = useState();
  
  const setTokens = (data) => {
    localStorage.setItem("tokens", JSON.stringify(data));
    setAuthTokens(data);
  }

  return (
    <AuthContext.Provider value={{ authTokens, setAuthTokens: setTokens }}>
      <Router>
          <div className="App">
            <Switch>
              <Route exact path="/" component={Show}/>
              <Route exact path="/signin" component={SignIn}/>
              <Route exact path="/signup" component={SignUp}/>
              <PrivateRoute exact path="/insert" component={Insert}/>
            </Switch>
          </div>
          </Router>
    </AuthContext.Provider>
  );
}

export default App;