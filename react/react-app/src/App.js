import React, { useState } from 'react';
import './App.css';
import {BrowserRouter as Router, Switch,Route} from "react-router-dom";
import { AuthContext } from "./context/auth";
import PrivateRoute from './PrivateRoute';
import Show from './Pages/Show';
import Insert from './Pages/Insert';
import SignIn from './Pages/SignIn';
import SignUp from './Pages/SignUp';
import Recipe from './Pages/Recipe';

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
              <Route exact path="/show" component={Show}/>
              <Route exact path="/signin" component={SignIn}/>
              <Route exact path="/signup" component={SignUp}/>
              <Route exact path="/insert" component={Insert}/> {/*<PrivateRoute</div> exact path="/insert" component={Insert}/>*/}
              <Route exact path="/show/recipe/:id" component={Recipe}/>
            </Switch>
          </div>
        </Router>
    </AuthContext.Provider>
  );
}

export default App;