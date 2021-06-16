import React, { useState, useEffect } from "react";
import "../../style/post.css";
import Post from "./Post";
import { useSelector } from 'react-redux';
import postService from './../../services/post.service'
import toastService from './../../services/toast.service'

const Posts = (props) => {
    const [posts, setPosts] = useState([]);
    const store = useSelector(state => state)

    useEffect(()=>{
        postService.getHomepagePosts({ jwt: store.user.jwt })
            .then(response => {
                if(response.status === 200) setPosts(response.data.posts)
            })
            .catch(err => {
                toastService.show("error", "Could not retrieve homepage posts.")
            })
    }, [])

    return(
        <div>
            {posts && posts.map((post) => {
                return (
                    <Post post={post} postUser={{ id: post.userId }}/>);
            })}
        </div>

    );
}

export default Posts;