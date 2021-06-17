import React, { useState, useEffect } from "react";
import { useSelector } from 'react-redux';

// TODO: moras da uvedes proveru ako je vec lajkovan neki post, po novom logovanju da mu ne da opet da lajkuje
const LikesAndDislikes = (props) => {
    const { likes, dislikes } = props;
    const [isLiked, setIsLiked] = useState(false);
    const [isDisliked, setIsDisliked] = useState(false);

    const store = useSelector(state => state);

    useEffect(()=>{
        likes && likes.forEach(like => like.userId === store.user.id ? setIsLiked(true) : null)
        dislikes && dislikes.forEach(dislike => dislike.userId === store.user.id ? setIsDisliked(true) : null)
    })

    const submitLike = () => {
        
    }

    const submitDislike = () => {
        
    }

    return(
        <div>
            <button id='like' className="big" onClick={submitLike}>👍</button>
            <button id='dislike' className="big"  onClick={submitDislike}>👎</button>
        </div>
    );
}

export default LikesAndDislikes;