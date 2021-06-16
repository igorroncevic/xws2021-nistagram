import React, { useEffect, useState } from "react";
import { useSelector } from 'react-redux';
import moment from 'moment';
import "../../style/post.css";
import Slider from './Slider'
import Comment from "./Comment";
import PostHeader from './PostHeader';
import { ReactComponent as CardButton } from './../../images/icons/cardButton.svg' 
import LikesAndDislikes from "./LikesAndDislikes";
import userService from './../../services/user.service';
import toastService from './../../services/toast.service';
import PostMenu from "./PostMenu";

function Post (props) {
    const { postUser, post } = props;
    const [user, setUser] = useState({});
    const [hoursAgo, setHoursAgo] = useState(0)
    const [daysAgo, setDaysAgo] = useState(0);
    const [minutesAgo, setMinutesAgo] = useState(0)
    const [likesText, setLikesText] = useState("");
    const [dislikesText, setDislikesText] = useState("");
    const store = useSelector(state => state)

    console.log(post);
    console.log(postUser);

    useEffect(() => {
        getUserInfo()
    }, [])

    useEffect(() => {
        const currentTime = moment(new Date())
        const difference = moment.duration(currentTime.diff(post.createdAt))
        if(difference.asHours() < 24){
            difference.asHours() < 0 ? setMinutesAgo(Math.floor(difference.asMinutes())) : setHoursAgo(Math.floor(difference.asHours()))
        }else{
            setDaysAgo(Math.floor(difference.asDays()))
        }
        changeLikesText();
        changeDislikesText()
    }, [])

    useEffect(()=>{
        changeLikesText()
        changeDislikesText()
    })

    const changeLikesText = () => {
        if(post.likes.length === 0) setLikesText("no one")
        if(post.likes.length === 1) setLikesText("1 person")
        if(post.likes.length >= 2) setLikesText(post.likes.length + " people")
    }

    const changeDislikesText = () => {
        if(post.dislikes.length === 0) setDislikesText("no one")
        if(post.dislikes.length === 1) setDislikesText("1 person")
        if(post.dislikes.length >= 2) setDislikesText(post.dislikes.length + " people")
    }

    const getUserInfo = async () => {
        if(postUser && postUser.id && postUser.id !== store.user.id){
            const response = await userService.getUserById({
                id: postUser.id
            })
            
            if(response.status === 200){
                setUser(response.data)
            }else{
                toastService.show("error", "Error retrieving info about user's creator.");
            }
        }else{
            setUser({
                id: store.user.id,
                username: store.user.username,
                photo: "",
            })
        }
    }

    return(
        <div className="Post">
            <header>
                <PostHeader 
                    username={user.username} 
                    hideUsername={false}
                    caption={post.location}
                    captionSize="small"
                    image={user.profilePhoto}
                    iconSize="medium"
                    />
                <CardButton className="cardButton" />
            </header>

            <div className="slider">
                <Slider media={post.media} />
            </div>
            <PostMenu 
                isLiked={post.likes.filter(like => like.userId === store.user.id).length === 1}
                isDisliked={post.dislikes.filter(dislike => dislike.userId === store.user.id).length === 1}
            />
            <div className="likes-dislikes">
                <PostHeader 
                    hideAccountName={true} 
                    image={user.profilePhoto} 
                    captionSize="large"
                    iconSize="small" />
                <span>Liked by {likesText} and disliked by {dislikesText}. </span>
            </div>
            <div className="Post-caption">
                <strong> {user.username} </strong> {post.description}
            </div>
            <div className="comments">
                {post.comments.length > 0 ? post.comments.map((comment) => {
                    return (
                        <Comment key={comment.id} username={comment.username} comment={comment.content} />
                    );
                }) : <p className="noComments">No comments yet...</p> }
            </div>
            <div className="timePosted">
                { daysAgo < 1 ? (
                    hoursAgo < 1 ? minutesAgo + " minutes ago" : hoursAgo + " hours ago" ) :  
                    daysAgo + " days ago"
                }
            </div>
            <div className="addComment">
                <input className="commentInput" placeholder="Add a comment..." />
                <button className="postText">Post</button>
            </div>
        </div>
    );
}
export default Post;