import React from 'react';
import HomePage from "./HomePage";
import Card from "../CardComponent/Card";


function PostsAndStories() {
    return (

        <div className='home'>
            <HomePage />

            <h1>Search, Postovi i storiji</h1>
            <Card/>

        </div>

    );

}export default PostsAndStories;