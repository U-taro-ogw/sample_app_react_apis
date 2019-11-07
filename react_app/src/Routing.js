import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Todo from './Todo';
import Menu from './Menu';

class Routing extends React.Component {
  render() {
    return (
      <BrowserRouter>
        <Switch>
          <Route exact path='/' component={Menu} />
          <Route exact path='/todos' component={Todo} />
        </Switch>
      </BrowserRouter>
    )
  }
}

export default Routing;