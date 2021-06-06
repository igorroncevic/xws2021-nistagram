import React from 'react';
import HomePage from "./HomePage";


function Chats(props) {
    console.log(props)
    console.log(props.location.state)
    console.log(props.location.state.user)
    console.log(props.location.state.user.accessToken)
    console.log(props.location.state.user.userId)
    console.log(props.location.state.user.role)
    console.log(props.location.state.follow)

    return (
        <div className='home'>
            <HomePage />

            <h1>Cet</h1>
        </div>

    );

}export default Chats;