import  React from "react";
import IndexPage from './pages/IndexPage.js'
import {ForgotPasswordPage} from './components/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPage from "./pages/RegistrationPage";
import Saved from "./components/HomePageComponents/Saved";
import NewProduct from "./components/Product/NewProduct";
import Home from "./components/HomePage/Home";
import UnauthorizedPage from "./helpers/UnauthorizedPage";

import ProfileInfo from "./components/UserData/ProfileInfo";
import EditProfile from "./components/UserData/EditProfile";
import ChangePassword from "./components/UserData/ChangePassword";

import EditProfileImage from "./components/UserData/EditProfileImage";
import Notifications from "./components/Notifications/Notifications";

import AuthenticatedRoute from './AuthenticatedRoute';
import AdminRoute from './AdminRoute';

const App = () => {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact  component={Home}/>
                <Route path='/login' exact={true} component={IndexPage}/>
                <Route path='/unauthorized' exact={true} component={UnauthorizedPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/info' exact component={ProfileInfo}/>
                <Route path='/newproduct' exact component={NewProduct}/>

                <AuthenticatedRoute path='/saved' exact component={Saved} isAdminProhibited={true} />
                <AuthenticatedRoute path='/notifications' exact component={Notifications} isAdminProhibited={true} />
                <AuthenticatedRoute path='/edit_profile' exact component={EditProfile} />
                <AuthenticatedRoute path='/password' exact component={ChangePassword} />
                <AuthenticatedRoute path='/edit_photo' exact component={EditProfileImage} />
            </Router>
        </div>
    );
}

export default App