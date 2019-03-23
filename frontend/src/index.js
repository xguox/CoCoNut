import React from 'react';
import ReactDOM from 'react-dom';
// import './index.css';

// import registerServiceWorker from './registerServiceWorker';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Home from './page/home/index.jsx';

class App extends React.Component {
  render() {
    return(
      <Router>
        <Switch>
          <Route exact component={Home} path='/' />
          {/* <Route exact component={AnimalIndex} path='/animals' />
          <Route exact component={AdoptionTips} path='/adoption-tips' />
          <Route exact component={About} path='/about' />

          <Route exact component={AnimalShow} path='/animals/:id' /> */}
        </Switch>
      </Router>

      // <div>123</div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
// registerServiceWorker();
