import React, {useEffect, useState} from 'react';
import ProfileInfo from "./ProfileInfo";
import {useDispatch, useSelector} from "react-redux";
import likeService from "../../services/like.service";
import toastService from "../../services/toast.service";
import Post from "../Post/Post";

function Disliked() {
    const dispatch = useDispatch()
    const store = useSelector(state => state);
    const [posts, setPosts] = useState([]);


    useEffect(() => {
        if(store.user.role === 'Admin' || store.user.role === "") window.location.replace("http://localhost:3000/unauthorized");
        getUserLikedOrDislikedPosts();
    }, []);

    function getUserLikedOrDislikedPosts() {
        likeService.getUserLikedOrDislikedPosts({ jwt: store.user.jwt , userId : store.user.id, isLike : false})
            .then(response => {
                if(response.status === 200) setPosts(response.data.posts)
            })
            .catch(err => {
                toastService.show("error", "Could not retrieve disliked posts.")
            })
    }

    return (
        <div  style={{display: 'flex'}}>
            <ProfileInfo />
            <div style={{marginRight: '20%',marginTop:'5%',display: 'flex', flexDirection: 'column'}}>
                {posts && posts.map((post) => {
                    return (
                        <Post post={post} postUser={{ id: post.userId }}/>);
                })}
            </div>
        </div>
    );
}

export default Disliked;