import React from "react";
import Sidebar from "./Sidebar";
import Posts from "./../Post/Posts";
import Stories from "../StoryCompoent/Stories";

const Feed = () => {
    return(
        <div>
            <Stories />
            <div className="container">
                <Posts />
                {/*<Sidebar/>*/}
            </div>
        </div>
    )
}

export default Feed;