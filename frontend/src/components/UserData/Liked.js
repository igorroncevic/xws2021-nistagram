import React, {useEffect, useState} from 'react';
import ProfileInfo from "./ProfileInfo";
import {useDispatch, useSelector} from "react-redux";

function Liked() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        if(store.user.role === 'Admin' || store.user.role === "") window.location.replace("http://localhost:3000/unauthorized");
    }, []);

    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>

            </div>
        </div>
    );
}

export default Liked;