import React, {useEffect, useState} from "react";
import {Button} from "react-bootstrap";
import axios from "axios";
//ovde dolazim ako sam usla u profil nekog drugog usera
//treba da mu posaljem moj id, i da uzmem njegov id kako bih proverila da li se pratimo, ili ne
function FollowAndUnfollow(props){
    const{user,loggedUser,followers,getFollowers}=props;
    const[follows,setFollows]=useState(false);

    useEffect(() => {
        setFollows(followers.some(item=>item.id==loggedUser.id))
    }, [followers])

    function follow(){
        axios
            .post('http://localhost:8005/api/followers/create_connection', {
                 UserId :loggedUser.id,
                 FollowerId : user.id,
                 IsMuted :false,
                 IsCloseFriends :false,
                 IsApprovedRequest :false,
                 IsNotificationEnabled : true
            })
            .then(res => {
              console.log("ZAPRACENO")
                props.getFollowers();
            }).catch(res => {
            console.log("NIJE ZAPRACENO")
        })
    }

    function unfollow(){
        axios
            .post('http://localhost:8005/api/followers/delete_directed', {
                UserId :loggedUser.id,
                FollowerId : user.id,
                IsMuted :false,
                IsCloseFriends :false,
                IsApprovedRequest :false,
                IsNotificationEnabled : true
            })
            .then(res => {
                console.log("otpratio")
                props.getFollowers();
            }).catch(res => {
            console.log("NIJE otpratio")
        })
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