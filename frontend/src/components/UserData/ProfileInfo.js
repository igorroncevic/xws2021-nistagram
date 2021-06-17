import Navigation from "../HomePage/Navigation";
import React from "react";
import {ProSidebar, Menu, MenuItem, SubMenu, SidebarHeader, SidebarContent, SidebarFooter} from 'react-pro-sidebar';
import {
    CgProfile,
    FaBan,
    FaGem,
    FaHeart,
    FaHeartBroken, HiOutlinePhotograph,
    IoMdSettings, RiLockPasswordLine, SiGnuprivacyguard,
} from "react-icons/all";
import 'react-pro-sidebar/dist/css/styles.css';
import {Link} from "react-router-dom";
import CloseFriends from "./CloseFriends";

function  ProfileInfo(){
    return(
        <div  style={{display: 'grid', gridTemplateColumns: '200px auto',marginLeft: '5%', marginRight: '20%',marginTop:'4.2%'}}>
            <Navigation/>
            <div style={{marginLeft: '5%', marginRight: '20%',marginTop:'4.2%',display: 'flex'}}>
                <ProSidebar >
                    <Menu iconShape="square">
                        <MenuItem icon={<FaGem/>}>Close friends  <Link to="/closefriends" /> </MenuItem>
                        <MenuItem icon={<FaBan/>}>Blocked users <Link to="/blocked" /> </MenuItem>
                        <MenuItem icon={<FaHeart/>}>Liked posts <Link to="/liked" /> </MenuItem>
                        <MenuItem icon={<FaHeartBroken/>}>Disliked posts <Link to="/disliked" /> </MenuItem>
                        <MenuItem icon={<IoMdSettings/>}>Edit profile info  <Link to="/editProfile" /> </MenuItem>
                        <MenuItem icon={<RiLockPasswordLine/>}>Edit password  <Link to="/password" /> </MenuItem>
                        <MenuItem icon={<SiGnuprivacyguard/>}>Edit privacy  <Link to="/privacy" /> </MenuItem>
                        <MenuItem icon={<HiOutlinePhotograph/>}>Edit profile photo  <Link to="/privacy" /> </MenuItem>
                    </Menu>
                </ProSidebar>
            </div>
        </div>
    );
}export default ProfileInfo;