import React, { useState, useEffect } from 'react';
import ScrollMenu from "react-horizontal-scrolling-menu";
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
                if(response.status === 200) {
                    setStories(convertToComponentArray([...response.data.stories]));
                }
            })
            .catch(err => {
                console.log(err)
                toastService.show("error", "Could not retrieve homepage stories.")
            })
    }, [])
    
    const convertToComponentArray = (stories) => stories.map((story) => <Story story={story} /> )

    return (
        <div className="stories">
            <ScrollMenu 
                alignCenter={false}
                data={stories}
                wheel={true}
                hideArrows={true}
                hideSingleArrow={true}
                />
        </div>
    );
}

export default Stories;