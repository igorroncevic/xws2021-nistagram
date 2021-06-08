import React, {useState} from 'react';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button, Modal} from "react-bootstrap";
import RegistrationPage from "../../pages/RegistrationPage";



function NewPost(props) {
    const [user,setUser] =useState(props.location.state.user);
    const[description,setDescription]=useState('');
    const[location,setLocation]=useState('');
    const[image,setImage]=useState('');
    const [hashtagList, setHashtagList] = useState([{ text: ""}]);
    const[showModal,setShowModal]=useState(false);
    const [tagList, setTagList] = useState([{ userId: ""}]);
    // console.log("new post:");
    // console.log(user);

    //todo hashtag list
    function handleHashtagListChange (e, index) {
        const { name, value } = e.target;
        const list = [...hashtagList];
        hashtagList[index][name] = value;
        setHashtagList(list);
    }

    function handleRemoveHashtagListClick (index) {
        const list = [...hashtagList];
        list.splice(index, 1);
        setHashtagList(list);
    }

    function handleAddHashtagListClick() {
        setHashtagList([...hashtagList, { text: ""}]);
    }

    //todo tags list
    function handleTagListChange (e, index) {
        const { name, value } = e.target;
        const list = [...tagList];
        tagList[index][name] = value;
        setTagList(list);
    }

    // handle click event of the Remove button
    function handleRemoveTagListClick (index) {
        const list = [...tagList];
        list.splice(index, 1);
        setTagList(list);
    }

    // handle click event of the Add button
    function handleAddTagListClick() {
        setTagList([...tagList, { userId: ""}]);
    }


    const postDetails = ()=>{
        axios
            .post('http://localhost:8080/api/content/posts', {
                id:'1',
                userId : user.id,
                isAd : false,
                type : 'Post',
                description : description,
                location : location,
                createdAt : "2017-01-15T01:30:15.01Z",
                media : [
                    {
                        id : "1",
                        type: "Image",
                        postId : "1",
                        content: image,
                        orderNum : 1,
                        tags: [
                            {
                                mediaId : "",
                                userId : "123",
                                username : "213eus"
                            }
                        ]
                    }
                ],
                comments : [],
                likes : [],
                dislikes : []
                // comments : '',
                // likes : 0,
                // dislikes: 0,
            })
            .then(res => {
                console.log(" RADI")

            }).catch(res => {
            console.log("NE RADI")
        })
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        var self = this;
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function(upload) {
            setImage(upload.target.result)
        };
        reader.readAsDataURL(file);
    }

    function handleModal(){
        setShowModal(!showModal)
    }

    function closeModal(){
        setShowModal(!showModal)
    }

    return (
        <div className='home'>
            <Navigation user={user} />

            <div className="card input-filed"
                 style={{ margin:"30px auto",maxWidth:"500px",padding:"20px", textAlign:"center", marginTop: "5%" }} >
                <input type="text" placeholder="description" value={description} onChange={(e)=>setDescription(e.target.value)} />
                <br/>
                <input type="text" placeholder="location" value={location} onChange={(e)=>setLocation(e.target.value)} />


                <h3>Hashtags:</h3>
                <div>
                    {hashtagList.map((x, i) => {
                        return (
                            <div className="box">
                                <input
                                    name="text"
                                    placeholder="text"
                                    value={x.text}
                                    onChange={e => handleHashtagListChange(e, i)}
                                />
                                <div className="btn-box">
                                    {hashtagList.length !== 1 && <button
                                        className="mr10"
                                        onClick={() => handleRemoveHashtagListClick(i)}>Remove</button>}
                                    {hashtagList.length - 1 === i && <button onClick={handleAddHashtagListClick}>Add</button>}
                                </div>
                            </div>
                        );
                    })}
                    {/*<div style={{ marginTop: 20 }}>{JSON.stringify(hashtagList)}</div>*/}
                </div>
                <br/><br/>
                {/*<Button type={"outline-primary"}   onClick={handleModal} style={{ maxWidth:"150px", textAlign:"center"}}*/}
                {/*>Add file</Button>*/}
                <input type="file" name="file"
                       className="upload-file"
                       id="file"
                       onChange={handleChangeImage}
                       formEncType="multipart/form-data"
                       required />

                <br/><br/>
                <h3>Tags:</h3>
                <div>
                    {tagList.map((x, i) => {
                        return (
                            <div className="box">
                                <input
                                    name="userId"
                                    placeholder="userId"
                                    value={x.userId}
                                    onChange={e => handleTagListChange(e, i)}
                                />
                                <div className="btn-box">
                                    {tagList.length !== 1 && <button
                                        className="mr10"
                                        onClick={() => handleRemoveTagListClick(i)}>Remove</button>}
                                    {tagList.length - 1 === i && <button onClick={handleAddTagListClick}>Add</button>}
                                </div>
                            </div>
                        );
                    })}
                    {/*<div style={{ marginTop: 20 }}>{JSON.stringify(hashtagList)}</div>*/}
                </div>

                <Button type={"primary"}   onClick={()=>postDetails()}> Submit post  </Button>

            </div>
            {/*<Modal show={showModal} onHide={closeModal} style={{ 'height': 650 }} >*/}
            {/*    <Modal.Header closeButton style={{ 'background': 'silver' }}>*/}
            {/*        <Modal.Title>Post files</Modal.Title>*/}
            {/*    </Modal.Header>*/}
            {/*    <Modal.Body style={{ 'background': 'silver' }}>*/}
            {/*        <div className="file-field input-field">*/}
            {/*            <div className="btn #64b5f6 blue darken-1">*/}
            {/*                /!*<span>Upload Image</span>*!/*/}
            {/*                <input type="file" name="file"*/}
            {/*                       className="upload-file"*/}
            {/*                       id="file"*/}
            {/*                       onChange={handleChangeImage}*/}
            {/*                       formEncType="multipart/form-data"*/}
            {/*                       required />*/}
            {/*            </div>*/}
            {/*        </div>*/}
            {/*    </Modal.Body>*/}
            {/*</Modal>*/}
        </div>

    );

}export default NewPost;