import {Button, Dropdown, Modal} from "react-bootstrap";
import {BsThreeDotsVertical} from "react-icons/all";
import React, {useEffect, useState} from "react";
import followersService from "../../services/followers.service";
import privacyService from "../../services/privacy.service";
import {useSelector} from "react-redux";
import {useHistory} from "react-router-dom";

function BlockMuteAndNotifications(props){
    const {isApprovedRequest,isMuted} = props;
    const[muted,setIsMuted]=useState({})
    const [showBlockModal, setBlockModal] = useState(false);
    const [showMuteModal, setMuteModal] = useState(false);
    const [showNotifiactionsModal, setNotificationsModal] = useState(false);
    const store = useSelector(state => state);
    const history = useHistory()

    useEffect(() => {
        setIsMuted(isMuted)
    },[isMuted]);



    function handleBlockModal(){
        setBlockModal(!showBlockModal)
    }
    function handleMuteModal(){
        setMuteModal(!showMuteModal)
    }
    function handleNotificationsModal(){
        setNotificationsModal(!showNotifiactionsModal)
    }
    async function blockUser() {
        const response = await privacyService.blockUser({
            UserId: store.user.id,
            BlockedUserId: store.followers.followerId,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
           console.log("USPEO")
            history.push({ pathname: '/home' })

        } else {
            console.log("Nije uspeo")
        }
    }

    async function muteUser() {
        const response = await followersService.updateUserConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:true,
            isNotificationEnabled:true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("JESTE mjut")
            setIsMuted(true)
            handleMuteModal()
        } else {
            console.log("NIJE mjut")
        }
    }
    async function unMuteUser() {
        const response = await followersService.updateUserConnection({
            userId: store.followers.userId,
            followerId: store.followers.followerId,
            isApprovedRequest: true,
            isCloseFriends: false,
            isMuted:false,
            isNotificationEnabled:true,
            jwt: store.user.jwt,
        })
        if (response.status === 200) {
            console.log("JESTE")
            setIsMuted(false)
            handleMuteModal()
        } else {
            console.log("NIJE ")
        }
    }

    return(
        <div>
            <Dropdown>
                <Dropdown.Toggle variant="link" id="dropdown-basic">
                    <BsThreeDotsVertical/>
                </Dropdown.Toggle>
                <Dropdown.Menu>
                    <Dropdown.Item onClick={handleBlockModal}>Block</Dropdown.Item>
                    {isApprovedRequest && <Dropdown.Item onClick={handleMuteModal}>Mute/Unmute</Dropdown.Item>}
                    {isApprovedRequest && <Dropdown.Item onClick={handleNotificationsModal}>Menage notifications</Dropdown.Item>}
                </Dropdown.Menu>
            </Dropdown>

            <Modal show={showBlockModal} onHide={handleBlockModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Block user</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <p style={{fontWeight:'bold'}}>Are you sure you want to block this user? Blocked users will not be abele to message you.</p>
                    <p style={{fontSize:'13px'}}>After blocking you will be redirected to home page.</p>
                    <br/> <br/>

                    <div style={{display:'flex',float:'right'}}>
                        <Button variant="danger" style={{marginRight:'10px'}} onClick={blockUser}>Yes</Button>
                        <Button variant="info" onClick={handleBlockModal}>No</Button>
                    </div>
                </Modal.Body>
            </Modal>

            <Modal show={showMuteModal} onHide={handleMuteModal}>
                <Modal.Header closeButton>
                    <Modal.Title>{muted ? "Unmute" :"Mute"}</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {!muted ?
                    <div>
                        <p style={{fontWeight:'bold',display: 'inline'}}>Are you sure you want to <p style={{color:'red',display: 'inline'}}>mute</p> this user?</p>
                        <div style={{display:'flex',float:'right'}}>
                            <Button variant="danger" style={{marginRight:'10px'}} onClick={muteUser}>Yes</Button>
                            <Button variant="info" onClick={handleMuteModal}>No</Button>
                        </div>
                    </div>
                    :
                    <div>
                        <div>
                            <p style={{fontWeight:'bold',display: 'inline'}}>Are you sure you want to <p style={{color:'red',display: 'inline'}}>unmute</p> this user?</p>
                            <div style={{display:'flex',float:'right'}}>
                                <Button variant="danger" style={{marginRight:'10px'}} onClick={unMuteUser}>Yes</Button>
                                <Button variant="info" onClick={handleMuteModal}>No</Button>
                            </div>
                        </div>
                    </div>}
                    <br/>


                </Modal.Body>
            </Modal>
        </div>



    );

}export default  BlockMuteAndNotifications;