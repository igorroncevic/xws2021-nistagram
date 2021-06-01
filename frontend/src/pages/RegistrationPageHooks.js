import React, {useEffect, useState} from 'react';
import PasswordStrengthBar from 'react-password-strength-bar';
import {Alert, Button, FormControl} from "react-bootstrap";
import axios from "axios";

export default function RegistrationPageHooks() {
    // Declare a new state variable, which we'll call "count"
    const [password, setPassword] = useState("");
    const [passwordStrength, setPasswordStrength] = useState("");
    const [birthDate, setBirthDate] = useState("");
    const [sex, setSex] = useState("");
    const [biography, setBiography] = useState("");
    const [username, setUsername] = useState("");
    const [usernameErr, setUsernameErr] = useState("Enter username");
    const [phoneNumber, setPhoneNumber] = useState("");
    const [phoneNumberErr, setPhoneNumberErr] = useState("Enter phone number in form +38160123456");
    const [id, setId] = useState("");
    const [email, setEmail] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [rePassword, setRePassword] = useState("");
    const [emailErr, setEmailErr] = useState("Enter email");
    const [passwordErr, setPasswordErr] = useState("Enter password");
    const [blacklistedPassword, setBlacklistedPassword] = useState("Password you entered is too common");
    const [firstNameErr, setFirstNameErr] = useState("Enter first name");
    const [lastNameErr, setLastNameErr] = useState ("Enter last name");
    const [rePasswordErr, setRePasswordErr] = useState("Repeat password");
    const [validForm, setValidForm] = useState(false);
    const [submitted, setSubmitted] = useState(false);
    const [successfullyReg, setSuccessfullyReg] = useState(false);
    const [disabled, setDisabled] = useState(false);
    const [errorMessage, setErrorMessage] = useState(false);
    const [blacklistedPasswords, setBlacklistedPasswords] = useState([]);

    // Similar to componentDidMount and componentDidUpdate:
    useEffect(() => {
        let response = axios.get('http://localhost:8080/security/passwords');
        if(response && response.status && response.status === 200)
            setBlacklistedPasswords([...response.data]);
        else
            console.log("No blacklisted passwords.")
    }, []);

    function handleInputChange(event) {
        const target = event.target;
        switch (target.name) {
            case "firstName" :
                setFirstName(target.value);
                break;
            case "lastName" :
                setLastName(target.value);
                break;
            case "email" :
                setEmail(target.value);
                break;
            case "password" :
                setPassword(target.value);
                break;
            case "rePassword" :
                setRePassword(target.value);
                break;
            case "birthDate" :
                setBirthDate(target.value);
                break;
            case "sex" :
                setSex(target.value);
                break;
            case "phoneNumber" :
                setPhoneNumber(target.value);
                break;
            case "username" :
                setUsername(target.value);
                break;
            case "biography" :
                setBiography(target.value);
                break;

        }
        validationErrorMessage(event);
    }

    function validationErrorMessage(event) {
        const { name, value } = event.target;

        switch (name) {
            case 'firstName':
                setFirstNameErr(checkNameAndSurname(firstName) ? '' : 'EnterFirstName')
                break;
            case 'lastName':
                setLastNameErr(checkNameAndSurname(lastName) ? '' : 'EnterLastName')
                break;
            case 'email':
                setEmailErr(isValidEmail(email) && email.length > 1 ? '' : 'Email is not valid!')
                break;
            case 'password':
                setPasswordErr(checkPassword(password) ? 'Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!' : '')
                break;
            case 'rePassword':
                setRePasswordErr( isValidRepeatedPassword(rePassword) ? '' : 'This password must match the previous!')
                break;
            case 'phoneNumber':
                setPhoneNumberErr( isPhoneNumberValid(phoneNumber) ? '' : 'Enter phone number')
                break;
            case 'username':
                setUsernameErr( isUsernameValid(username) ? '' : 'Enter username')
                break;
            default:
                /*this.setState({
                    validForm: true
                })*/
                break;
        }

    }

    function isUsernameValid(value) {
        return /^[a-z0-9_.]+$/.test(value);
    }

    function isPhoneNumberValid(value) {
        return /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/.test(value);
    }

    function checkNameAndSurname(value) {
        return /^[a-zA-Z ,.'-]+$/.test(value);

    }

    function checkPassword (password) {
        console.log("Checking")
        if(/^(?=.*[\d])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*])[\w!@#$%^&*]{8,}$/.test(password)){
            setPasswordStrength(password);
            return false;
        }else if(blacklistedPasswords.includes(password)){
            setPasswordStrength(password);
            return false;
        } else {
            setPasswordStrength("");
            return true;
        }
    }

    function isValidEmail (value) {
        return !(value && !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,64}$/i.test(value));
    }
    function isValidRepeatedPassword (value)  {
        return password === rePassword;
    }

    async function submitForm (event) {
        setSubmitted(true);

        event.preventDefault();
        const errors = ['email', 'password', 'firstName', 'rePassword', 'lastName'];
        if (validateForm(errors)) {
            await sendParams()
        } else {
            console.log('Invalid Form')
        }
    }

    function validateForm(errors) {
        let valid = true;
        for(const Error of errors) {
            validationErrorMessage(createTarget(Error));
        }
        //Promeniti!
        if(emailErr !== "" || passwordErr !== "" || firstNameErr !== "" ||
            lastNameErr !== "" || rePasswordErr !== "" )
            return !valid;
        return valid;
    }

    function createTarget (error) {
        return {target : {value : error, name : error}}
    }

    async function sendParams() {
        axios
            .post('http://localhost:8001/api/users', {
                'id':'1',
                'firstName' : firstName,
                'lastName' : lastName,
                'email' : email,
                'username' : username,
                'password' : password,
                'role' : 'Basic',
                'birthdate' : "2017-01-15T01:30:15.01Z",
                'profilePhoto' : 'idk',
                'sex' : 'MAN',
                'isActive' : true
            })
            .then(res => {
                setErrorMessage(false);
                setSuccessfullyReg(true);
                setDisabled(!disabled);
            }).catch(res => {
                setErrorMessage(true);
                console.log("NE RADI")
        })

    }

    return (
        <div  className="App">
            {/*<h2 id="createCertifiacate"> Create certificate </h2>*/}
            <div className="row">
                <label className="col-sm-2 col-form-label">Name</label>
                <div className="col-sm-5 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""} type="text" value={firstName} name="firstName" onChange={(e) =>
                        handleInputChange(e)
                    } className="form-control" placeholder="First Name"/>
                    {submitted && firstNameErr.length > 0 &&
                    <span className="text-danger">{firstNameErr}</span>}

                </div>
                <div className="col-sm-5 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}  type="text" value={lastName} name="lastName" onChange={(e) => handleInputChange(e) } className="form-control" placeholder="Last Name"/>
                    {submitted && lastNameErr.length > 0 && <span className="text-danger">{lastNameErr}</span>}
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Birth date</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="date" value={birthDate} name="birthDate" onChange={(e) => handleInputChange(e) } className="form-control" id="birthDate" />
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Sex</label>
                <div className="col-sm-6 mb-2">
                    <select onChange={(e) => handleInputChange(e)} name={"sex"} value={sex}>
                        <option value="MALE">Male</option>
                        <option value="FEMALE">Female</option>
                        <option value="OTHER">Other</option>
                    </select>
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Phone number</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="text" value={phoneNumber} name="phoneNumber" onChange={(e) => handleInputChange(e) } className="form-control" id="phoneNumber" placeholder="+38160123456" />
                    {submitted && phoneNumberErr.length > 0 && <span className="text-danger">{phoneNumberErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Username</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="text" value={username} name="username" onChange={(e) => handleInputChange(e) } className="form-control" id="username" />
                    {submitted && usernameErr.length > 0 && <span className="text-danger">{usernameErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Biography</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="text" value={biography} name="biography" onChange={(e) => handleInputChange(e) } className="form-control" id="biography" />
                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Email</label>
                <div className="col-sm-6 mb-2">
                    <input  disabled = {(disabled)? "disabled" : ""}   type="email" value={email} name="email" onChange={(e) => handleInputChange(e) } className="form-control" id="email" placeholder="example@gmail.com" />
                    {submitted && emailErr.length > 0 && <span className="text-danger">{emailErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>
            <div className="row" style={{marginTop: '1rem'}}>
                <label className="col-sm-2 col-form-label">Password</label>
                <div className="col-sm-6 mb-2">
                    <FormControl disabled = {(disabled)? "disabled" : ""}  name="password" type="password" placeholder="Password"  value={password} onChange={(e) => handleInputChange(e) }/>
                    {submitted && passwordErr.length > 0 &&  <span className="text-danger">{passwordErr}</span>}
                    <PasswordStrengthBar password={passwordStrength} />
                </div>
                <div className="col-sm-4">
                </div>
            </div>

            <div className="row" style={{marginTop: '1rem'}}>
                <label  className="col-sm-2 col-form-label">Repeat password</label>
                <div className="col-sm-6 mb-2">
                    <FormControl  disabled = {(disabled)? "disabled" : ""}  name="rePassword" type="password" placeholder="Repeat new Password" value={rePassword} onChange={(e) => handleInputChange(e) }/>
                    {submitted && rePasswordErr.length > 0 &&  <span className="text-danger">{rePasswordErr}</span>}

                </div>
                <div className="col-sm-4">
                </div>
            </div>

            {
                successfullyReg ?
                    <Alert variant='success' show={true}  style={({textAlignVertical: "center", textAlign: "center"})}>
                        Successfully registered please login.
                    </Alert>
                    :
                    <div className="row" style={{marginTop: '1rem'}}>
                        <div className="col-sm-5 mb-2">
                        </div>
                        <div className="col-sm-4">
                            <Button variant="success" onClick={submitForm}>Confirm</Button>
                        </div>
                    </div>
            }

            {
                errorMessage &&
                <Alert variant='danger' show={true}  style={({textAlignVertical: "center", textAlign: "center"})}>
                    The e-mail address must be unique! Please try again
                </Alert>
            }
        </div>
    );
}