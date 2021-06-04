import React from 'react';
import HomePage from "./HomePage";
import Card from "../CardComponent/Card";
import StoryContent from "../Story/StoryContent";
import Stories from "../Story/Stories";


function PostsAndStories() {
    return (

        <div className='home'>
            <HomePage />
            <Stories />
            <Card/>

        </div>

    );

}export default PostsAndStories;