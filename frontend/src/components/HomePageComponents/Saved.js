import React, {useEffect} from 'react';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";


function Saved() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        if(store.user.role === 'Admin' || store.user.role === "") window.location.replace("http://localhost:3000/unauthorized");
    }, []);

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <h1>Saved</h1>
        </div>
    );
}

export default Saved;