import "../../style/menu.css";
import { ReactComponent as Home } from "../../images/icons/home.svg";
import { ReactComponent as Inbox } from "../../images/icons/inbox.svg";
import { ReactComponent as Notifications } from "../../images/icons/notifications.svg";
import { ReactComponent as Bookmark } from "../../images/icons/bookmark.svg";
import { ReactComponent as StoryArchive } from "../../images/icons/story-archive.svg";
import { ReactComponent as Plus } from "../../images/icons/plus.svg";
import { ReactComponent as Explore } from "../../images/icons/more.svg";
import { ReactComponent as VerificationSymbol } from "../../images/icons/verification-symbol.svg";

import ProfileIcon from "../ProfileComponent/ProfileIcon";
import { NavLink, useHistory } from "react-router-dom";
import React, {useEffect, useState} from "react";
import { useSelector} from "react-redux";
import { Dropdown, Button } from "react-bootstrap";

function Menu() {
    const[username, setUsername]=useState('')
    const store = useSelector(state => state);
    const history = useHistory()


    useEffect(() => {
        setUsername(store.user.username ? store.user.username : "")
    }, []);

    function verificationRedirect(text) {
        switch(text) {
            case 'submit-verification-request' :
                history.push({ pathname: '/submit-verification-request' });
                break;
            case 'view-my-verification-request' :
                history.push({ pathname: '/view-my-verification-request' })
                break;
            case 'view-pending-verification-request' :
                history.push({ pathname: '/view-pending-verification-request' })
                break;
            case 'view-all-verification-request' :
                history.push({ pathname: '/view-all-verification-request' })
                break;
            default: return;
        }
    }

    return (
        <div className="menu">
            <NavLink  to={{ pathname: "/home" }} >
                <Home className="icon"/>
            </NavLink>
            <NavLink  to={{ pathname: "/chats" }} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}} >
                <Inbox className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/notifications" }} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}} >
                <Notifications className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/saved" }} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}} >
                <Bookmark className="icon" /> 
            </NavLink>
            <NavLink  to={{ pathname: "/story-archive" }} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}} >
                <StoryArchive className="icon" /> 
            </NavLink>
            <NavLink  to={{ pathname: "/newpost" }} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}} >
                <Plus className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/info" }} >
                <Explore className="icon" />
            </NavLink>
            <Dropdown>
                <Dropdown.Toggle variant="link" id="dropdown-basic">
                    <VerificationSymbol className="icon" />
                </Dropdown.Toggle>

                <Dropdown.Menu>

                    <Dropdown.Item onClick={() => verificationRedirect('submit-verification-request')} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>Submit verification request

                    </Dropdown.Item>
                    <Dropdown.Item onClick={() => verificationRedirect('view-my-verification-request')} style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}}>View my verification requests</Dropdown.Item>
                    <Dropdown.Item onClick={() => verificationRedirect('view-pending-verification-request')} style={store.user.role === 'Admin' ? {display : 'block'} : {display: 'none'}}>View pending verification requests</Dropdown.Item>
                    <Dropdown.Item onClick={() => verificationRedirect('view-all-verification-request')} style={store.user.role === 'Admin' ? {display : 'block'} : {display: 'none'}}>View all verification requests</Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>

            <NavLink  to={"/profile/"+username } style={store.user.role !== 'Admin' ? {display : 'block'} : {display: 'none'}} >
                <ProfileIcon iconSize="medium" image={store.user.photo ? store.user.photo : 'https://i.pravatar.cc/150?img=1'} />
            </NavLink>

            <Button variant="outline-danger" href='/' style={{width: "220px", display: "block"}}>Logout</Button>

        </div>
    );
}

export default Menu;