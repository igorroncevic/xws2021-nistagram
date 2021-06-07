import "../../style/navigation.css";
import Menu from "./Menu";
import React, {useEffect, useState} from "react";
import {Button, Modal} from "react-bootstrap";
import RegistrationPage from "../../pages/RegistrationPage";
import Search from "../../pages/Search";
import {useHistory} from "react-router-dom";


function Navigation(props) {
    const [user, setUser] = useState({...props.user});
    const history = useHistory()

    useEffect(() => {
        setUser(props.user);
    }, [props.user])

    function redirectToSearch(){
        history.push({
            pathname: '/search',
        })
    }

    return (
        <div className="navigation">
            <div className="container">
                <font face = "Comic Sans MS" size = "5" style={{marginRight:'5em'}}>Ništagram</font>
                <Button variant="outline-dark" style={{marginRight:'25em'}} onClick={redirectToSearch}>Search...</Button>
                {/*<input type="text" placeholder="Search.." style={{marginRight:'25em'}} onClick={props.getUser}/>*/}
                <Menu user={user} getUser={props.getUser}/>
            </div>
        </div>
    );
}

export default Navigation;