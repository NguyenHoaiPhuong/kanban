import React, { Component } from 'react'

import { BrowserRouter as Router, Route } from "react-router-dom";

import Home from './home/home';
import Signin from './login/signin'
import Signup from './login/signup'

class App extends Component {
    render() {
        return (
            <Router>
                <div>
                    <Route exact path="/" component={Home} />
                    <Route exact path="/signin" component={Signin} />
                    <Route exact path="/signup" component={Signup} />
                </div>
            </Router>
        )
    }
}

export default App
