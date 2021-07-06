import React, { useEffect, useState } from 'react';
import {Button, Modal, ListGroup, Table, Dropdown} from "react-bootstrap";
import { useSelector } from 'react-redux';
import toastService from './../../services/toast.service'
import './../../style/CollectionsModal.css'
import Navigation from "../HomePage/Navigation";
import complaintService from "../../services/complaint.service";
import userService from "../../services/user.service";
import storyService from "../../services/story.service";
import postService from "../../services/post.service";
import PostPreviewModal from "../Post/PostPreviewModal";
import {ReactComponent as Plus} from "../../images/icons/plus.svg";
import DatePicker from "react-datepicker";

const ComplaintPreview = (props) => {
    const [complaints, setComplaints] = useState([]);
    const [users,setUsers]=useState([]);
    const store = useSelector(state => state);
    const [post, setPost] = useState({})
    const [story, setStory] = useState({})
    const [showModal, setShowModal] = useState(false);
    const [showModalStory, setShowModalStory] = useState(false);
    const [modalStory, setModalStory] = useState(false);

    useEffect(() => {
        getComplaints()
    }, []);

    function getComplaints() {
        complaintService.getAllContentComplaints({ jwt: store.user.jwt })
            .then(response => {
                if(response.status === 200)
                    console.log(response.data)
                    setComplaints(response.data.contentComplaints)
            })
            .catch(err => {
                toastService.show("error", "Error")
            })
    }


     async function getPostById(id) {
         await postService.getPostById({id: id, jwt: store.user.jwt})
             .then(response => {
                 setPost(response.data)
                 setShowModal(true);
             })
             .catch(err => {
                 console.log("BLA")
                 toastService.show("error", "Error")
             })
    }

    async function getStoryById(id) {
        await storyService.getStoryById({id: id, jwt: store.user.jwt})
            .then(response => {
                console.log(response.data)
                setStory(response.data)
                setModalStory(response.data)
                setShowModalStory(!showModalStory)

            })
            .catch(err => {
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

    function closeModalStory() {
        setShowModalStory(!showModalStory)
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
                    {complaints.map((complaint,index) => {
                        return (
                            <tr>
                                <td>{index+1}</td>
                                <td>{complaint.category}</td>
                                <td>
                                    { complaint.isPost===true ?
                                        <Button variant="link" style={{color: 'black'}} onClick={() =>getPostById(complaint.postId)} >post</Button>
                                        :
                                        <Button variant="link" style={{color: 'black'}} onClick={() =>getStoryById(complaint.postId)} >story</Button>
                                    }
                                </td>
                                <td>bla</td>
                                <td>{complaint.status}</td>
                                {complaint.status === "Pending" && <td>
                                    <Button variant={"danger"} onClick={() => changeStatus(complaint, 'Delete')}>Delete
                                        content</Button>
                                    <span style={{marginLeft: '5%'}}/>
                                    <Button variant={"danger"} onClick={() => changeStatus(complaint, 'Block')}>Block
                                        account</Button>
                                    <span style={{marginLeft: '5%'}}/>
                                    <Button variant={"success"} onClick={() => changeStatus(complaint, 'Refused')}>Refuse
                                        complaint</Button>

                                </td>
                                }
                            </tr>
                        )
                    })}
                    </tbody>
                </Table>
            </div>

            <PostPreviewModal
                post={post}
                postUser={{id: post.userId}}
                showModal={showModal}
                setShowModal={setShowModal}/>

            <Modal show={showModalStory} onHide={closeModalStory}>
                <Modal.Header closeButton>
                    <Modal.Title>Story</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <img  src={modalStory} alt="document photo" style={{width:'200px', height: '200px'}}/>
                </Modal.Body>
            </Modal>
        </div>


    );
}

export default ComplaintPreview;