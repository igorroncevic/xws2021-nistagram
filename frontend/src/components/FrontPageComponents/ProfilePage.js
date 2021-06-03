import React, {useState} from 'react';
import {Container, Row, Col, Table, FormControl} from "react-bootstrap";
import HomePage from "./HomePage";
import UsersInfo from "./UsersInfo";
import Card from "../CardComponent/Card";


function ProfilePage() {

    return (
        <div>
        <HomePage/>
            <p style={{display:'inline-block', float:'left'}}>PROFILE PHOTO  ></p>

            <div style={{ display:'inline-block', float:'right'}}>
                <UsersInfo/>
            </div>
        </div>
    )

}export default ProfilePage;