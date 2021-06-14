import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import { Provider } from 'react-redux'
import { PersistGate } from 'redux-persist/integration/react'
import 'bootstrap/dist/css/bootstrap.css';
import "react-bootstrap";

import configureStore from './store'

const { store, persistor } = configureStore();
 
// Add ToastContainer
ReactDOM.render(
    <Provider store={store}>
        <PersistGate loading={null} persistor={persistor}>
            <App/> 
        </PersistGate>
    </Provider>, 
document.getElementById('root'))
