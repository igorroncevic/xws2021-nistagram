import "../../style/menu.css";
import { ReactComponent as Home } from "../../images/home.svg";
import { ReactComponent as Inbox } from "../../images/inbox.svg";
import { ReactComponent as Explore } from "../../images/explore.svg";
import { ReactComponent as Notifications } from "../../images/notifications.svg";
import ProfileIcon from "./ProfileIcon";
import {NavLink} from "react-router-dom";

function Menu(props) {
    console.log("menuu")
    const{user}=props;

    return (
        <div className="menu">
            <Home className="icon"/>
            <Inbox className="icon" />
            <Explore className="icon" />
            <NavLink  to={{ pathname: "/", state: { user:user}  }}  >
                    <Notifications className="icon" />
            </NavLink>

            <NavLink  to={{ pathname: "/profile", state: { user:user, follow:false}  }}  >
                <ProfileIcon iconSize="small" image='https://i.pravatar.cc/150?img=1' />
            </NavLink>
        </div>
    );
}

export default Menu;