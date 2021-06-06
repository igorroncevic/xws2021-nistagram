import "../../style/menu.css";

import { ReactComponent as Home } from "../../images/home.svg";
import { ReactComponent as Inbox } from "../../images/inbox.svg";
import { ReactComponent as Explore } from "../../images/explore.svg";
import { ReactComponent as Notifications } from "../../images/notifications.svg";
import image from "../../images/profile.jpg";
import ProfileIcon from "./ProfileIcon";
import {NavLink} from "react-router-dom";
import {useEffect, useState} from "react";

function Menu(props) {
    console.log("menuu")
    const{user}=props;


    console.log(user)
    function click(){
        console.log("JEEEJ")
    }
    return (
        <div className="menu">
            <Home className="icon"/>
            <button onClick={click}><Inbox className="icon" /></button>
            <Explore className="icon" />
            <Notifications className="icon" />
            <NavLink  to={{ pathname: "/profile", state: { user:user, follow:false}  }}  >
                <ProfileIcon iconSize="small" image='https://i.pravatar.cc/150?img=1' />
            </NavLink>
        </div>
    );
}

export default Menu;