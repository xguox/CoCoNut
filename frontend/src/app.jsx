import React, { Component } from 'react';
import ReactDOM from 'react-dom';

export default class App extends Component{
    render(){
        return (
            <div id="page-wrapper">hello COCO</div>
        )
    }
}


ReactDOM.render(
    <App />,
    document.getElementById('app')
);