import * as React from "react";
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import Home from './containers/home'
import Signin from './containers/signin'
import Signup from './containers/signup'

const URLs = {
    home: '/',
    signin: '/signin',
    signup: '/signup',
}

export const routes = (
    <Router>
        <Switch>
            <Route exact path={URLs.home} component={Home} key={0}/>
            <Route exact path={URLs.signin} component={Signin} key={1}/>
            <Route exact path={URLs.signup} component={Signup} key={2}/>
        </Switch>
    </Router>
);
