import React, {useEffect, useState} from 'react';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button, Modal} from "react-bootstrap";
import RegistrationPage from "../../pages/RegistrationPage";
import UserAutocomplete from "../PostComponent/UserAutocomplete";
import ProfileForAutocomplete from "../PostComponent/ProfileForAutocomplete";
import AutocompleteHashtags from "../PostComponent/AutocompleteHashtags";
import userService from "../../services/user.service";
import {useDispatch, useSelector} from "react-redux";

function NewPost(props) {
    // const [user,setUser] =useState(props.location.state.user);
    const [user,setUser] =useState({});

    const[description,setDescription]=useState('');
    const[location,setLocation]=useState('');
    const[image,setImage]=useState('');
    const [hashtagList, setHashtagList] = useState([]);
    const [hashtagListForPrint, setHashtagListForPrint] = useState([]);
    const[showModal,setShowModal]=useState(false);
    const [tagList, setTagList] = useState([]);
    const [tagListForPrint, setTagListForPrint] = useState([]);
    const [mediaList, setMediaList] = useState([]);
    const [postPrint, setPostPrint] = useState([]);
    const [imageName, setImageName] = useState("");
    const [allUsers, setAllUsers] = useState([]);
    const [allHashtags, setAllHashtags] = useState([]);

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    useEffect(() => {
        //if(!props.location.state) window.location.replace("http://localhost:3000/unauthorized");
        console.log(props);
        //setUser(props.location.state.user);
        getUserInfo();
        getAllUsers();
        getAllHashtags();
        console.log("all users:")
        console.log(allUsers)
    }, []);

    async function getUserInfo() {
        const response = await userService.getUserById({
            id: store.user.id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setUser(response.data)
        } else {
            console.log("getuser error")
        }
    }


    function setupHeaders(){
        return {
            Accept: 'application/json',
            Authorization: 'Bearer ' + store.user.jwt,
        }
    }

    function getAllUsers() {
        console.log(setupHeaders())
        axios.get("http://localhost:8080/api/users/api/users", {
            headers : setupHeaders()
        }).then(res => setAllUsers(res.data.users));
    }

    function getAllHashtags() {
        axios.get("http://localhost:8080/api/content/hashtag/get-all", {
            headers : setupHeaders()
        }).then(res => setAllHashtags(res.data.hashtags));
    }

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


    function handleTagAutocompleteClick(tag) {
        if (tagListForPrint.some((someTag) => someTag.id === tag.id)) //prevents adding the same tag
            return;
        // setTagList([...tagList, {userId: tag.userId, username: tag.username, mediaId: "1"}]);
        setTagListForPrint([...tagListForPrint, tag]);
        setTagList([...tagList, {userId : tag.id, mediaId : "1", username : tag.username}]);
    }

    function handleHashtagAutocompleteClick(tag) {
        if (hashtagListForPrint.some((someTag) => someTag.id === tag.id)) //prevents adding the same tag
            return;
        // setTagList([...tagList, {userId: tag.userId, username: tag.username, mediaId: "1"}]);
        setHashtagListForPrint([...hashtagListForPrint, tag]);
        setHashtagList([...hashtagList, {id : tag.id, text : tag.text}]);
    }

    function handleHashtagAutocompleteNewSuggestion(newTag) {
        setHashtagListForPrint([...hashtagListForPrint, {text : newTag}]);
        setHashtagList([...hashtagList, {text : newTag}]);
    }


    const postDetails = ()=>{
        let date = new Date();
        let month = date.getMonth() + 1;
        if (month < 10)
            month = "0" + month;
        let jsonDate = date.getFullYear() + "-" + month + "-" + date.getDate() + "T01:30:15.01Z";
        console.log(mediaList)
        axios
            .post('http://localhost:8080/api/content/posts', {
                id:'1',
                userId : user.id,
                isAd : false,
                type : 'Post',
                description : description,
                location : location,
                createdAt : jsonDate,
                media : mediaList,
                comments : [],
                likes : [],
                dislikes : [],
                hashtags : hashtagList
            }, {
                headers : setupHeaders()

            })
            .then(res => {
                alert("Post created successfully!");

            }).catch(res => {
            console.log("NE RADI")
        })
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        setImageName(evt.target.files[0].name);
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
        setTagList([]);
        setTagListForPrint([]);
        setImage("");
        setShowModal(!showModal)
    }

    async function saveModal() {
        let tagListFilter = await tagList.filter(tag => tag.userId !== "");
        let media = {
            id : "1",
            type: "Image",
            postId : "1",
            content: image,
            orderNum : 1,
            tags: tagListFilter
        };
        setMediaList([...mediaList, media]);
        setPostPrint([...postPrint, {filename : imageName, tags : tagListFilter}]);

        setTagList([]);
        setTagListForPrint([]);
        setImage("");
        closeModal();
    }

    return (
        <div className='home'>
            <Navigation user={user} />

            <div className="card input-filed"
                 style={{ margin:"30px auto",maxWidth:"500px",padding:"20px", textAlign:"center", marginTop: "5%" }} >
                <input type="text" placeholder="description" value={description} onChange={(e)=>setDescription(e.target.value)} />
                <br/>
                <input type="text" placeholder="location" value={location} onChange={(e)=>setLocation(e.target.value)} />

                <br/>
                <AutocompleteHashtags addToHashtaglist={handleHashtagAutocompleteClick}
                                  suggestions={allHashtags} handleHashtagAutocompleteNewSuggestion={handleHashtagAutocompleteNewSuggestion}
                />
                <br/><br/>
                <h3>Hashtags:</h3>
                <div>
                    <ul>
                    {hashtagListForPrint.map((hashtag, i) => {
                        return (
                            <li>
                                {hashtag.text}
                            </li>
                        );
                    })}
                    </ul>
                    {/*<div style={{ marginTop: 20 }}>{JSON.stringify(hashtagList)}</div>*/}
                </div>
                <br/><br/>
                <Button type={"outline-primary"}   onClick={handleModal} style={{ maxWidth:"150px", textAlign:"center"}}
                >Add file</Button>
                {postPrint.map((x, i) => {
                    return (
                        <div className="box">
                            filename : {x.filename} <br/>
                            tag number : {x.tags.length}
                            <br/><br/>
                        </div>

                    );
                })}
                <br/><br/>
                <Button type={"primary"}   onClick={()=>postDetails()}> Submit post  </Button>

            </div>
            <Modal show={showModal} onHide={closeModal} style={{ 'height': 650 }} >
                <Modal.Header closeButton style={{ 'background': 'silver' }}>
                    <Modal.Title>Post files</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{ 'background': 'silver' }}>
                    <input type="file" name="file"
                           className="upload-file"
                           id="file"
                           onChange={handleChangeImage}
                           formEncType="multipart/form-data"
                           required />

                    <br/><br/>
                    <UserAutocomplete addToTaglist={handleTagAutocompleteClick}
                        suggestions={allUsers}
                    />
                    <h3>Tags:</h3>
                    <div>
                        <ul>
                        {tagListForPrint.map((tag, i) => {
                            return (
                                <li>
                                    <ProfileForAutocomplete username={tag.username} firstName={tag.firstName} lastName={tag.lastName}  caption={tag.biography} urlText="Follow" iconSize="medium" captionSize="small" storyBorder={true} />
                                </li>
                            );
                        })}
                        </ul>
                        <br/><br/>
                        <Button type={"primary"}   onClick={()=>saveModal()}> Save  </Button>
                    </div>
                </Modal.Body>
            </Modal>
        </div>

    );

}export default NewPost;