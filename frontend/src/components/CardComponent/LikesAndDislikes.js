import React, {useState} from "react";

function LikesAndDislikes(){
    const[likes,setLikes]=useState(0);
    const[disLikes,setDisLikes]=useState(0);
    const[boolParams,setBoolParams]=useState({isLiked:false, isDisliked:false});

    function  submitLikes(){
        if(boolParams.isDisliked==false && boolParams.isLiked==false) {
            document.getElementById('like').style.backgroundColor = 'red'
            setBoolParams({
                ...boolParams,
                isLiked: true,
            });
            setLikes(likes + 1);
        }else if(boolParams.isLiked==true){
            document.getElementById('like').style.backgroundColor = 'white'
            setBoolParams({
                ...boolParams,
                isLiked: false,
            });
            setLikes(likes - 1);
        }
    }
    function submitDislikes(){
        if(boolParams.isDisliked==false && boolParams.isLiked==false) {

            document.getElementById('dislike').style.backgroundColor = 'red'
            setBoolParams({
                ...boolParams,
                isDisliked: true,
            });
            setDisLikes(disLikes + 1);
        }else if(boolParams.isDisliked==true){
            document.getElementById('dislike').style.backgroundColor = 'white'
            setBoolParams({
                ...boolParams,
                isDisliked: false,
            });
            setDisLikes(disLikes - 1);
        }
    }

    return(
        <div>
            <button id='like' className="big" onClick={submitLikes}> 👍</button>
            <span> or </span>
            <button id='dislike' className="big"  onClick={submitDislikes}>👎</button>
            <div style={{fontSize:'small'}}>Number of likes: {likes} <br /> Number of dislikes: {disLikes}</div>
        </div>

    );


}export default  LikesAndDislikes;