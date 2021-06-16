import React, {useEffect, useState} from 'react';
import ProfileForSug from "../HomePage/ProfileForSug";
import userService from "../../services/user.service";
import {useDispatch, useSelector} from "react-redux";
import {Button} from "react-bootstrap";


function FollowersAndFollowings(props) {
    const [users,setUsers]=useState([]);

    const store = useSelector(state => state);

    useEffect(() => {
        props.ids.map((id, i) =>
             getUserById(id.UserId)
        );
    }, [])

    async function getUserById(id) {
        const response = await userService.getUserById({
            id: id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
           setUsers(users=>[...users,response.data])
        } else {
            console.log("getuserbyid error")
        }
    }




    return(
        <div>
            {users.map((user, i) =>
                <Button variant="link"  onClick={props.handleModal}>
                    <ProfileForSug user={user} username={user.username} firstName={user.firstName} lastName={user.lastName}
                          caption={user.biography} urlText="Follow" iconSize="big" captionSize="small" storyBorder={true}/>
                </Button>



            )
            }
        </div>
    );

}export default FollowersAndFollowings;