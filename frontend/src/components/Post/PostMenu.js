import React from 'react';
import "./../../style/PostMenu.css"
import { ReactComponent as Inbox } from './../../images/icons/inbox.svg'
import { ReactComponent as Comments } from './../../images/icons/comments.svg'
import { ReactComponent as Bookmark } from './../../images/icons/bookmark.svg'

import { ReactComponent as Heart } from './../../images/icons/heart.svg'
import { ReactComponent as HeartFilled } from './../../images/icons/heart-filled.svg'
import { ReactComponent as BrokenHeart } from './../../images/icons/broken-heart.svg'
import { ReactComponent as BrokenHeartFilled } from './../../images/icons/broken-heart-filled.svg'

const PostMenu = (props) => {
    const { isLiked, isDisliked, postId } = props;

    return (
        <div className="postMenu">
            <div className="interactions">
                { isLiked ? <HeartFilled fill="red" className="icon" /> : <Heart className="icon" /> }
                { isDisliked ? <BrokenHeartFilled fill="red" className="icon" /> : <BrokenHeart className="icon" /> }
                <Comments className="icon" />
                <Inbox className="icon" />
            </div>
            <Bookmark className="icon" />
        </div>
    )
}

export default PostMenu;