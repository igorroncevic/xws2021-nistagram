import React, {useEffect, useState} from "react";
import Navigation from "../HomePage/Navigation";
import userService from "../../services/user.service";
import toastService from "../../services/toast.service";
import {useSelector} from "react-redux";
import ProfileForSug from "../HomePage/ProfileForSug";
import {Button, Dropdown, FormControl, Modal} from "react-bootstrap";
import followersService from "../../services/followers.service";
import {ReactComponent as Plus} from "../../images/icons/plus.svg";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

const Influencers = () => {
    const [influencers, setInfluencers] = useState([])
    const [renderInfluencers, setRenderInfluencers] = useState([])
    const store = useSelector(state => state);
    const [showModal, setShowModal] = useState(false);
    const [dateTime, setDateTime] = useState(new Date());
    const [startDate, setStartDate] = useState(new Date());
    const [modalUser, setModalUser] = useState({});

    useEffect(() => {
        getInfluencers()
    }, []);



    async function getInfluencers() {
        const response = await userService.getInfluencers({
            //   id: store.user.id,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
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
                console.log(user)
                let temp={id:user.id,username:user.username, firstname:user.firstName, lastname:user.lastName,profilePhoto:user.profilePhoto,isApprovedRequest:true,requestIsPending:false}
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
            let temp={id:value.id,username:value.username,firstname:value.firstName, lastname:value.lastName,profilePhoto:value.profilePhoto, isApprovedRequest:response.data.isApprovedRequest,requestIsPending:response.data.requestIsPending }
            setRenderInfluencers(renderInfluencers=>[...renderInfluencers, temp])
        } else {
            console.log("followings ne radi")
        }
    }

    function handleModal(user) {
        setModalUser(user)
        setShowModal(!showModal)

    }

    function closeModal(user) {
        setShowModal(!showModal)
    }

    async function createCampaignRequest(modalUser) {
        console.log(modalUser)
        const response = await userService.createCampaignRequest({
            agentId: store.user.id,
            influencerId:modalUser.id,
            campaignId:'',
            status:'Pending',
            postAt:dateTime,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            toastService.show("success", "Successfully created!!");

        } else {
            toastService.show("error", "Something went wrong.Please try again!");
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
                                        variant="success"  onClick={() => handleModal(user)}>Hire for compaign </Button>

                                </div>
                            }
                        </div>
                    )}

                </div>
            </div>

            <Modal show={showModal} onHide={closeModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Hire for campaign</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <div>
                        <div style={{borderBottom: '1px solid #dbe0de'}}>
                            <h6>Influencers info:</h6>
                            <tr>
                                <td>Username:</td>
                                <td ><p style={{marginLeft:'15px'}}>{modalUser.username}</p></td>
                            </tr>

                            <tr>
                                <td>First name:</td>
                                <td ><p style={{marginLeft:'15px'}}>{modalUser.firstname}</p></td>
                            </tr>
                            <tr>
                                <td>Last name:</td>
                                <td ><p style={{marginLeft:'15px'}}>{modalUser.lastname}</p></td>
                            </tr>
                        </div>
                      <tr>
                          <td>Campaign:</td>
                          <Dropdown>
                              <Dropdown.Toggle variant="link" id="dropdown-basic">
                                  <Plus className="icon" style={{maxWidth:'20px', marginLeft:'20px'}}/>
                              </Dropdown.Toggle>

                              <Dropdown.Menu>
                                  kampanja1
                                  kampanja2
                              </Dropdown.Menu>
                          </Dropdown>
                      </tr>
                        <tr>
                            <td>Date and time posted at:  </td>
                            <DatePicker
                                onChange={date => setDateTime(date)}

                                selected={dateTime}
                                onChange={date => setDateTime(date)}
                                showTimeSelect
                                timeFormat="HH:mm"
                                timeIntervals={15}
                                timeCaption="time"
                                dateFormat="MMMM d, yyyy h:mm aa"
                                startDate={startDate}
                                minDate={startDate}
                            />

                        </tr>
                        <tr>
                            <Button
                                style={{marginLeft: '5px', marginTop: '22px', height: '32px', fontSize: '15px'}}
                                variant="success"  onClick={() => createCampaignRequest(modalUser)}>Create campaign request</Button>
                        </tr>
                    </div>

                </Modal.Body>
            </Modal>
        </div>
    );
}
export  default  Influencers;
