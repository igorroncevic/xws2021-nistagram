import React from 'react';
import HomePage from "./HomePage";
import Card from "../CardComponent/Card";
import Story from "../FrontPageComponents/Story";


function PostsAndStories() {
    return (

        <div className='home'>
            <HomePage />
            {/*<Story />*/}
            <h1>Search, Postovi i storiji</h1>
            <Card/>

        </div>

    );

}export default PostsAndStories;