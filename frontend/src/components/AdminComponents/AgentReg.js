import React, { useEffect, useState } from 'react';
import { Button, Modal, ListGroup } from "react-bootstrap";
import { useSelector } from 'react-redux';
import { ReactComponent as Check } from './../../images/icons/check.svg'
import collectionsService from './../../services/collections.service'
import favoritesService from './../../services/favorites.service'
import toastService from './../../services/toast.service'
import './../../style/CollectionsModal.css'
import RegistrationPage from "../../pages/RegistrationPage";
import Navigation from "../HomePage/Navigation";

const AgentReg = (props) => {


    return (
        <div>
            <Navigation/>
            <div style={{marginTop:'5%',marginLeft:'20%', marginRight:'20%', marginBottom:'20%'}}>
                <h3 style={{borderBottom:'1px solid black'}}>Agent Registration</h3>
                <RegistrationPage role={'Agent'}/>

        </div>
        </div>

    )
}

export default AgentReg;