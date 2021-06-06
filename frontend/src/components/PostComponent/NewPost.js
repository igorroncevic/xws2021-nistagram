import React, {useState} from 'react';
import HomePage from "../HomePageComponents/HomePage";
import axios from "axios";



function NewPost() {
    const[description,setDescription]=useState('');
    const[image,setImage]=useState('');

    const postDetails = ()=>{
        axios
            .post('http://localhost:8080/api/posts', {
                id:'1',
                userId : '',
                isAd : false,
                type : '',
                description : description,
                createdAt : new Date(),
                media : ' ',
                comments : '',
                likes : 0,
                dislikes: 0,
            })
            .then(res => {
                console.log(" RADI")

            }).catch(res => {
            console.log("NE RADI")
        })
    }

    return (
        <div className='home'>
            <HomePage />

            <div className="card input-filed"
                 style={{ margin:"30px auto",maxWidth:"500px",padding:"20px", textAlign:"center" }} >
                <input type="text" placeholder="description" value={description} onChange={(e)=>setDescription(e.target.value)} />
                <div className="file-field input-field">
                    <div className="btn #64b5f6 blue darken-1">
                        <span>Upload Image</span>
                        <input type="file" onChange={(e)=>setImage(e.target.files[0])} />
                    </div>
                </div>
                <button className="btn waves-effect waves-light #64b5f6 blue darken-1"   onClick={()=>postDetails()}> Submit post  </button>
            </div>
        </div>

    );

}export default NewPost;