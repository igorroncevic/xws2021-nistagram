import React, {useEffect, useState} from 'react';
import Navigation from "../HomePage/Navigation";


function Notifications(props) {
    console.log("chats")
    console.log(props)
    const [user,setUser] =useState({});

    useEffect(() => {
        if(!props.location.state) window.location.replace("http://localhost:3000/unauthorized");
        setUser(props.location.state.user);
    },[]);

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation user={user}/>

            <h1>Notifikacije</h1>
        </div>

    );

}export default Notifications;