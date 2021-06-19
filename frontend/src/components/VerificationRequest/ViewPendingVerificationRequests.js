import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button, Table} from "react-bootstrap";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";
import verificationRequestService from "../../services/verificationRequest.service";
import toastService from "../../services/toast.service";



function ViewPendingVerificationRequests() {
    const [verificationRequests, setVerificationRequests] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        if(store.user.role !== 'Admin') window.location.replace("http://localhost:3000/unauthorized");
        getPendingVerificationRequests();
    }, [])

    async function getPendingVerificationRequests() {
        const response = await verificationRequestService.getPendingVerificationRequests({jwt : store.user.jwt})
        setVerificationRequests(response.data.verificationRequests);
    }

    async function changeVerificationRequestStatus(verificationRequest, status) {
        const response = await verificationRequestService.changeVerificationRequestStatus({
            id : verificationRequest.id,
            status : status,
            jwt : store.user.jwt
        });
        if (response.status === 200) {
            toastService.show("success", "Verification request status changed successfully")
            getPendingVerificationRequests();
        }
        else
            toastService.show("error", "Something went wrong. Try again")


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