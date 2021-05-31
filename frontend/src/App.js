import  React from "react";
import {IndexPage} from './pages/IndexPage.js'
import {ForgotPasswordPage} from './pages/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import {ResetPasswordCode} from "./pages/forgotPass/ResetPasswordCode";
import {PasswordChange} from "./pages/forgotPass/PasswordChange";

function App () {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact={true} component={IndexPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/reset' exact={true} component={PasswordChange}/>
             </Router>
        </div>
    );
}
export default App