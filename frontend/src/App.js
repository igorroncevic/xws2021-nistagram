import  React from "react";
import {LoginPage} from './pages/LoginPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

function App () {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact={true} component={LoginPage}/>
             </Router>
        </div>
    );
}
export default App