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
import storyService from "../../services/story.service";
import postService from "../../services/post.service";

const ComplaintPreview = (props) => {
    const [posts, setPosts] = useState([]);
    const [users,setUsers]=useState([]);
    const store = useSelector(state => state);

    useEffect(() => {
        getComplaints()
    }, []);

    function getComplaints() {
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
        complaints.map((complaint, i) =>
             getUserById(complaint.userId)
        );
        setPosts(complaints)
    }


     async function getUserById(id) {

        await userService.getUserById({id: id, jwt: store.user.jwt})
             .then(response => {
                 setUsers(users => [...users, response.data])
             })
             .catch(err => {
                 console.log("BLA")
                 toastService.show("error", "Error")
             })

    }

    function  changeStatus(post,status){
        if(status==="Refused"){
            rejectComplaint(post.id)
        }else if(status==="Block"){
            changeUserActiveStatus(post.userId)
        }else if(status==="Delete"){
            if(post.isPost===false){
                deleteStory(post.postId)
            }else{
                deletePost(post.postId)
            }
        }
    }

    async function rejectComplaint(complaintId) {
        await complaintService.rejectById({id: complaintId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully rejected!")
                getComplaints()
            })
            .catch(err => {
                console.log("BLA")
                toastService.show("error", "Error")
            })
    }

    async function changeUserActiveStatus(userId) {
        console.log(userId)
        await userService.changeUserActiveStatus({id: userId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully updated!")
                getComplaints()
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }

    async function deleteStory(storyId) {
        await storyService.deleteStory({id: storyId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully updated!")
                getComplaints()
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }

    async function deletePost(postId) {
        await postService.deletePost({id: postId, jwt: store.user.jwt})
            .then(response => {
                toastService.show("success", "Successfully updated!")
                getComplaints()
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }



    return (
        <div>
            <Navigation/>
            <div style={{marginTop:'5%',marginLeft:'10%', marginRight:'10%', marginBottom:'20%'}}>
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
                                <td>bla</td>
                                <td>{post.status}</td>
                                {post.status === "Pending" && <td>
                                    <Button variant={"danger"} onClick={() => changeStatus(post, 'Delete')}>Delete
                                        content</Button>
                                    <span style={{marginLeft: '5%'}}/>
                                    <Button variant={"danger"} onClick={() => changeStatus(post, 'Block')}>Block
                                        account</Button>
                                    <span style={{marginLeft: '5%'}}/>
                                    <Button variant={"success"} onClick={() => changeStatus(post, 'Refused')}>Refuse
                                        complaint</Button>

                                </td>
                                }
                            </tr>
                        )
                    })}
                    </tbody>
                </Table>
            </div>
        </div>


    );
}

export default ComplaintPreview;