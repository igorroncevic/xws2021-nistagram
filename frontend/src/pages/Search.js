import React, {useEffect, useState} from 'react';
import PasswordStrengthBar from 'react-password-strength-bar';
import {Alert, Button, Dropdown, DropdownButton, FormControl, InputGroup} from "react-bootstrap";
import axios from "axios";

export default function Search() {
    // Declare a new state variable, which we'll call "count"
    const [searchCategory, setSearchCategory] = useState("Search category");

    // Similar to componentDidMount and componentDidUpdate:
    // useEffect(() => {
    //     let response = axios.get('http://localhost:8080/security/passwords');
    //     if(response && response.status && response.status === 200)
    //         setBlacklistedPasswords([...response.data]);
    //     else
    //         console.log("No blacklisted passwords.")
    // }, []);


    return (
        <div  className="App">
            <h1 style={{marginLeft : "30px"}}>Search</h1>
            <br/>

            <InputGroup  style={{marginLeft : "30px", width : "30%"}}>
                <FormControl/>

                <DropdownButton onSelect={(e) => setSearchCategory(e) }
                    as={InputGroup.Append}
                    variant="outline-secondary"
                    title={searchCategory}
                    id="input-group-dropdown-2"
                >
                    <Dropdown.Item eventKey={"user"} >User</Dropdown.Item>
                    <Dropdown.Item eventKey={"location"} >Location</Dropdown.Item>
                    <Dropdown.Item eventKey={"tag"} >Tag</Dropdown.Item>
                </DropdownButton>

                <Button variant="primary">Search</Button>{' '}
            </InputGroup>
        </div>
    );
}