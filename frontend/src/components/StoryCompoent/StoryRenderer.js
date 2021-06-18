import React, { useState, useEffect } from 'react';
import Slider from './../Post/Slider';
import { WithHeader } from 'react-insta-stories';
import moment from 'moment';


const StoryRenderer = ({ story, action, isPaused, config }) => {
    const { width, height, loader, storyStyles, header } = config;

    const [localStory, setLocalStory] = useState(story);

    useEffect(()=>{
        let timeCreated = "";

        const currentTime = moment(new Date())
        const difference = moment.duration(currentTime.diff(story.createdAt))

        if(difference.asHours() < 1){
            difference.asMinutes() < 1 ? 
                timeCreated = Math.floor(difference.asSeconds()) + "s" : 
                timeCreated = Math.floor(difference.asMinutes()) + "m"
        }else{
            timeCreated = Math.floor(difference.asHours()) + "h"
        }

        setLocalStory({
            ...story,
            header: {
                heading: header.username ? header.username : "username",
                subheading: timeCreated + " ago @ " + story.location,
                profileImage: header.profileImage ? header.profileImage : "" 
            }
        })

        action('play'); // Doesn't auto start if there are multiple stories

        console.log("useEffect")
    }, [story])

    return (
        <WithHeader story={localStory}>
            <Slider 
                showStoryCaption={true}
                storyCaption={story.description} 
                showTags={true} 
                media={story.media}
                />
        </WithHeader>
    )
}

const Tester = (story) => {
    return { condition: true, priority: 10 }
}

export default {
    renderer: StoryRenderer,
    tester: Tester,
}