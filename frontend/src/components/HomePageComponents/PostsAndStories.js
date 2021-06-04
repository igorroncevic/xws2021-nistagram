import React from 'react';
import HomePage from "./HomePage";
import Stories from "../StoryCompoent/Stories";
import Posts from "../PostComponent/Posts";


function PostsAndStories() {
    return (

        <div className='home'>
            <HomePage />
            <Stories />
            <Posts/>
        </div>

    );

}export default PostsAndStories;