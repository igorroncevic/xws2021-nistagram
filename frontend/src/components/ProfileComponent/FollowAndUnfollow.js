import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import followersService from "../../services/followers.service";
import {useDispatch, useSelector} from "react-redux";
import Switch from "react-switch";

function FollowAndUnfollow(props) {
    const {user, followers, getFollowers} = props;
    const [follows, setFollows] = useState(false);
    const [closeFriend, setCloseFriend] = useState(false);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        props.getFollowers(store.followers.followerId)
        getFollowersConnection()
    }, [])

    useEffect(() => {
        setFollows(followers.some(item => item.UserId === store.user.id))
    }, [followers])


    async function getFollowersConnection() {
        const response = await followersService.getFollowersConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
        })

        if (response.status === 200) {
            setFollows(response.data.isApprovedRequest)
            setCloseFriend(response.data.isCloseFriends)
            props.getFollowers(store.followers.followerId)
        } else {
            console.log("followings ne radi")
        }
    }

    async function follow() {
        const response = await followersService.createConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("ZAPRACENO")
            props.getFollowers(store.followers.followerId)
            getFollowersConnection()
        } else {
            console.log("NIJE ZAPRACENO")
        }
    }

    async function unfollow() {
        const response = await followersService.unfollow({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            props.getFollowers(store.followers.followerId)
            getFollowersConnection()
        } else {
            console.log("NIJE otpratio")
        }
    }

    function handleCloseFriends() {
        setCloseFriend(!closeFriend)
        console.log(closeFriend)

        setCloseFriends()
    }

    async function setCloseFriends() {
        const response = await followersService.updateUserConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            isCloseFriends: !closeFriend,
            isNotificationEnabled:true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("JESTE")
        } else {
            console.log("NIJE ")
        }
    }


    return (
        <div>
            {!follows ?
                <Button variant="primary" style={{margin: "10px"}} onClick={follow}>Follow</Button>
                :
                <div>
                    <div className='row'>
                        <p  style={{ marginLeft:'15px',marginRight:'3em', color:'#64f427'}}  >Close friend: </p>
                    <Switch onChange={handleCloseFriends} checked={closeFriend}/>
                    </div>
                    <Button style={{margin: "10px", marginRight: '78px'}} onClick={unfollow}>UnFollow</Button>

                </div>

            }
        </div>
    );
}export default FollowAndUnfollow;