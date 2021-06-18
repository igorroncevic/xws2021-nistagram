import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button, Table} from "react-bootstrap";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";



function ViewPendingVerificationRequests() {
    const [verificationRequests, setVerificationRequests] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        if(store.user.role !== 'Admin') window.location.replace("http://localhost:3000/unauthorized");
        getPendingVerificationRequests();
    }, [])

    function getPendingVerificationRequests() {
        axios.get("http://localhost:8080/api/users/api/users/get-pending-verification-requests", {
            headers : userService.setupHeaders(store.user.jwt)
        }).then(res => setVerificationRequests(res.data.verificationRequests))
    }

    function changeVerificationRequestStatus(verificationRequest, status) {
        axios.post("http://localhost:8080/api/users/api/users/change-verification-request-status", {
            id : verificationRequest.id,
            status : status
        },{
            headers : userService.setupHeaders(store.user.jwt)
        }).then(res => {
            alert("Verification request status changed successfully!")
            getPendingVerificationRequests();
        }).catch(err => {
            alert("Error while attempting to change verification request status!")
        })
    }

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <h1>Pending verification requests</h1>

            <Table striped bordered hover>
                <thead>
                <tr>
                    <th>#</th>
                    <th>First name</th>
                    <th>Last name</th>
                    <th>Category</th>
                    <th>Status</th>
                    <th>Created at</th>
                    <th>Document photo</th>
                    <th>Profile</th>
                    <th></th>
                </tr>
                </thead>
                <tbody>
                {verificationRequests.map((verificationRequest,index) => {
                    return (
                        <tr>
                            <td>{index+1}</td>
                            <td>{verificationRequest.firstName}</td>
                            <td>{verificationRequest.lastName}</td>
                            <td>{verificationRequest.category}</td>
                            <td>{verificationRequest.status}</td>
                            <td>{verificationRequest.createdAt}</td>
                            <td>
                                <img  src={verificationRequest.documentPhoto} alt="document photo" style={{width:'200px', height: '200px'}}/>
                            </td>
                            <td>Profile</td>
                            <td>
                                <Button variant={"success"} onClick={() => changeVerificationRequestStatus(verificationRequest,'Accepted')}>Accept</Button>
                                <span  style={{marginLeft: '5%'}}/>
                                <Button variant={"danger"} onClick={() => changeVerificationRequestStatus(verificationRequest,'Refused')}>Refuse</Button>

                            </td>
                        </tr>
                    )
                })}
                </tbody>
            </Table>
        </div>
    );
}

export default ViewPendingVerificationRequests;