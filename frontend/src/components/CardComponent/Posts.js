import React, {Component, useEffect, useState} from "react";
import "../../style/Card.css";
import Post from "./Post";


function Posts () {
    //prvo cemo dobavljati postove iz baze sve
    const content=[
        {   id:'1',
            userId:'2',
            isAd:false,
            type:'Post',
            description:'a cool new post',
            location:'Novi Sad, Serbia',
            createdAt: '2021-06-02T17:33:17.541716Z',
            mediaContent:'https://picsum.photos/800/1000'
        },{id:'2',
            userId:'3',
            isAd:false,
            type:'Post',
            description:'Vidite kako je lepo',
            location:'Zlatibor, Serbia',
            createdAt: '2021-06-02T17:33:17.541716Z',
            mediaContent:'https://picsum.photos/700/1000'
        }

    ];

    return(
        <div>
            {content.map((content) => {
                return (
                    <Post content={content} userId={content.userId}/>);
            })}

        </div>

    );
}
export default Posts;