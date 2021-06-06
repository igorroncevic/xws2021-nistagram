import React, {useEffect, useState} from "react";
import '../../style/Profile.css';
import {Button, Modal} from "react-bootstrap";
import EditProfile from "./EditProfile";
import Post from "../PostComponent/Post";
import ChangePassword from "./ChangePassword";
import FollowAndUnfollow from "./FollowAndUnfollow";
import axios from "axios";
import Navigation from "../HomePage/Navigation";
//user ce se pokupiti kroz props isto kao i follow true/false
//kad idem na myprofile, saljem mu sebe i follow false
//kad stisnem na nekog saljem mu tog nekog i saljem mu follow true
//kad ulazis u bilo ciji profil trebaju ti info o pratiocima kao i info o postovima
//ako je tudji profil ne trebaju ti update funkcionalnosti, treba ti mogucnost follow/unfollow


function Profile(props) {

   const [user,setUser] =useState(props.location.state.user);
   const [follow,setFollow] =useState(props.location.state.follow);

    useEffect(() => {
        setFollow(props.location.state.follow);
    }, [])
    useEffect(() => {
        setUser(props.location.state.user);
    }, [])

    const [image, setImage] = useState('');
    const [showModal, setModal] = useState(false);
    const [showModalPass, setModalPass] = useState(false);
    const [followers, setFollowers] = useState([]);
    const [following, setFollowings] = useState([]);
    const [posts, setPosts] = useState([]);

    console.log("PROFILE")
    console.log(user);
    console.log(follow);

    const updatePhoto = (file) => {
        setImage(file)
    }

    useEffect(() => {
        getFollowers()
        getFollowing()
        getPosts()
    });
    function getFollowing(){
        axios
            .post('http://localhost:8080/api/followers/get_followings', {
                UserId:'5190c16f-7886-4fad-9d76-ef0b5e304639'
            })
            .then(res => {
                console.log("RADI")
                setFollowers(res);
            }).catch(res => {
            console.log("NE RADIs")
        })
    }

    function getFollowers(){
        axios
            .post('http://localhost:8005/api/followers/get_followers', {
                    UserId:'5190c16f-7886-4fad-9d76-ef0b5e304639'
            })
            .then(res => {
                console.log("RADI")
                setFollowings(res);
            }).catch(res => {
            console.log("NE RADIs")
        })
    }

    function getPosts(){
        axios
            .get('http://localhost:8080/api/content/api/posts',+user.id)
            .then(res => {
                console.log("RADI")
                setPosts(res);
            }).catch(res => {
            console.log("NE RADIs")
        })
    }

    function handleModal() {
        setModal(!showModal)
    }
    function handleModalPass() {
        setModalPass(!showModalPass)
    }

    const postsMock = [
        {
            id: '1',
            userId: '2dsdsd',
            isAd: false,
            type: 'Post',
            description: 'a cool new post',
            location: 'Novi Sad, Serbia',
            createdAt: '2021-06-02T17:33:17.541716Z',
            mediaContent: 'https://picsum.photos/800/1000'
        }, {
            id: '2',
            userId: '3dsdss',
            isAd: false,
            type: 'Post',
            description: 'Vidite kako je lepo',
            location: 'Zlatibor, Serbia',
            createdAt: '2021-06-02T17:33:17.541716Z',
            mediaContent: 'https://picsum.photos/600/1000'
        }

    ];

    return (
        <div>
            <Navigation user={user}/>
            <div style={{marginLeft: '20%', marginRight: '20%',marginTop:'15%'}}>
                <div style={{margin: "18px 0px", orderBottom: "1px solid "}}>
                    <div style={{display: "flex", justifyContent: "space-around",}}>
                        <div>
                            <img style={{width: "180px", height: "160px", borderRadius: "80px"}}
                                 src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                        </div>
                        <div>
                            <h4>{user.firstName} {user.lastName}</h4>
                            <h4>{user.username}</h4>
                            <div style={{display: "flex"}}>
                                <h6> 15 posts </h6>
                                <h6 style={{marginLeft:'13px'}}> {followers.length} followers </h6>
                                <h6 style={{marginLeft:'13px'}}> {following.length} following </h6>
                            </div>
                            {follow ?
                                <FollowAndUnfollow following={following}/>
                                :
                                <div>
                                    <Button variant="link" style={{marginTop:'2em', borderTop: '1px solid red', display: "flex",  justifyContent: "space-between", width: "108%", color: 'red', float: "right"}} onClick={handleModal}>Update profile info?</Button>
                                    <Button variant="link" style={{borderBottom: '1px solid red', display: "flex",  justifyContent: "space-between", width: "108%", color: 'red', float: "right"}} onClick={handleModalPass}>Change password?</Button>
                                </div>


                            }
                        </div>
                        <div>

                        </div>
                    </div>

                </div>

                <div className="gallery">
                    
            <img className="item"
                 src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>

                <img className="item"
                     src='https://images.unsplash.com/photo-1581882898166-634d30416957?ixid=MnwxMjA3fDB8MHxzZWFyY2h8OXx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>

                    <img className="item"
                         src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>

                    {/*  {
                mypics.map(item=>{
                    return(
                        <img key={item._id} className="item" src={item.photo} alt={item.title}/>
                    )
                })
            }
                               {postsMock .map(item => {
                        return (
                            <div>
                                <Post post={item}/>
                            </div>
                        )
                    })
                    }
        */}



                    <Modal show={showModal} onHide={handleModal}>
                        <Modal.Header closeButton>
                            <Modal.Title>Edit profile</Modal.Title>
                        </Modal.Header>
                        <Modal.Body>
                            <EditProfile user={user}/>
                        </Modal.Body>
                        <Modal.Footer>

                        </Modal.Footer>
                    </Modal>

                    <Modal show={showModalPass} onHide={handleModalPass}>
                        <Modal.Header closeButton>
                        </Modal.Header>
                        <Modal.Body>
                            <ChangePassword user={user}/>
                        </Modal.Body>
                        <Modal.Footer>

                        </Modal.Footer>
                    </Modal>

                </div>

            </div>
        </div>
    );
}


export default Profile