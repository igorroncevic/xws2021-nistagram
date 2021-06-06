import Navigation from "./Navigation";
import Sidebar from "./Sidebar";
import Posts from "../PostComponent/Posts";
import "../../style/home.css";
import Stories from "../StoryCompoent/Stories";
import axios from "axios";
import {useEffect, useState} from "react";

function Home(props) {
    const[user,setUser]=useState('');
    useEffect(() => {
        getUser();
    },[])

    function getUser(){
        axios
            .post('http://localhost:8080/api/users/api/users/searchByUser', {
                    username:props.location.state.user.username
            })
            .then(res => {
                console.log("RADI")
                setUser(res.data.users[0])
            }).catch(res => {
            console.log("NE RADIs")
        })
    }
    return (
        <div className="App">
            <Navigation user={user}/>
            <main>
                <div>
                    <Stories/>
                <div className="container">
                    <Posts />
                    <Sidebar/>
                </div>
                </div>
            </main>
        </div>

    );
}export default Home;