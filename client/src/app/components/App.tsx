import * as React from "react";
import { BrowserRouter as Router, Switch, Route, Redirect, withRouter } from 'react-router-dom';
import { connect } from "react-redux";

import "./App.scss";
import Home from './Home/home';
import Signin from './Signin/signin';
import Signup from './Signin/signup';
import LandingPage from './LandingPage/LandingPage';

import { AppState } from '../states/AppState';
import { UserState, InitialUserState } from '../states/UserState'

const URLs = {
    home: '/',
    signin: '/signin',
    signup: '/signup',
}

const App = (isUser: boolean, isGuest: boolean ) => {
    // Serve different pages depending on if user is logged in or not
    if (isUser || isGuest) {
        return (
            <Router>
                <Switch>
                    <Route exact path={URLs.home} component={Home} key={0}/>
                    <Route exact path={URLs.signin} component={Signin} key={1}/>
                    <Route exact path={URLs.signup} component={Signup} key={2}/>
                </Switch>
            </Router>
        );
    }
  
    // If not logged in, always redirect to landing page
    return (
        <Switch>
            <Route exact path="/" component={LandingPage} />
            <Redirect to="/" />
        </Switch>
    );
};

const mapStateToProps = (state: AppState) => ({ user: state.user, isGuest: state.isGuest });

// Use withRouter to prevent strange glitch where other components
// lower down in the component tree wouldn't update from URL changes
export default withRouter(connect(mapStateToProps)(App));
