import Navigation from "../HomePage/Navigation";
import {Button, Table} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import userService from "../../services/user.service";
import toastService from "../../services/toast.service";
import {useSelector} from "react-redux";

const CampaignRequests = () => {
    const[requests,setRequests]=useState([{campaign:'',influencer:'idinfluenc',status:'Pending', postAt:'danas' }])
    const store = useSelector(state => state);

    useEffect(() => {
       // getAllRequests()
    }, []);

    async function getAllRequests() {
        const response = await userService.getCampaignRequests({
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Prosao, pazi sta treba da setuje!");
            setRequests(response.data)
        } else {
            toastService.show("error", "Something went wrong.Please try again!");
        }
    }

    return(
        <div>
            <Navigation/>
            <div style={{marginTop:'5%',marginLeft:'10%', marginRight:'10%', marginBottom:'20%'}}>
                <h3 style={{borderBottom:'1px solid black'}}>Campaign requests</h3>
                <Table striped bordered hover style={{marginTop:'3%'}}>
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Campaign</th>
                        <th>Influencer</th>
                        <th>Status</th>
                        <th>Post at</th>
                    </tr>
                    </thead>
                    <tbody>
                    {requests.map((request,index) => {
                        return (
                            <tr>
                                <td>{index+1}</td>
                                <td>cekaj jos</td>
                                <td>{request.influencer}</td>
                                <td>{ request.status==="Pending" &&<p  style={{color:'blueviolet' }}>{request.status}</p> }
                                    { request.status==="Accepted" &&<p  style={{color:'yellowgreen' }}>{request.status}</p> }
                                    { request.status==="Rejected" &&<p  style={{color:'red' }}>{request.status}</p> }
                                </td>
                                <td>{request.postAt}</td>
                            </tr>
                        )
                    })}
                    </tbody>
                </Table>
            </div>
        </div>
    );
}
export default CampaignRequests;
