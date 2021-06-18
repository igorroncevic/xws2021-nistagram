import React, { useState, useEffect } from "react";
import Stories from "react-insta-stories";
import { Modal } from "react-bootstrap";
import Slider from './../Post/Slider';
import renderer from './StoryRenderer';
import storyService from './../../services/story.service'
import '../../style/Story.css';
import '../../style/ProfileIcon.css';



function Story(props) {
    const { story } = props;
    const [showModal, setModal] = useState(false);
    const [convertedStory, setConvertedStory] = useState([])
    const [header, setHeader] = useState({})

    // Convert story with multiple media to multiple stories with single media, to comply with react-insta-stories
    useEffect(()=>{
        const convertedStories = [];
        story.stories.forEach(singleStory => {  
            convertedStories.push(...storyService.convertStory(singleStory))
        })
        setConvertedStory(convertedStories);

        setHeader({
            username: story.username,
            profileImage: story.userPhoto
        })
    }, [])

    return (
        <div>
            <div className="story">
                <div className={true ? "storyBorder" : ""}>
                    <img className={`profileIcon big`} src={story.userPhoto} alt="profile" onClick={() => setModal(!showModal)}/>
                </div>
                <span className="accountName">{story.username}</span>
            </div>
            
            <Modal show={showModal} onHide={() => setModal(!showModal)}>
                <Stories
                    onAllStoriesEnd={() => setModal(!showModal)}
                    renderers={[renderer]}
                    stories={convertedStory} 
                    defaultInterval={10000} 
                    header={header}
                    width={600} 
                    height={800}/>
            </Modal>
        </div>
    );
}

export default Story;