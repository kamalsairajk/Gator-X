import React from "react";
import { BrowserRouter as Router, Routes, Route, 
Link } from "react-router-dom";
import Homepage from './components/Homepage'
import Review from './components/Review'

function App(){
  axios.defaults.baseURL = "http://localhost:9000/"
  console.log(localStorage.getItem("user"));

  return (
  <Router>
    <nav>
      <Link to = "/"> Homepage </Link>
      <Link to = "/review"> Review </Link>
    </nav>
      <Routes>
      <Route exact path="/" render={() => <HomePage />} />
      <Route exact path="/register" render={() => <RegisterComponent />} />
      <Route exact path="/login" render={() => <LoginComponent />} />
      </Routes>
  </Router>
  );
}

export default App;
