import React, { Component } from 'react'

import { BrowserRouter as Router, Route } from "react-router-dom";

import PrivateRoute from './privateRoute/privateRoute'
import Home from './home/home';
import Signin from './login/signin'
import Signup from './login/signup'

class App extends Component {
    render() {
        return (
            <Router>
                <PrivateRoute exact path="/" component={Home} />
                <Route path="/signin" component={Signin} />
                <Route path="/signup" component={Signup} />                        
            </Router>
        )
    }
}

export default App
