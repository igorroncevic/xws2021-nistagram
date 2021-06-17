import React, { useState } from 'react'

import './../../style/Slider.css';
import { ReactComponent as ChevronRight } from './../../images/icons/chevron-right-solid.svg';
import { ReactComponent as ChevronLeft } from './../../images/icons/chevron-left-solid.svg';
import { ReactComponent as Circle } from './../../images/icons/circle.svg';

const Slider = (props) => {
    const { media } = props;
    const [slideCount, setSlideCount] = useState(1)

    const nextImage = () => {
      setSlideCount((slideCount + 1 <= media.length) ? slideCount + 1 : slideCount)
    }

    const previousImage = () => {
      setSlideCount((slideCount - 1 >= 1) ? slideCount - 1 : slideCount)
    }

    const BackArrow = () => <ChevronLeft onClick={previousImage} className="Arrow-back" />
    const NextArrow = () => <ChevronRight onClick={nextImage} className="Arrow-next" />
    const CircleWrapper = () => {
      if(media.length < 2) return null;
      
      return (
      <div className="CircleWrapper">
        { media.map((photo, key) => {
          const checkActive = photo.orderNum === slideCount ? "active" : "";
          return (
            <Circle className={`Circle ${checkActive}`} />
          )
        })}
      </div>
    )}

    return (
      <div className="Slide">
        {media.map((photo, key) => {
          if (photo.orderNum === slideCount) {
            return (
              <div key={photo.id}>
                <img src={photo.content} alt=''/>
              </div>
            )
          }
          return ''
        })}
        <CircleWrapper />
        {slideCount !== 1 ? <BackArrow /> : ''}
        {slideCount !== media.length ? <NextArrow /> : ''}
      </div>
    );
}

export default Slider;