import * as React from "react";
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import Home from './components/Home/home'
import Signin from './components/Signin/signin'
import Signup from './components/Signin/signup'

const URLs = {
    home: '/',
    signin: '/signin',
    signup: '/signup',
}

export const Routes = (
    <Router>
        <Switch>
            <Route exact path={URLs.home} component={Home} key={0}/>
            <Route exact path={URLs.signin} component={Signin} key={1}/>
            <Route exact path={URLs.signup} component={Signup} key={2}/>
        </Switch>
    </Router>
);
