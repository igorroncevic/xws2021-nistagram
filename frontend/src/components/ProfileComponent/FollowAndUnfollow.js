import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import axios from "axios";
import userService from "../../services/user.service";
import followersService from "../../services/followers.service";
import {useDispatch, useSelector} from "react-redux";
//treba srediti da ne moze da zaprati sam sebe i da salje zahtev za pracenje drugima
function FollowAndUnfollow(props){
    const{user,loggedUser,followers,getFollowers}=props;
    const[follows,setFollows]=useState(false);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        setFollows(followers.some(item=>item.UserId===store.user.id))
        console.log(follows)
    }, [followers])

    async function follow() {
        const response = await followersService.follow({
            userId: store.user.id,
            followerId: user.id,
            isApprovedRequest: true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("ZAPRACENO")
            props.getFollowers();
        } else {
            console.log("NIJE ZAPRACENO")
        }
    }

    async function unfollow() {
        const response = await followersService.unfollow({
            userId: store.user.id,
            followerId: user.id,
            isApprovedRequest: true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("otpratio")
            props.getFollowers();
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