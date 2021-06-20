import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Button, Modal } from "react-bootstrap";
import { useParams } from 'react-router-dom'

import FollowAndUnfollow from "./FollowAndUnfollow";
import Navigation from "../HomePage/Navigation";
import userService from "../../services/user.service";
import { userActions } from "../../store/actions/user.actions";
import FollowersAndFollowings from "./FollowersAndFollowings";
import BlockMuteAndNotifications from "./BlockMuteAndNotifications";
import PostPreviewGrid from './../Post/PostPreviewGrid';
import Spinner from './../../helpers/spinner';

import privacyService from "../../services/privacy.service";
import followersService from "../../services/followers.service";
import postService from './../../services/post.service';
import toastService from './../../services/toast.service';

import '../../style/Profile.css';


const Profile = () => {
    const { username } = useParams()

    const [loading, setLoading] = useState(true);

    const [user, setUser] = useState({});
    const [follow, setFollow] = useState({});
    const [publicProfile,setPublicProfile] = useState(false);

    const [showModalFollowers, setModalFollowers] = useState(false);
    const [showModalFollowings, setModalFollowings] = useState(false);
    const [followers, setFollowers] = useState([]);
    const [following, setFollowings] = useState([]);

    const [closeFriend, setCloseFriend] = useState(false);
    const [isApprovedRequest, setIsApprovedRequest] = useState(false);
    const [isMuted, setIsMuted] = useState(false);
    const [isNotificationEnabled, setNotifications] = useState(false);

    const [posts, setPosts] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        (async function(){
            const tempUser = await getUserByUsername(); // Since it doesn't get saved in time for other requests
            if(tempUser) {
                getUserPrivacy(tempUser.id);
                getFollowers(tempUser.id)
                getFollowing(tempUser.id)
                checkUser(tempUser.id);
                getUserPrivacy(tempUser.id);
                getFollowers(tempUser.id);
                getFollowing(tempUser.id);
                getPosts(tempUser.id);
            }
        })();
    }, [username]);

    const getPosts = async (userId) => {
        const response = await postService.getPostsForUser({
            jwt: store.user.jwt,
            userId: userId
        })
        
        if (response.status === 200){ 
            setPosts([...response.data.posts])
            setLoading(false);
        }
        else{
            console.log(response);
            toastService.show("error", "Could not retrieve user's posts.")
        }
    }

    async function getUserByUsername() {
        const response = await userService.getUserByUsername({
            username: username,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
            return response.data
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
            setCloseFriend(response.data.isCloseFriends)
            setIsApprovedRequest(response.data.isApprovedRequest)
            setIsMuted(response.data.isMuted)
            setNotifications(response.data.isNotificationEnabled)
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
        if(publicProfile || isApprovedRequest || !follow)
         setModalFollowers(!showModalFollowers)
    }

    function handleModalFollowings() {
        if(publicProfile || isApprovedRequest || !follow)
            setModalFollowings(!showModalFollowings)
    }

    return (
        <div>
            <Navigation/>
            <div style={{marginLeft: '20%', marginRight: '20%',marginTop:'10%'}}>
                <div style={{margin: "18px 0px", orderBottom: "1px solid "}}>
                    <div style={{display: "flex", justifyContent: "space-around",}}>
                        <div>
                            <img alt="" style={{width: "180px", height: "160px", borderRadius: "80px"}} src={user.profilePhoto ? user.profilePhoto : ""}/>
                        </div>
                        <div>
                            <div  style={{display: "flex"}}>
                            <h4>{user.firstName} {user.lastName}</h4>
                                {follow && <div  style={{ marginLeft:'10em',color:'white'}}>
                                    <BlockMuteAndNotifications isApprovedRequest={isApprovedRequest} isMuted={isMuted} isNotificationEnabled={isNotificationEnabled}/>
                                </div>
                                }
                            </div>
                            <h4>{user.username}</h4>
                            <div style={{display: "flex"}}>
                                <h6 style={{marginTop:'9px'}}>{posts.length} posts </h6>
                                <Button variant="link" style={{color:'black'}} onClick={handleModalFollowers}>{followers.length} followers</Button>
                                <Button variant="link"  style={{color:'black'}} onClick={handleModalFollowings}> {following.length} following </Button>
                            </div>

                            { follow && <FollowAndUnfollow user={user} isCloseFriends={closeFriend} funcIsCloseFriend={isCloseFriend} followers={followers}  getFollowers={getFollowers}/> }
                        </div>
                    </div>

                </div>
                {/*prikazi kad sam na svom profilu*/}
                {!follow &&
                (loading ?
                        <div style={{ position: "relative", left: "45%", marginTop: "50px" }}>
                            <Spinner type="MutatingDots" height="100" width="100" />
                        </div> :
                        <PostPreviewGrid posts={posts} />
                )
                }
                {follow && publicProfile  &&
                    (loading ? 
                        <div style={{ position: "relative", left: "45%", marginTop: "50px" }}>
                            <Spinner type="MutatingDots" height="100" width="100" />
                        </div> :
                        <PostPreviewGrid posts={posts} />
                    )
                }


                {follow && !publicProfile && !isApprovedRequest &&
                    <div style={{ borderTop: '1px solid black'}}>
                        <p style={{textAlign: 'center', marginTop:'6%', fontWeight:'bold'}}> This Account is Private </p>
                        <p style={{textAlign: 'center', marginTop:'2%'}}>Follow to see their photos and videos!</p>
                    </div>
                }

                <Modal show={showModalFollowers} onHide={handleModalFollowers}>
                    <Modal.Header closeButton>
                        <Modal.Title>Followers:</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <FollowersAndFollowings ids={followers} handleModal={handleModalFollowers}/>
                    </Modal.Body>
                </Modal>
                <Modal show={showModalFollowings} onHide={handleModalFollowings}>
                    <Modal.Header closeButton>
                        <Modal.Title>Following:</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <FollowersAndFollowings ids={following} handleModal={handleModalFollowings}/>
                    </Modal.Body>
                </Modal>

            </div>
        </div>
    );
}

export default Profile;
