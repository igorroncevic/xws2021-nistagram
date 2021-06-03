import HomePage from "../components/FrontPageComponents/HomePage";
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import PostsAndStories from "../components/FrontPageComponents/PostsAndStories";
import ProfilePage from "../components/FrontPageComponents/ProfilePage";
import Chats from "../components/FrontPageComponents/Chats";
import {useEffect} from "react";

function FrontPage() {
    useEffect(() => {
        document.body.style.backgroundColor = "  #ffeecc"
    });
    return(
        <div >
        <Router>
            <HomePage />
            <Switch>
                <Route path='/' exact component={PostsAndStories} />
                <Route path='/profile' exact component={ProfilePage} />
                <Route path='/chats' exact component={Chats} />

            </Switch>
        </Router>
        </div>
    );
}export default  FrontPage;