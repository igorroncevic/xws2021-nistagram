import ProfileForSug from "../HomePage/ProfileForSug";
import {Button} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import userService from "../../services/user.service";
import {useSelector} from "react-redux";
import followersService from "../../services/followers.service";
import toastService from "../../services/toast.service";

function Notification(props){
    const {creatorId,userId,text,type} = props;
    const[user,setUser]=useState({});
    const[privateFollow,setPrivateFollow]=useState(false);
    const store = useSelector(state => state);

    useEffect(() => {
        getUserById()
        checkType()
    }, []);

    async function getUserById(id) {
        const response = await userService.getUserById({
            id: creatorId,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
        console.log(response.data)
            setUser(response.data)
        } else {
            console.log("getuserbyid error")
        }
    }

    function  checkType(){
        if(type==="FollowPrivate"){
            setPrivateFollow(true)
        }
    }
    async function acceptRequest() {
        const response = await followersService.updateUserConnection({
            userId: userId ,
            followerId: creatorId,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:false,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully accepted!");
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    async function removeRequest() {
        const response = await followersService.unfollow({
            userId: userId,
            followerId: creatorId,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully removed!");
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }


    return(
        <div style={{display: "flex", marginLeft:'10%'}}>
            <ProfileForSug user={user} username={user.username} caption={user.biography} urlText="Follow" iconSize="big" captionSize="small" storyBorder={true}
                           firstName={user.firstName} lastName={user.lastName} image={user.profilePhoto}/>
            <font face = "Comic Sans MS" size = "3" style={{marginRight:'5em', fontWeight:'bold'}}>{text}</font>

            {privateFollow &&
                <div  style={{display: "flex", marginLeft:'85px'}}>

                    <Button  style={{ height:'27px',  fontSize:'12px'}}  variant="success"  onClick={() => acceptRequest()}  >Accept</Button>
                    <Button  style={{marginLeft:'5px', height:'27px', fontSize:'12px'}}  variant="secondary"  onClick={() => removeRequest()} >Reject</Button>

                </div>
            }
        </div>
    );

}
export default Notification;