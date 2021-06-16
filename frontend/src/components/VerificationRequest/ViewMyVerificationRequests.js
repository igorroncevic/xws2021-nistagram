import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button} from "react-bootstrap";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";



function ViewMyVerificationRequests() {

    const dispatch = useDispatch()
    const store = useSelector(state => state);
    
    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <h1>My verification requests</h1>
        </div>
    );
}

export default ViewMyVerificationRequests;