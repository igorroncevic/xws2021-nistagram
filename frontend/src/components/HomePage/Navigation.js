import "../../style/navigation.css";
import Menu from "./Menu";
import {useEffect, useState} from "react";

function Navigation(props) {
    const [user, setUser] = useState({...props.user});
    useEffect(() => {
        setUser(props.user);
    }, [props.user])

    return (
        <div className="navigation">
            <div className="container">
                <font face = "Comic Sans MS" size = "5" style={{marginRight:'5em'}}>Ništagram</font>
                    <input type="text" placeholder="Search.."style={{marginRight:'25em'}}/>
                <Menu user={user}/>
            </div>
        </div>
    );
}

export default Navigation;