import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import axios from "axios";
import userService from "../../services/user.service";
import followersService from "../../services/followers.service";
import {useDispatch, useSelector} from "react-redux";
//treba srediti da ne moze da zaprati sam sebe i da salje zahtev za pracenje drugima
function FollowAndUnfollow(props){
    const{user,followers,getFollowers}=props;
    const[follows,setFollows]=useState(false);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        props.getFollowers(props.user.id)
        getFollowersConnection()
    }, [])

    useEffect(() => {
        setFollows(followers.some(item=>item.UserId===store.user.id))
    }, [followers])


    async function getFollowersConnection() {
        const response = await followersService.getFollowersConnection({
             userId : store.followers.userId,
             followerId :store.followers.followerId,
        })

        if (response.status === 200) {
            console.log(response.data)
            setFollows(response.data.isApprovedRequest)
            props.getFollowers(props.user.id)
        } else {
            console.log("followings ne radi")
        }
    }

    async function follow() {
        const response = await followersService.follow({
            userId : store.followers.userId,
            followerId :store.followers.followerId,
            isApprovedRequest: true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("ZAPRACENO")
            props.getFollowers(props.user.id)
            getFollowersConnection()
        } else {
            console.log("NIJE ZAPRACENO")
        }
    }

    async function unfollow() {
        const response = await followersService.unfollow({
            userId : store.followers.userId,
            followerId :store.followers.followerId,
            isApprovedRequest: true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("otpratio")
            props.getFollowers(props.user.id)
            getFollowersConnection()
        } else {
            console.log("NIJE otpratio")
        }
    }

    return(
        <div>
            {!follows ?
                <Button variant="primary" style={{margin: "10px"}} onClick={follow}>Follow</Button>
                :
                <Button style={{margin: "10px"}} onClick={unfollow}>UnFollow</Button>
            }
        </div>
    );
}export default FollowAndUnfollow;