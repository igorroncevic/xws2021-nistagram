import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import 'bootstrap/dist/css/bootstrap.css';
import "react-bootstrap";
import { Provider } from 'react-redux'

import { store } from './store'
 
// Add ToastContainer
ReactDOM.render(
    <Provider store={store}>
        <App/> 
    </Provider>
, document.getElementById('root'))
