import React, {useEffect, useState} from 'react';
import ProfileInfo from "./ProfileInfo";
import userService from "../../services/user.service";
import {userActions} from "../../store/actions/user.actions";
import followersService from "../../services/followers.service";
import {useSelector} from "react-redux";

function CloseFriends() {
    const store = useSelector(state => state);
    useEffect(() => {
        getCloseFriends()
    },[]);


    async function getCloseFriends(){
        const response = await followersService.getCloseFriends({
            id: store.user.id,
            jwt: store.user.jwt,
        })

        //console.log(response)

        if(response.status === 200) {
            console.log("USPEO");
            console.log(response)
        }else{
            console.log("Nije uspeo")
        }
    }



    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>

            </div>
        </div>
    );
}

export default CloseFriends;