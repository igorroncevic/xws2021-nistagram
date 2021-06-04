import React from 'react';
import HomePage from "./HomePage";
import Stories from "../Story/Stories";
import Posts from "../CardComponent/Posts";


function PostsAndStories() {
    return (

        <div className='home'>
            <HomePage />
            <Stories />
            <Posts/>
        </div>

    );

}export default PostsAndStories;