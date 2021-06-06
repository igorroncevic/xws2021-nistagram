import React, {Component, useEffect, useState} from "react";
import "../../style/Post.css";
import Post from "./Post";


function Posts () {
    //prvo cemo dobavljati postove iz baze sve

    const posts=[
        {id:'1',
            userId:'2dsdsd',
            isAd:false,
            type:'Post',
            description:'a cool new post',
            location:'Novi Sad, Serbia',
            createdAt: '2021-06-02T17:33:17.541716Z',
            mediaContent:'https://picsum.photos/800/1000'
        },{id:'2',
            userId:'3dsdss',
            isAd:false,
            type:'Post',
            description:'Vidite kako je lepo',
            location:'Zlatibor, Serbia',
            createdAt: '2021-06-02T17:33:17.541716Z',
            mediaContent:'https://picsum.photos/600/1000'
        }

    ];

    return(
        <div>
            {posts.map((post) => {
                return (
                    <Post post={post} userId={post.userId}/>);
            })}

        </div>

    );
}
export default Posts;