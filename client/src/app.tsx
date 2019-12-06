import * as React from "react";
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './containers/home'
import Signin from './containers/signin'
import Signup from './containers/signup'

const URLs = {
    home: '/',
    signin: '/signin',
    signup: '/signup'
}

export default class App extends React.Component {
    render() {
        return (
            <Router>
                <Switch>
                    <Route exact path={URLs.home} component={Home} key={1}/>    
                    <Route exact path={URLs.signin} component={Signin} key={2}/>
                    <Route exact path={URLs.signup} component={Signup} key={3}/>                                    
                </Switch>
            </Router>
        )
    }
}
