import  React from "react";
import {IndexPage} from './pages/IndexPage.js'
import {ForgotPasswordPage} from './pages/forgotPass/ForgotPasswordPage.js'
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import RegistrationPage from "./pages/RegistrationPage";
import PostsAndStories from "./components/HomePageComponents/PostsAndStories";
import Chats from "./components/HomePageComponents/Chats";
import Saved from "./components/HomePageComponents/Saved";
import HomePage from "./components/HomePageComponents/HomePage";
import Search from "./pages/Search";
import Profile from "./components/ProfileComponent/Profile";
import NewPost from "./components/PostComponent/NewPost";
import Home from "./components/HomePage/Home";
import ChangePassword from "./components/ProfileComponent/ChangePassword";



function App () {
    return (
        <div className="App">
            <Router>
                <Route path='/' exact={true} component={IndexPage}/>
                <Route path='/forgotten' exact={true} component={ForgotPasswordPage}/>
                <Route path='/registration' exact={true} component={RegistrationPage}/>
                <Route path='/posts' exact component={PostsAndStories} />
                <Route path='/newpost' exact component={NewPost} />
                <Route path='/profile' exact component={Profile}/>
                <Route path='/chats' exact component={Chats} />
                <Route path='/saved' exact component={Saved} />
                <Route path='/search' exact={true} component={Search}/>
                <Route path='/pass' exact={true} component={ChangePassword}/>



                <Route path='/aj' exact  component={Home}/>
             </Router>
        </div>
    );
}
export default App