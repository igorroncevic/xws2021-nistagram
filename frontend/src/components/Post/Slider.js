import React, { useState } from 'react'

import './../../style/Slider.css';
import faChevronRight from './../../images/icons/chevron-right-solid.svg';
import faChevronLeft from './../../images/icons/chevron-left-solid.svg';

const Slider = (props) => {
    const { media } = props;
    const [slideCount, setSlideCount] = useState(0)

    const nextImage = () => {
        setSlideCount((slideCount + 1 <= media.length) ? slideCount + 1 : slideCount)
    }

    const previousImage = () => {
        setSlideCount((slideCount - 1 >= 0) ? slideCount - 1 : slideCount)
    }

    const BackArrow = (props) => (
        <div onClick={props.previousImage} className="Arrow">
           <img src={faChevronLeft} alt="left" />
        </div>
    )

    const NextArrow = (props) => (
        <div onClick={props.nextImage} className="Arrow">
          <img src={faChevronRight} alt="right" />
        </div>
    )

    return (
      <div className="Slide">
        {slideCount !== 0 ? <BackArrow previousImage={previousImage}/> : ''}
        {media.map((photo, key) => {
          if (media.indexOf(photo) === slideCount) {
            return (
              <div key={photo.id}>
                <img src={photo.content} alt=''/>
              </div>
            )
          }
          return ''
        })}
        {slideCount !== (media.length - 1) ? <NextArrow nextImage={nextImage}/> : ''}
      </div>
    );
}

export default Slider;