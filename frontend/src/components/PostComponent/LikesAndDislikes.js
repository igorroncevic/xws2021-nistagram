import React, {useState} from "react";
import axios from "axios";
//TODO:moras da uvedes proveru ako je vec lajkovan neki post, po novom logovanju da mu ne da opet da lajkuje
function LikesAndDislikes(props){
    const{post}=props;
    const[numOfLikes,setNumOfLikes]=useState(0);
    const[numberOfDislikes,setNumOfDislikes]=useState(0);

    const[likes,setLikes]=useState([]); //ovo dolazi sa beka, i izvuci cemo kolika je lista i to je num of likes
    const[disLikes,setDisLikes]=useState([])
    const[boolParams,setBoolParams]=useState({isLiked:false, isDisliked:false});

    /*treba da dobavimo lajkove posta na osnovu id-a
    useEffect(() => {
            getLikes();
    });
    */
    function getLikes(){
        axios
            .get('http://localhost:8080/api/likes/'+post.id)
            .then(res => {
                console.log("RADI")
                //treba da nam vrati username od ussera koji su lajkovali
            }).catch(res => {
            console.log("NE RADI")
        })
    }

    function  submitLikes(){
        if(boolParams.isDisliked==false && boolParams.isLiked==false) {
            document.getElementById('like').style.backgroundColor = 'red'

            setBoolParams({
                ...boolParams,
                isLiked: true,
            });
            setNumOfLikes(numOfLikes + 1);
        }else if(boolParams.isLiked==true){
            document.getElementById('like').style.backgroundColor = 'white'
            setBoolParams({
                ...boolParams,
                isLiked: false,
            });
            setNumOfLikes(numOfLikes - 1);
        }
    }
    function submitDislikes(){
        if(boolParams.isDisliked==false && boolParams.isLiked==false) {

            document.getElementById('dislike').style.backgroundColor = 'red'
            setBoolParams({
                ...boolParams,
                isDisliked: true,
            });
            setNumOfDislikes(numberOfDislikes + 1);
        }else if(boolParams.isDisliked==true){
            document.getElementById('dislike').style.backgroundColor = 'white'
            setBoolParams({
                ...boolParams,
                isDisliked: false,
            });
            setNumOfDislikes(numberOfDislikes - 1);
        }
    }

    return(
        <div>
            <button id='like' className="big" onClick={submitLikes}> 👍</button>
            <span> or </span>
            <button id='dislike' className="big"  onClick={submitDislikes}>👎</button>
            <div style={{fontSize:'small'}}>Number of likes: {numOfLikes} <br /> Number of dislikes: {numberOfDislikes}</div>
        </div>

    );


}export default  LikesAndDislikes;