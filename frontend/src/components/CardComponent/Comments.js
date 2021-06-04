import React, {useState} from "react";
import axios from "axios";
import {Button} from "react-bootstrap";

function Comments(props) {
    const{post}=props;
    const[comments,setComments]=useState([]);
    const[newComment, setNewComment]=useState('');

    function getComments(){
        axios
            .get('http://localhost:8080/api/comments/'+post.id)
            .then(res => {
                console.log("RADI")
                setComments([
                    ...comments,{
                        id:'',
                        userId:'',
                        content:'',
                        createdAt:'',
                    }
                ]);
            }).catch(res => {
            console.log("NE RADI")
        })
    }

    const commentsOne = [
        {
            userId: "raffagrassetti",
            content: "Woah dude, this is awesome! 🔥",
            id: 1,
            createdAt:'f'
        },
        {
            userId: "therealadamsavage",
            content: "Like!",
            id: 2,
            createdAt:'f'
        },
        {
            userId: "mapvault",
            content: "Niceeeee!",
            id: 3,
            createdAt:'f'
        },
    ];
    function handleSubmitComment(){
        axios
            .post('http://localhost:8080/api/comments',{
                PostId: post.id,
                UserId: post.userId,
                Content:newComment,
                CreatedAt:new Date()
            })
            .then(res => {
                console.log("RADI")
                //treba da nam vrati username od ussera koji su lajkovali
            }).catch(res => {
            console.log("NE RADI")
        })
    }




    return (
        <div className="comments">
            {commentsOne.map((comment) => {
                return (
                    <div className="commentContainer">
                        <div className="accountName">{comment.userId}</div>
                        <div className="comment">{comment.content}</div>
                    </div>

                );
            })}
            <div className="row">
                <div style={{ marginLeft: '3em'}}className="accountName"><strong>IDUSERA</strong></div>
                <input  style={{marginLeft:'2em', height:'1.4em'}} aria-label="Add a comment" autoComplete="off"  type="text"
                        name="add-comment" placeholder="Add a comment..." value={newComment} onChange={({ target }) => setNewComment(target.value)} />
                <Button size="sm" style={{marginLeft:'0.6em',fontSize :'small', width:'9em'}} variant="light" onClick={handleSubmitComment}><strong> Add </strong></Button>
            </div>

        </div>

    );
}export default Comments;