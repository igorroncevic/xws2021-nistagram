import React, {Component, useEffect, useState} from "react";
import "../../style/Card.css";
import axios from "axios";


function Post (props) {
    const{content,userId}=props;
    const[user,setUser]=useState({username:'majanokti123'});
    const[comments,setComments]=useState([]);
    const[likes,setLikes]=useState([]);
    /*treba da dobavimo usera na osnovu id-a
    useEffect(() => {
            getUserInfo();
            getComments();
            getLikes();
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

    function getComments(){
        axios
            .get('http://localhost:8080/api/comments/'+content.id)
            .then(res => {
                console.log("RADI")
                setComments([
                    ...comments,{
                        content:'',
                        createdAt:'',
                    }
                ]);
            }).catch(res => {
            console.log("NE RADI")
        })
    }

    function getLikes(){
        axios
            .get('http://localhost:8080/api/likes/'+content.id)
            .then(res => {
                console.log("RADI")
                //treba da nam vrati username od ussera koji su lajkovali
            }).catch(res => {
            console.log("NE RADI")
        })
    }

    return(
            <article className="Post">
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
                        <img alt="Icon Living" src={content.mediaContent} />
                    </div>
                </div>
                <div className="Post-caption">
                    <strong>{user.username} </strong>{content.description}
                </div>
            </article>);
}
export default Post;