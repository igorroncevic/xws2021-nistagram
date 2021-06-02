import React, {useEffect, useState} from 'react';
import {Alert, Button, Dropdown, DropdownButton, FormControl, InputGroup} from "react-bootstrap";
import axios from "axios";

export default function Search() {
    // Declare a new state variable, which we'll call "count"
    const [searchCategory, setSearchCategory] = useState("Search category");
    const [input, setInput] = useState("");
    const [inputErr, setInputErr] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");



    function searchByUser() {
        axios
            .post('http://localhost:8080/api/users/api/users/searchByUser', {
                'Username' : searchCategory,
            })
            .then(res => {
            }).catch(res => {
        })
    }

    function searchByTag() {
        axios
            .post('http://localhost:8080/api/content/api/searchByLocation', {
                'Username' : searchCategory,
            })
            .then(res => {
            }).catch(res => {
        })
    }

    function searchByLocation() {
        axios
            .post('http://localhost:8080/api/content/api/searchByLocation', {
                'Username' : searchCategory,
            })
            .then(res => {
            }).catch(res => {
        })
    }

    function search() {
        if (input === "" && searchCategory !== "User") {
            setInputErr("Enter search value")
            return
        }
        else if (input ==="" && firstName === "" && lastName === "" && searchCategory === "User") {
            setInputErr("Enter search value")
            return
        }
        switch (searchCategory) {
            case "User" :
                searchByUser();
                break;
            case "Tag" :
                searchByTag();
                break;
            case "Location" :
                searchByLocation();
                break;
            default:
                alert("Select search category");
                return;
        }
    }

    return (
        <div  className="App">
            <h1 style={{marginLeft : "30px"}}>Search</h1>
            <br/>
            <div className="row" style={{marginLeft : "10px"}}>
                <div className="col-sm-5 mb-2">
                    <input name="input" className="form-control" placeholder={"search value"} value={input} onClick={(e) => setInputErr("")} onChange={(e) => setInput(e.target.value)}/>
                    {inputErr.length > 0 &&
                    <span className="text-danger">{inputErr}</span>}
                </div>
                <div className="col-sm-5 mb-2" style={{display : "inline"}}>
                    {searchCategory === "User" &&
                    <input name="input" className="form-control" placeholder={"first name"} value={firstName} onClick={(e) => setInputErr("")} onChange={(e) => setFirstName(e.target.value)}/>}
                </div>
                <div className="col-sm-5 mb-2" style={{display : "inline"}}>
                    {searchCategory === "User" &&
                    <input name="input" className="form-control" placeholder={"last name"} value={lastName} onClick={(e) => setInputErr("")} onChange={(e) => setLastName(e.target.value)}/>}
                </div>
                <div className="col-sm-5 mb-2">
                    <DropdownButton onSelect={(e) => setSearchCategory(e) }
                                    as={InputGroup.Append}
                                    variant="outline-secondary"
                                    title={searchCategory}
                                    id="input-group-dropdown-2"
                    >
                        <Dropdown.Item eventKey={"Location"} >Location</Dropdown.Item>
                        <Dropdown.Item eventKey={"Tag"} >Tag</Dropdown.Item>
                        <Dropdown.Item eventKey={"User"} >User</Dropdown.Item>
                    </DropdownButton>
                </div>
                <div className="col-sm-4">
                    <Button variant="primary" onClick={search}>Search</Button>{' '}
                </div>
            </div>
        </div>
    );
}