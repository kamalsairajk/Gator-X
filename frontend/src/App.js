import React from "react";
import { BrowserRouter as Router, Routes, Route, 
Link } from "react-router-dom";
import Homepage from './components/Homepage'
import Review from './components/Review'

const App = () => {
  <Router>
    <nav>
      <Link to = "/"> Homepage </Link>
      <Link to = "/review"> Review </Link>
    </nav>
      <Routes>
          <Route path = '/' element = {Homepage}/>
          <Route path = '/Review' element = {<Review />}/>
      </Routes>
  </Router>
};

export default App;
