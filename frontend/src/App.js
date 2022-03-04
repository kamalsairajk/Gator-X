import React from "react";
import axios from 'axios'
import { BrowserRouter as Router, Routes, Route, 
Link } from "react-router-dom";
import Homepage from './components/Homepage'
import Review from './components/Review'
import {Navbar, Nav, Button} from 'react-bootstrap'

function App(){
  axios.defaults.baseURL = "http://localhost:8080/"
  console.log(localStorage.getItem("user"));

  return (
  <div className="App">
    <Router>
      <nav>
        <Link to = "/"> Homepage </Link>
        <Link to = "/login"> Login </Link>
        <Link to = "/register"> RegisterComponent </Link>
        <Link to = "/review"> Review </Link>
      </nav>
        <Routes>
        <Route exact path="/" render={() => <HomePage />} />
        <Route exact path="/register" render={() => <RegisterComponent />} />
        <Route exact path="/login" render={() => <Login/>} />
        </Routes>
    </Router>
  </div>
  );
}

export default App;
