import React, { useState, useEffect } from 'react';
import HorizontalScroll from "react-scroll-horizontal";
import Story from "./Story";
import '../../style/Stories.css';
import { useSelector } from 'react-redux';
import storyService from './../../services/story.service'
import toastService from './../../services/toast.service'

const Stories = () => {
    const [stories, setStories] = useState([]);
    const store = useSelector(state => state);

    useEffect(()=>{
        storyService.getHomepageStories({ jwt: store.user.jwt })
            .then(response => {
                if(response.status === 200) setStories([...response.data.stories])
                console.log(response.data.stories)
            })
            .catch(err => {
                console.log(err)
                toastService.show("error", "Could not retrieve homepage stories.")
            })
    }, [])

    return (
        <div className="stories">
            {/* <HorizontalScroll className="scroll" reverseScroll={false}> */}
                {stories.map((story) => <Story story={story} /> )}
            {/* </HorizontalScroll> */}
        </div>
    );
}

export default Stories;