import React, { Component } from 'react'
import { Helmet } from 'react-helmet'
import Header from '../header/header';
import BoardAdder from './boardAdder';

class Home extends Component {
    render() {
        return (
            <>
                <Helmet>
                    <title>Home | Task Management</title>
                </Helmet>

                <Header />
                <BoardAdder />
            </>           
        
        )
    }
}

export default Home
