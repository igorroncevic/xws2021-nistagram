import React, {useEffect, useState} from 'react';
import {Button, Form, Modal} from "react-bootstrap";
import ReCAPTCHA from "react-google-recaptcha";
import axios from "axios";
import { useHistory } from 'react-router-dom';

import RegistrationPage from "./RegistrationPage";
const TEST_SITE_KEY = "6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI";

export function IndexPage(){
    const[details,setDetails]=useState({email:"", password:""});
    const[badCredentials,setCredentials]=useState(true);
    const[badEmail,setEmailErr]=useState(true);
    const[submitted,setSubmitted]=useState(false);
    const[reCaptcha,setCaptcha]=useState(0);
    const[logInDisabled,setLogInDisabled]=useState(false);
    const[showModal,setShowModal]=useState(false);
    const history = useHistory()

    useEffect(() => {
        document.body.style.backgroundColor = "#C0C0C0"
    });

    async function sendParams(){
        axios
            .post("http://localhost:8080/api/users/api/users/login", {
                email: details.email,
                password: details.password
            })
            .then(res => {
                alert("Login successful ");
                history.push({
                    pathname: '/home',
                    state: { user:res.data, follow:false }
                })
            })
            .catch(res => {
                if (reCaptcha >= 2) {
                    setCaptcha(reCaptcha+1);
                    setLogInDisabled(true);
                    setCredentials(false);
                } else {
                    setCaptcha(reCaptcha+1);
                    setCredentials(false);
                }
            });
    }
    function submitHandler(e){
        setSubmitted(true);
        e.preventDefault();
        if(checkEmail(details.email)){
            setEmailErr(true);
            sendParams();
        }else{
            setEmailErr(false)
        }
    }

    function handleChange(event) {
        setDetails({
            ...details,
            [event.target.name]: event.target.value,
        });

        if(event.target.name==="email" && submitted){
            if(checkEmail(event.target.value)==false) setEmailErr(false);
            else setEmailErr(true);
        }
    };

    function checkEmail(value){
        return !(value && !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,64}$/i.test(value));
    }

    function closeCaptcha(){
        setCaptcha(0);
        setLogInDisabled(false);
    }

    function handleModal(){
        setShowModal(!showModal)
    }

    function closeModal(){
        setShowModal(!showModal)
    }

    return (
        <div style={{ padding: '60px 0', margin: '0 auto', maxWidth: '320px' }}>
            <br />
            <Form onSubmit={submitHandler}>
                <Form.Group size="lg" controlId="email">
                    <Form.Label>Email</Form.Label>
                    <Form.Control autoFocus type="email" name="email" onChange={handleChange}  value={details.email} />
                    <p hidden={badEmail} style={{ color: "red" }}> Email is not valid!</p>

                </Form.Group>
                <Form.Group size="lg" controlId="password">
                    <Form.Label>Password</Form.Label>
                    <Form.Control  name="password" type="password" onChange={handleChange} value={details.password}/>
                </Form.Group>
                <p hidden={badCredentials} style={{ color: "red" }}> Invalid username or password!</p>
                <div style={{display : "flex"}}>
                    <a href={'/forgotten'} style={{'color': '#089A87',float : "right"}}> Forgot password?</a>
                </div>
                {reCaptcha >= 3 &&
                <ReCAPTCHA
                    style={{ display: "inline-block" }}
                    theme="light"
                    ref={React.createRef()}
                    sitekey={TEST_SITE_KEY}
                    onChange={closeCaptcha}
                    // asyncScriptOnLoad={asyncScriptOnLoad}
                />
                }
                <Button disabled={logInDisabled} block size="lg" onClick={submitHandler}>Login </Button>
            </Form>
            <br />
            <div style={{ display: " table" }}>
                <p style={{ display: "table-cell" }}>Don't have account?</p>
                <a style={{ display: "table-cell" }} className="nav-link" style={{ 'color': '#00d8fe', 'fontWeight': 'bold' }} href='#' name="workHours" onClick={handleModal}>Register</a>
            </div>

            <Modal show={showModal} onHide={closeModal} style={{ 'height': 650 }} >
                <Modal.Header closeButton style={{ 'background': 'silver' }}>
                    <Modal.Title>Registration</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{ 'background': 'silver' }}>
                    <RegistrationPage/>
                </Modal.Body>
            </Modal>
        </div>
    );
}