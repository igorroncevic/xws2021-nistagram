import "../../style/menu.css";
import { ReactComponent as Home } from "../../images/home.svg";
import { ReactComponent as Inbox } from "../../images/inbox.svg";
import { ReactComponent as Notifications } from "../../images/notifications.svg";
import { ReactComponent as Bookmark } from "../../images/bookmark.svg";
import { ReactComponent as Plus } from "../../images/plus.svg";
import ProfileIcon from "../ProfileComponent/ProfileIcon";
import {NavLink, Route} from "react-router-dom";
import Profile from "../ProfileComponent/Profile";
import {useState} from "react";

function Menu() {
    const[username,setUsername]=useState("andja")
    return (
        <div className="menu">
            <NavLink  to={{ pathname: "/home" }}  >
                <Home className="icon"/>
            </NavLink>
            <NavLink  to={{ pathname: "/chats" }}  >
                <Inbox className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/notifications" }}  >
                <Notifications className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/saved" }}  >
                <Bookmark className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/newpost" }}  >
                <Plus className="icon" />
            </NavLink>

            <NavLink  to={"/profile/"+username }  >
                <ProfileIcon iconSize="medium" image='https://i.pravatar.cc/150?img=1' />
            </NavLink>

            <a href='/' style={{marginLeft:'10px'}}>Log out</a>

        </div>
    );
}

export default Menu;