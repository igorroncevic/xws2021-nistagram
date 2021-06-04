import React, {useState} from 'react';
import {Container, Row, Col, Table, FormControl} from "react-bootstrap";
import HomePage from "./HomePage";
import EditProfile from "./EditProfile";


function ProfilePage() {

    return (
        <div>
        <HomePage/>
            <p style={{display:'inline-block', float:'left'}}>PROFILE PHOTO  ></p>

            <div style={{ display:'inline-block', float:'right'}}>
                <EditProfile/>

            </div>
        </div>
    )

}export default ProfilePage;