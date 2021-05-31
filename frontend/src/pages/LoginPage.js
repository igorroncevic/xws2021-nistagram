import React, {useState} from 'react';
import {Button, Form, Modal} from "react-bootstrap";
import ReCAPTCHA from "react-google-recaptcha";
import RegistrationPage from "./RegistrationPage";


export function LoginPage(){
    const[details,setDetails]=useState({email:"", password:""});

    const submitHandler=e=>{
        e.preventDefault();
        console.log(details.email);
        console.log(details.password);
    }
    return (
        <div style={{ padding: '60px 0', margin: '0 auto', maxWidth: '320px' }}>
            <br />
            <Form onSubmit={submitHandler}>
                <Form.Group size="lg" controlId="email">
                    <Form.Label>Email</Form.Label>
                    <Form.Control autoFocus type="email" name="email" onChange={e => setDetails({...details,email:e.target.value})} value={details.email} />
                </Form.Group>
                <Form.Group size="lg" controlId="password">
                    <Form.Label>Password</Form.Label>
                    <Form.Control  name="password" type="password" onChange={e => setDetails({...details,password:e.target.value})} value={details.password}/>
                </Form.Group>



                <Button block size="lg" onClick={submitHandler}>Login </Button>
            </Form>
            <br />


        </div>
    );
}