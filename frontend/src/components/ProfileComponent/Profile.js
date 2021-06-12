import React, {useEffect, useState} from "react";
import axios from "axios";
import '../../style/Profile.css';
import {Button, Modal} from "react-bootstrap";
import EditProfile from "./EditProfile";
import ChangePassword from "./ChangePassword";
import FollowAndUnfollow from "./FollowAndUnfollow";
import Navigation from "../HomePage/Navigation";


function Profile(props) {
    const [user, setUser] =useState(props.location.state.user ? props.location.state.user : {});
    const [follow,setFollow] =useState(props.location.state.follow ? props.location.state.follow : {});
    const [publicProfile,setPublicProfile]=useState(false);

    const [image, setImage] = useState('');
    const [showModal, setModal] = useState(false);
    const [showModalPass, setModalPass] = useState(false);

    const [followers, setFollowers] = useState([]);
    const [following, setFollowings] = useState([]);
    const [posts, setPosts] = useState([]);
    const [loggedUser, setLoggedUser] = useState();


    var loggedUsername = sessionStorage.getItem("username");
    var isSSO = sessionStorage.getItem("isSSO")

    useEffect(() => {
        if(!props.location.state) window.location.replace("http://localhost:3000/unauthorized");

        setUser(props.location.state.user);
        setFollow(props.location.state.follow)
        getUserByUsername();
        getUserPrivacy();
        getFollowers()
        getFollowing()
        //getPosts()
    },[]);

    function getUserByUsername(){
        axios
            .post('http://localhost:8080/api/users/api/users/searchByUser', {
                username:loggedUsername
            })
            .then(res => {
                setLoggedUser(res.data.users[0])
                console.log(loggedUser)
            }).catch(res => {
            console.log("NE RADI get user")
        })
    }

    function  getUserPrivacy(){
        axios
            .post('http://localhost:8080/api/users/api/privacy/isProfilePublic', {
                userId: user.id
            })
            .then(res => {
                setPublicProfile(res.data.response)
               // console.log("privacy radi")
            }).catch(res => {
            console.log("privacy ne radi")
        })
    }

    function getUser(){ //zbog azuriranja podataka nakon izmene profila! mora ovako jer navigation link ne moze da salje funkciju
        axios
            .post('http://localhost:8080/api/users/api/users/searchByUser', {
                username:user.username
            })
            .then(res => {
              //  console.log("RADI get user")
                setUser(res.data.users[0])
            }).catch(res => {
            console.log("NE RADI get user")
        })
    }

    const updatePhoto = (file) => {
        setImage(file)
    }

    function getFollowing(){
        axios
            .post('http://localhost:8005/api/followers/get_followings', {
                UserId:user.id
            })
            .then(res => {
                console.log("following radi")
              //  console.log(res.data.users)
                setFollowings(res.data.users);

            }).catch(res => {
            console.log("following ne radi")
        })
    }

    function getFollowers(){
        axios
            .post('http://localhost:8005/api/followers/get_followers', {
                UserId: user.id
            })
            .then(res => {
                console.log("followers radi")
                console.log(res.data.users)
                setFollowers(res.data.users);

            }).catch(res => {
            console.log("followers ne radi")
        })
    }

    function getPosts(){
        axios
            .get('http://localhost:8080/api/content/api/posts',+user.id)
            .then(res => {
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

    return (
        <div>
            <Navigation user={loggedUser}/>
            <div style={{marginLeft: '20%', marginRight: '20%',marginTop:'10%'}}>
                <div style={{margin: "18px 0px", orderBottom: "1px solid "}}>
                    <div style={{display: "flex", justifyContent: "space-around",}}>
                        <div>
                            <img style={{width: "180px", height: "160px", borderRadius: "80px"}}
                                 src={user.profilePhoto}/>
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
                                <FollowAndUnfollow user={user} loggedUser={loggedUser} followers={followers} getFollowers={getFollowers}/>
:
                                <div>
                                    <Button variant="link" style={{marginTop:'2em', borderTop: '1px solid red', display: "flex",  justifyContent: "space-between", width: "108%", color: 'red', float: "right"}} onClick={handleModal}>Update profile info</Button>
                                    { !isSSO && <Button variant="link" style={{borderBottom: '1px solid red', display: "flex",  justifyContent: "space-between", width: "108%", color: 'red', float: "right"}} onClick={handleModalPass}>Change password</Button> }
                                </div>

                            }
                        </div>
                        <div>

                        </div>
                    </div>

                </div>
                {publicProfile ?
                    <div className="gallery">
                        <img className="item"
                             src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                        <img className="item"
                             src='https://images.unsplash.com/photo-1581882898166-634d30416957?ixid=MnwxMjA3fDB8MHxzZWFyY2h8OXx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                        <img className="item"
                             src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                    </div>
                    :
                    <div style={{ borderTop: '1px solid black'}}>
                        <p style={{textAlign: 'center',marginTop:'6%', fontWeight:'bold'}}> This Account is Private </p>
                        <p style={{textAlign: 'center',marginTop:'2%'}}>Follow to see their photos and videos!</p>
                    </div>


                }
                <Modal show={showModal} onHide={handleModal}>
                    <Modal.Header closeButton>
                        <Modal.Title>Edit profile</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <EditProfile user={user} updateUser={getUser}/>
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
    );
}export default Profile;

/*
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
        {
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


*/