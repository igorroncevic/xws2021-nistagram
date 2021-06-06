import React from 'react';
import HomePage from "./HomePage";
import Stories from "../StoryCompoent/Stories";
import Posts from "../PostComponent/Posts";


function PostsAndStories(props) {
  // console.log(props.location.state);
    console.log(props.location.state.user)
   // const{user}=props.location.state.user;
    return (

        <div >
            <HomePage user={props.location.state.user}/>
            <Stories />
            <Posts/>
        </div>

    );

}export default PostsAndStories;