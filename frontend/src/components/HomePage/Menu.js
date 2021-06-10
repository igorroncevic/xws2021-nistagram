import "../../style/menu.css";
import { ReactComponent as Home } from "../../images/home.svg";
import { ReactComponent as Inbox } from "../../images/inbox.svg";
import { ReactComponent as Notifications } from "../../images/notifications.svg";
import { ReactComponent as Bookmark } from "../../images/bookmark.svg";
import ProfileIcon from "../ProfileComponent/ProfileIcon";
import {NavLink} from "react-router-dom";

function Menu(props) {
    const{user}=props;

    return (
        <div className="menu">
            <NavLink  to={{ pathname: "/home", state: { user:user}  }}  >
                <Home className="icon"/>
            </NavLink>
            <NavLink  to={{ pathname: "/chats", state: { user:user}  }}  >
                <Inbox className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/notifications", state: { user:user}  }}  >
                    <Notifications className="icon" />
            </NavLink>
            <NavLink  to={{ pathname: "/saved", state: { user:user}  }}  >
                <Bookmark className="icon" />
            </NavLink>

            <NavLink  to={{ pathname: "/profile", state: { user:user, follow:false}  }}  >
                <ProfileIcon iconSize="medium" image='https://i.pravatar.cc/150?img=1' />
            </NavLink>
            <a href='/' style={{marginLeft:'10px'}}>Log out</a>

        </div>
    );
}

export default Menu;