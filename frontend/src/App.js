import  React from "react";
import IndexPage from './pages/IndexPage.js'
import {ForgotPasswordPage} from './components/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPage from "./pages/RegistrationPage";
import Chats from "./components/HomePageComponents/Chats";
import Saved from "./components/HomePageComponents/Saved";
import Search from "./pages/Search";
import Profile from "./components/ProfileComponent/Profile";
import NewPost from "./components/HomePageComponents/NewPost";
import Home from "./components/HomePage/Home";
import ChangePassword from "./components/ProfileComponent/ChangePassword";
import Notifications from "./components/HomePageComponents/Notifications";
import UnauthorizedPage from "./helpers/UnauthorizedPage";

const App = () => {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact={true} component={IndexPage}/>
                <Route path='/unauthorized' exact={true} component={UnauthorizedPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/home' exact  component={Home}/>
                <Route path='/search' exact={true} component={Search}/>
                <Route path='/profile' exact component={Profile}/>

                <Route path='/newpost' exact component={NewPost} />
                <Route path='/chats' exact component={Chats} />
                <Route path='/saved' exact component={Saved} />
                <Route path='/notifications' exact component={Notifications} />
            </Router>
        </div>
    );
}

export default App