import  React from "react";
import {IndexPage} from './pages/IndexPage.js'
import {ForgotPasswordPage} from './pages/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPageHooks from "./pages/RegistrationPageHooks";


function App () {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact={true} component={IndexPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration-hooks' exact={true} component={RegistrationPageHooks}/>
             </Router>
        </div>
    );
}
export default App