import React, {useEffect, useState} from 'react';
import {useDispatch, useSelector} from "react-redux";

function Archived() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        if(store.user.role === 'Admin' || store.user.role === "") window.location.replace("http://localhost:3000/unauthorized");
    }, []);

    return (
        <div style={{marginTop:'55%'}}>
        </div>
    );
}

export default Archived;