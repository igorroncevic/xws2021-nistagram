import React, {useEffect, useState} from "react";
import axios from "axios";
import '../../style/Profile.css';
import {Button, Dropdown, Modal} from "react-bootstrap";
import FollowAndUnfollow from "./FollowAndUnfollow";
import Navigation from "../HomePage/Navigation";
import {   useParams } from 'react-router-dom'
import userService from "../../services/user.service";
import {userActions} from "../../store/actions/user.actions";

import {useDispatch, useSelector} from "react-redux";
import privacyService from "../../services/privacy.service";
import followersService from "../../services/followers.service";
import FollowersAndFollowings from "./FollowersAndFollowings";
import Switch from "react-switch";
import {BsThreeDotsVertical, FaGem} from "react-icons/all";
import BlockMuteAndNotifications from "./BlockMuteAndNotifications";



function Profile() {
    const{username}=useParams()
    const [user, setUser] =useState({});
    const [follow,setFollow] =useState( {});
    const [publicProfile,setPublicProfile]=useState(false);

    const [showModalFollowers, setModalFollowers] = useState(false);
    const [showModalFollowings, setModalFollowings] = useState(false);

    const [followers, setFollowers] = useState([]);
    const [following, setFollowings] = useState([]);
    const [posts, setPosts] = useState([]);
    const [closeFriend, setCloseFriend] = useState(false);
    const [isApprovedRequest, setIsApprovedRequest] = useState(false);
    const [isMuted, setIsMuted] = useState(false);

    const dispatch = useDispatch()
    const store = useSelector(state => state);
    const [isSSO,setIsSSO] =useState( store.user.isSSO);


    useEffect(() => {
     //   if(!props.location.state) window.location.replace("http://localhost:3000/unauthorized");
        getUserByUsername();
        getUserPrivacy();
        getFollowers()
        getFollowing()
    },[username]);



    async function getUserByUsername() {
        const response = await userService.getUserByUsername({
            username: username,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
            getUserPrivacy(response.data.id);
            getFollowers(response.data.id)
            getFollowing(response.data.id)
            checkUser(response.data.id);
        } else {
            console.log("getuserbyusername error")
        }
    }

    async function checkUser(value) {
        if (value === store.user.id) {
            setFollow(false)
        } else {
            setFollow(true)
            dispatch(userActions.followRequest({
                userId: store.user.id,
                followerId: value,
            }))
            isCloseFriend(value)
        }
    }

    async function isCloseFriend(value) {
        const response = await followersService.getFollowersConnection({
            userId: store.user.id,
            followerId: value,
        })
        if (response.status === 200) {
            console.log(response.data)
            setCloseFriend(response.data.isCloseFriends)
            setIsApprovedRequest(response.data.isApprovedRequest)
            setIsMuted(response.data.isMuted)
        } else {
            console.log("followings ne radi")
        }
    }

    async function getUserPrivacy(value) {
        const response = await privacyService.getUserPrivacy({
            userId: value,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setPublicProfile(response.data.response)
        } else {
            console.log("privacy ne radi")
        }
    }

    async function getFollowing(value) {
        const response = await followersService.getFollowing({
            userId:value,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setFollowings(response.data.users);
        } else {
            console.log("followings ne radi")
        }
    }

    async function getFollowers(value) {
        const response = await followersService.getFollowers({
            userId: value,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setFollowers(response.data.users);
        } else {
            console.log("followers ne radi")
        }

    }

    function handleModalFollowers() {
        setModalFollowers(!showModalFollowers)
    }

    function handleModalFollowings() {
        setModalFollowings(!showModalFollowings)
    }



    return (
        <div>
            <Navigation/>
            <div style={{marginLeft: '20%', marginRight: '20%',marginTop:'10%'}}>
                <div style={{margin: "18px 0px", orderBottom: "1px solid "}}>
                    <div style={{display: "flex", justifyContent: "space-around",}}>
                        <div>
                            <img style={{width: "180px", height: "160px", borderRadius: "80px"}} src={user.profilePhoto ? user.profilePhoto : ""}/>
                        </div>
                        <div>
                            <div  style={{display: "flex"}}>
                            <h4>{user.firstName} {user.lastName}</h4>
                                {follow && <div  style={{ marginLeft:'10em',color:'white'}}>
                                    <BlockMuteAndNotifications isApprovedRequest={isApprovedRequest} isMuted={isMuted}/>
                                </div>
                                }
                            </div>
                            <h4>{user.username}</h4>
                            <div style={{display: "flex"}}>
                                <h6 style={{marginTop:'9px'}}> 15 posts </h6>
                                <Button variant="link" style={{color:'black'}} onClick={handleModalFollowers}>{followers.length} followers</Button>
                                <Button variant="link"  style={{color:'black'}} onClick={handleModalFollowings}> {following.length} following </Button>
                            </div>

                            {follow &&
                                <div>

                                <FollowAndUnfollow user={user} isCloseFriends={closeFriend} funcIsCloseFriend={isCloseFriend} followers={followers} getFollowers={getFollowers}/>
                                </div>
                            }
                        </div>
                        <div>

                        </div>
                    </div>

                </div>
                {!follow &&
                <div className="gallery">
                    <img className="item"
                         src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                    <img className="item"
                         src='https://images.unsplash.com/photo-1581882898166-634d30416957?ixid=MnwxMjA3fDB8MHxzZWFyY2h8OXx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                    <img className="item"
                         src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                </div>
                }
                {follow && publicProfile &&
                <div className="gallery">
                    <img className="item"
                         src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                    <img className="item"
                         src='https://images.unsplash.com/photo-1581882898166-634d30416957?ixid=MnwxMjA3fDB8MHxzZWFyY2h8OXx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                    <img className="item"
                         src='https://images.unsplash.com/photo-1522228115018-d838bcce5c3a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZmlsZSUyMHBpY3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60'/>
                </div>
                }
                {follow && !publicProfile &&
                    <div style={{ borderTop: '1px solid black'}}>
                        <p style={{textAlign: 'center',marginTop:'6%', fontWeight:'bold'}}> This Account is Private </p>
                        <p style={{textAlign: 'center',marginTop:'2%'}}>Follow to see their photos and videos!</p>
                    </div>
                }

                <Modal show={showModalFollowers} onHide={handleModalFollowers}>
                    <Modal.Header closeButton>
                        <Modal.Title>Followers:</Modal.Title>

                    </Modal.Header>
                    <Modal.Body>
                        <FollowersAndFollowings ids={followers} handleModal={handleModalFollowers}/>
                    </Modal.Body>
                    <Modal.Footer>

                    </Modal.Footer>
                </Modal>
                <Modal show={showModalFollowings} onHide={handleModalFollowings}>
                    <Modal.Header closeButton>
                        <Modal.Title>Following:</Modal.Title>

                    </Modal.Header>
                    <Modal.Body>
                        <FollowersAndFollowings ids={following} handleModal={handleModalFollowings}/>
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