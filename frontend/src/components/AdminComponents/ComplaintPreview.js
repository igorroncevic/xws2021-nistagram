import React, { useEffect, useState } from 'react';
import {Button, Modal, ListGroup, Table} from "react-bootstrap";
import { useSelector } from 'react-redux';
import { ReactComponent as Check } from './../../images/icons/check.svg'
import collectionsService from './../../services/collections.service'
import favoritesService from './../../services/favorites.service'
import toastService from './../../services/toast.service'
import './../../style/CollectionsModal.css'
import RegistrationPage from "../../pages/RegistrationPage";
import Navigation from "../HomePage/Navigation";
import ProfileInfo from "../UserData/ProfileInfo";
import PostPreviewGrid from "../Post/PostPreviewGrid";
import likeService from "../../services/like.service";
import complaintService from "../../services/complaint.service";
import userService from "../../services/user.service";
import {NavLink} from "react-router-dom";

const ComplaintPreview = (props) => {
    const [posts, setPosts] = useState([]);
    const [users,setUsers]=useState([]);
    const store = useSelector(state => state);

    useEffect(() => {
        getComplaints()
    }, []);

    function getComplaints() {
        console.log("EVE ME")
        complaintService.getAllContentComplaints({ jwt: store.user.jwt })
            .then(response => {
                if(response.status === 200)
                    setParams(response.data.contentComplaints)
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }
    async function  setParams(complaints){
        console.log("I tu sam")
        complaints.map((complaint, i) =>
             getUserById(complaint.userId)
        );
        setPosts(complaints)
    }


     async function getUserById(id) {
         console.log("I ovde")
         console.log(id)

        await userService.getUserById({id: id, jwt: store.user.jwt})
             .then(response => {
                 setUsers(users => [...users, response.data])
                 console.log(users)
             })
             .catch(err => {
                 console.log("BLA")
                 toastService.show("error", "Error")
             })

    }


    return (
        <div>
            <Navigation/>
            <div style={{marginTop:'5%',marginLeft:'20%', marginRight:'20%', marginBottom:'20%'}}>
                <h3 style={{borderBottom:'1px solid black'}}>Complaints</h3>
                <Table striped bordered hover>
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Category</th>
                        <th>Post/Story</th>
                        <th>User</th>
                        <th>Status</th>
                        <th></th>
                    </tr>
                    </thead>
                    <tbody>
                    {posts.map((post,index) => {
                        return (
                            <tr>
                                <td>{index+1}</td>
                                <td>{post.category}</td>
                                <td>bla</td>
                                <td><a href={`/profile/${users[index].username}`}> {users[index].username} </a></td>
                                <td>bla</td>
                                <td>{post.status}</td>

                            </tr>
                        )
                    })}
                    </tbody>
                </Table>
            </div>
        </div>


    )
}

export default ComplaintPreview;