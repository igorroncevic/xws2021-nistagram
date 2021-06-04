import React from 'react';
import HomePage from "./HomePage";
import Card from "../CardComponent/Card";
import StoryContent from "./StoryContent";
import Stories from "../Story/Stories";


function PostsAndStories() {
    return (

        <div className='home'>
            <HomePage />
            <Stories />
            <h1>Search, Postovi i storiji</h1>
            <Card/>

        </div>

    );

}export default PostsAndStories;