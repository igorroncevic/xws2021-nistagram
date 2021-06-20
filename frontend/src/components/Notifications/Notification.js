import ProfileForSug from "../HomePage/ProfileForSug";
import {Button} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import userService from "../../services/user.service";
import {useSelector} from "react-redux";
import followersService from "../../services/followers.service";
import toastService from "../../services/toast.service";
import moment from "moment";
import "../../style/notification.css";

function Notification(props){
    const {id,creatorId,userId,text,type,createdAt} = props;
    const[user,setUser]=useState({});
    const[privateFollow,setPrivateFollow]=useState(false);
    const store = useSelector(state => state);
    const [hoursAgo, setHoursAgo] = useState(0)
    const [daysAgo, setDaysAgo] = useState(0);
    const [minutesAgo, setMinutesAgo] = useState(0)

    useEffect(() => {
        setDate()
        getUserById()
        checkType()
    }, []);

    async function getUserById(id) {
        const response = await userService.getUserById({
            id: creatorId,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
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
            userId: creatorId ,
            followerId: userId,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:false,
            requestIsPending:false,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully accepted!");
            deleteNotification()
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    function handleReject(){
        unfollow()
        deleteNotification()
    }

    async function unfollow() {
        const response = await followersService.unfollow({
            userId: creatorId,
            followerId:  userId,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully rejected!");
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    async function deleteNotification() {
        const response = await userService.deleteNotification({
            id: id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            props.getUserNotifications()
        } else {
            toastService.show("error", "Something went wrong, please try again!");
        }
    }

    function setDate(){
        const currentTime = moment(new Date())

        if(Math.floor(moment.duration(currentTime.diff(createdAt)).asHours()) < 24){
            if (Math.floor(moment.duration(currentTime.diff(createdAt)).asHours()) <= 0) {
                setMinutesAgo(Math.floor(moment.duration(currentTime.diff(createdAt)).asMinutes()))
            }
            else{
                setHoursAgo(Math.floor(moment.duration(currentTime.diff(createdAt)).asHours()))}
        }else{
            setDaysAgo(Math.floor(moment.duration(currentTime.diff(createdAt)).asDays()))
        }
    }


    return(
        <div style={{display: "flex", marginLeft:'10%'}}>
            <ProfileForSug user={user} username={user.username} caption={user.biography} urlText="Follow" iconSize="big" captionSize="small" storyBorder={true}
                           firstName={user.firstName} lastName={user.lastName} image={user.profilePhoto}/>
            <font face = "Comic Sans MS" size = "3" style={{marginRight:'5em', fontWeight:'bold'}}>{text}</font>
            <div className="timePosted">
                { daysAgo < 1 ? (
                        hoursAgo < 1 ? <p style={{fontSize:'11px'}}>{minutesAgo}  minutes ago</p>: <p> {hoursAgo}  hours ago</p>
                    ) : <p style={{fontSize:'11px'}}>{daysAgo} days ago</p>
                }
            </div>
            {privateFollow &&
                <div  style={{display: "flex", marginLeft:'85px'}}>

                    <Button  style={{ height:'27px',  fontSize:'12px'}}  variant="success"  onClick={() => acceptRequest()}  >Accept</Button>
                    <Button  style={{marginLeft:'5px', height:'27px', fontSize:'12px'}}  variant="secondary"  onClick={() => handleReject()} >Reject</Button>

                </div>
            }
        </div>
    );

}
export default Notification;