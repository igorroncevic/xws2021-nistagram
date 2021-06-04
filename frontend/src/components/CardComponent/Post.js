import React, {Component, useEffect, useState} from "react";
import "../../style/Card.css";
import axios from "axios";
import {Button} from "react-bootstrap";
import {LikeButton, UpdownButton} from '@lyket/react';
import Comments from "./Comments";
import LikesAndDislikes from "./LikesAndDislikes";


function Post (props) {
    const{post,userId}=props;
    const[user,setUser]=useState({username:'majanokti123'});
    const[likes,setLikes]=useState([]);
    const[numOfLikes,setNumOfLikes]=useState(0);

    /*treba da dobavimo usera na osnovu id-a
    useEffect(() => {
            getUserInfo();
\            getLikes();
    });
    */


    function getUserInfo(){
        axios
            .get('http://localhost:8080/api/users/getUserById'+userId)
            .then(res => {
                console.log("RADI")
                //setUser();
            }).catch(res => {
            console.log("NE RADI")
        })
    }

    function getLikes(){
        axios
            .get('http://localhost:8080/api/likes/'+post.id)
            .then(res => {
                console.log("RADI")
                //treba da nam vrati username od ussera koji su lajkovali
            }).catch(res => {
            console.log("NE RADI")
        })
    }


    return(
            <div className="Post">
                <header>
                    <div className="Post-user">
                        <div className="Post-user-avatar">
                            <img src="https://picsum.photos/800/1000" alt={user.username}/>
                        </div>
                        <div className="Post-user-nickname">
                            <span>{user.username}</span>
                        </div>
                    </div>
                </header>

                <div className="Post-image">
                    <div className="Post-image-bg">
                        <img alt="Icon Living" src={post.mediaContent} />
                    </div>
                </div>
                <LikesAndDislikes/>

                <div className="Post-caption">
                    <strong>{user.username} </strong>{post.description}
                </div>

                <Comments post={post}/>


            </div>);
}
export default Post;