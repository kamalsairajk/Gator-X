import React from "react";
import { BrowserRouter as Router, Routes, Route, 
Link } from "react-router-dom";
import Layout from '/frontend/Page-elements/Layout';
import Homepage from '/frontend/components/Homepage'
import Review from '/frontend/components/Review'

const App = () => {
  <Router>
      <Routes>
          <Route path = '/' component = {Homepage}/>
      </Routes>
  </Router>
};

export default App;
