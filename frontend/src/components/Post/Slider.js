import React, { useState } from 'react'

import './../../style/Slider.css';
import { ReactComponent as ChevronRight } from './../../images/icons/chevron-right-solid.svg';
import { ReactComponent as ChevronLeft } from './../../images/icons/chevron-left-solid.svg';
import { ReactComponent as Circle } from './../../images/icons/circle.svg';
import { ReactComponent as Tag } from './../../images/icons/tag.svg';
import {useDispatch, useSelector} from "react-redux";
import {userActions} from "../../store/actions/user.actions";

const Slider = (props) => {
    const { media, showStoryCaption, storyCaption } = props;
    const [slideCount, setSlideCount] = useState(1)
    const [showTags, setShowTags] = useState(props.showTags);
    const [storyForReportFlag, setStoryForReportFlag] = useState(true)

    const store = useSelector(state => state);
    const dispatch = useDispatch()

    const nextImage = () => {
      setSlideCount((slideCount + 1 <= media.length) ? slideCount + 1 : slideCount)
        setStoryForReportFlag(true)
    }

    const previousImage = () => {
      setSlideCount((slideCount - 1 >= 1) ? slideCount - 1 : slideCount)
        setStoryForReportFlag(true)
    }

    const Tags = (props) => {
      const { photo } = props;
      return (
        <div className="tagWrapper">
          <Tag onClick={() => setShowTags(!showTags)} className="tagIcon" />
          <div className="tags"> 
            {showTags && photo.tags.length > 0 && photo.tags.map(tag => {
              return (
                <div className="tag">
                    <div className="tagText">{tag.username}</div>
                </div>
              )
            })} 
          </div>
        </div>
      )
    }

    const BackArrow = () => <ChevronLeft onClick={previousImage} className="Arrow-back" />
    const NextArrow = () => <ChevronRight onClick={nextImage} className="Arrow-next" />
    const CircleWrapper = () => {
      if(media.length < 2) return null;
      
      return (
      <div className="CircleWrapper">
        { media.map((photo) => {
          const checkActive = photo.orderNum === slideCount ? "active" : "";
          return (
            <Circle className={`Circle ${checkActive}`} />
          )
        })}
      </div>
    )}

    const StoryCaption = () => (
      <div className="storyCaptionWrapepr">
        <div className="storyCaptionText">{storyCaption}</div>
      </div>
    )

    const aa = (photo) => {
        if (storyForReportFlag) {
            dispatch(userActions.setStory({
                storyId : photo.postId
            }))
            setStoryForReportFlag(false)
            console.log(store.storyForReport)
        }

    }

    return (
      <div className="Slide" >
          {media.map((photo) => {
            if (photo.orderNum === slideCount) {
              return (
                <div key={photo.id} >
                  <img src={photo.content} alt='' onClick={ aa(photo)}/>
                  <div className="tags-caption">
                    { photo.tags.length !== 0 && <Tags photo={photo} /> }
                    { showStoryCaption && <StoryCaption /> }
                  </div>
                </div>
              )
            }
            return ''
          })}
        
        <CircleWrapper />
        <BackArrow />
        <NextArrow />
      </div>
    );
}

export default Slider;