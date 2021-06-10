import React, {useState} from 'react';
import Navigation from "../HomePage/Navigation";


function Saved(props) {
    const [user,setUser] =useState(props.location.state.user);

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation user={user}/>

            <h1>Saved</h1>
        </div>

    );

}export default Saved;