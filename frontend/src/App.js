import  React from "react";
import {LoginPage} from './pages/LoginPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';

function App () {
    return (
        <div className="App">
            <Router>
                <Route path='/login' exact={true} component={LoginPage}/>

             </Router>
        </div>
    );
}
export default App