import React from 'react';
import "./../../style/PostMenu.css"
import { ReactComponent as Inbox } from './../../images/icons/inbox.svg'
import { ReactComponent as Comments } from './../../images/icons/comments.svg'
import { ReactComponent as Bookmark } from './../../images/icons/bookmark.svg'
import { ReactComponent as BookmarkFilled } from './../../images/icons/bookmark-filled.svg'

import { ReactComponent as Heart } from './../../images/icons/heart.svg'
import { ReactComponent as HeartFilled } from './../../images/icons/heart-filled.svg'
import { ReactComponent as BrokenHeart } from './../../images/icons/broken-heart.svg'
import { ReactComponent as BrokenHeartFilled } from './../../images/icons/broken-heart-filled.svg'

const PostMenu = (props) => {
    const { isLiked, isDisliked, likeClicked, dislikeClicked, commentClicked, saveClicked, isSaved } = props;

    return (
        <div className="postMenu">
            <div className="interactions">
                { isLiked ? 
                    <HeartFilled onClick={likeClicked} fill="red" className="icon" /> : 
                    <Heart onClick={likeClicked} className="icon" /> 
                }
                { isDisliked ? 
                    <BrokenHeartFilled onClick={dislikeClicked} fill="red" className="icon" /> : 
                    <BrokenHeart onClick={dislikeClicked} className="icon" /> 
                }
                <Comments onClick={commentClicked} className="icon" />
                <Inbox className="icon" />
            </div>
            { isSaved ? 
                <BookmarkFilled onClick={saveClicked} fill="black" className="icon" /> : 
                <Bookmark onClick={saveClicked} className="icon" />
            }
        </div>
    )
}

export default PostMenu;