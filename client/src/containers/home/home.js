import React, { Component } from 'react'
import { Helmet } from 'react-helmet'

import Signin from '../login/signin'
import Signup from '../login/signup'
// import Header from '../header/header';
// import BoardAdder from './boardAdder';

class Home extends Component {
    render() {
        return (
            <>
                <Helmet>
                    {/* <title>Home | Task Management</title> */}
                    <title>Login | Task Management</title>
                </Helmet>

                {/* <Header />
                <BoardAdder /> */}
                {/* <Signin /> */}
                <Signup />
            </>           
        
        )
    }
}

export default Home
