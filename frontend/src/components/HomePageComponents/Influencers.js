import ProfileInfo from "../UserData/ProfileInfo";
import Spinner from "../../helpers/spinner";
import PostPreviewGrid from "../Post/PostPreviewGrid";
import React, {useEffect, useState} from "react";
import Navigation from "../HomePage/Navigation";
import Notification from "../Notifications/Notification";
import userService from "../../services/user.service";
import toastService from "../../services/toast.service";
import {useSelector} from "react-redux";
import ProfileForSug from "../HomePage/ProfileForSug";
import {Button} from "react-bootstrap";
import followersService from "../../services/followers.service";

const Influencers = () => {
    const [influencers, setInfluencers] = useState([])
    const [renderInfluencers, setRenderInfluencers] = useState([])
    const store = useSelector(state => state);

    useEffect(() => {
        getInfluencers()
    }, []);



    async function getInfluencers() {
        const response = await userService.getInfluencers({
            //   id: store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log(response.data.users)
            setInfluencers(response.data.users)
            checkConnection(response.data.users)
        } else {
            toastService.show("error", "Something went wrong.Please try again!");
        }
    }

     function checkConnection(users){
        users.map((user, i) => {
            if(user.isProfilePublic==false){
                GetFollowersConnection(user)
            }else{
                let temp={username:user.username, firstname:user.firstname, lastname:user.lastname,profilePhoto:user.profilePhoto,isApprovedRequest:true,requestIsPending:false}
                setRenderInfluencers(renderInfluencers=>[...renderInfluencers, temp])

            }
        });
    }

    async function GetFollowersConnection(value) {
        const response = await followersService.getFollowersConnection({
            userId: store.user.id,
            followerId: value.id,
        })
        if (response.status === 200) {
            let temp={username:value.username, firstname:value.firstname, lastname:value.lastname,profilePhoto:value.profilePhoto, isApprovedRequest:response.data.isApprovedRequest,requestIsPending:response.data.requestIsPending }
            setRenderInfluencers(renderInfluencers=>[...renderInfluencers, temp])
        } else {
            console.log("followings ne radi")
        }
    }

    return (
        <div style={{marginTop: '6%'}}>
            <Navigation/>
            <div style={{marginLeft: '20%', marginRight: '20%'}}>
                <h3 style={{borderBottom: '1px solid black'}}>Influencers</h3>
                <div style={{marginTop: '4%'}}>
                    {renderInfluencers.map((user, i) =>
                        <div style={{display: 'flex', borderBottom: '1px solid #dbe0de', marginTop: '5px'}}>
                            <ProfileForSug user={user} username={user.username} firstName={user.firstName}
                                           lastName={user.lastName} urlText="Follow" iconSize="big" captionSize="small" caption="influencer"
                                           image={user.profilePhoto} storyBorder={true}/>
                            {(!user.isProfilePublic && !user.isApprovedRequest && ! user.requestIsPending) &&
                                <div>
                                    <p style={{fontSize: '0.75em', paddingLeft: '250px',paddingBottom: '0.2em', paddingTop: '1.5em', color: 'red'}}>
                                        This account is private. Follow for more info.</p>
                                </div>
                            }
                            { user.requestIsPending &&
                                <div>
                                    <p style={{fontSize: '0.75em', paddingLeft: '250px',paddingBottom: '0.2em', paddingTop: '1.5em', color: 'green'}}>
                                        Follow request is pending</p>
                                </div>
                            }
                            { (user.isProfilePublic || (!user.isProfilePublic && user.isApprovedRequest && !user.requestIsPending)) &&
                                <div style={{paddingLeft: '250px'}}>
                                    <Button
                                        style={{marginLeft: '5px', marginTop: '22px', height: '32px', fontSize: '15px'}}
                                        variant="success">Hire for compaign </Button>
                                </div>
                            }
                        </div>
                    )}

                </div>
            </div>
        </div>
    );
}
export  default  Influencers;
