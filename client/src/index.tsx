import * as React from "react";
import * as ReactDOM from "react-dom";
import { BrowserRouter as Router} from 'react-router-dom';
import { Provider } from 'react-redux';

import './index.css';
import App from './app/components/App';
import { store } from './app/store/store';

// ReactDOM.render(<App />, document.getElementById('root'));

store.subscribe(() => {
    console.log(store.getState())
})

ReactDOM.hydrate(
    <Provider store={store}>
      <Router>
        <App />
      </Router>
    </Provider>,
    document.getElementById("root")
  );
  
