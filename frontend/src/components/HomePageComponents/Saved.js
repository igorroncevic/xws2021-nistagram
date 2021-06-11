import React, {useEffect, useState} from 'react';
import Navigation from "../HomePage/Navigation";


function Saved(props) {
    const [user,setUser] =useState({});

    useEffect(() => {
        if(!props.location.state) window.location.replace("http://localhost:3000/unauthorized");
        setUser(props.location.state.user);
    },[]);

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation user={user}/>

            <h1>Saved</h1>
        </div>

    );

}export default Saved;