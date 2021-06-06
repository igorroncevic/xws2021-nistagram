import React, {useState} from "react";
import "../../style/post.css";
import Post from "./Post";


function Posts (props) {
    //prvo cemo dobavljati postove iz baze sve
    const [user, setUser] = useState({...props.user});

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