import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Home from './page/home/index.jsx';
import Login from './page/login/index.jsx';

import './app.scss';

class App extends React.Component {
  render() {
    return(
      <Router>
        <Switch>
          <Route exact component={Home} path='/admin' />
          <Route exact component={Login} path='/admin/login' />
          {/* <Route exact component={AnimalIndex} path='/animals' />
          <Route exact component={AdoptionTips} path='/adoption-tips' />
          <Route exact component={About} path='/about' />

          <Route exact component={AnimalShow} path='/animals/:id' /> */}
        </Switch>
      </Router>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
