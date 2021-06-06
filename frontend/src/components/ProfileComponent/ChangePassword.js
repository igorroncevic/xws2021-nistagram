import React, {useState} from "react";
import {Col, Container, FormControl, Row} from "react-bootstrap";
import axios from "axios";
import PasswordStrengthBar from "react-password-strength-bar";

function ChangePassword(props) {
    const [user, setUser] = useState(props.user);
    const[passwords,setPasswords]=useState({oldPass:'',newPass:'',repPass:''})
    const[submitted,setSubmitted]=useState(false)
    const[oldErr,setOldErr]=useState('');
    const[newErr,setNewErr]=useState('');
    const[repErr,setRepErr]=useState('');
    const [passwordStrength, setPasswordStrength] = useState("");

    function changePass(){
        axios
            .post('http://localhost:8080/api/users/api/users/update_password', {
                password: {
                    Id: '5190c16f-7886-4fad-9d76-ef0b5e304639',
                    OldPassword: passwords.oldPass,
                    NewPassword: passwords.newPass,
                    RepeatedPassword: passwords.repPass
                }
            })
            .then(res => {
                console.log("RADI")

            }).catch(res => {
            console.log("NE RADI")
        })
    }



    function handleChange(event) {
        console.log(event.target.name)
        console.log(event.target.value)
        setPasswords({
            ...passwords,
            [event.target.name]: event.target.value,
        });
        validationErrorMessage(event);
    }
    function validationErrorMessage(event) {
        const { name, value } = event.target;
        switch (name) {
            case 'newPass':
                setNewErr(checkPassword(passwords.newPass) ? 'Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!' : '')
                break;
            case 'repPass':
                setRepErr( isValidRepeatedPassword(passwords.repPass) ? '' : 'This password must match the previous!')
                break;
            default:
                break;
        }
    }
    function isValidRepeatedPassword (value)  {
        var pera= passwords.newPass === passwords.repPass;
        return pera;
    }

    function  activateUpdateMode(event){
        setSubmitted(true);
        validatePasswords();

        if(newErr=='' && oldErr=='' && repErr==''){
            changePass();
        }
    }
    function checkPassword (password) {
        console.log("Checking")
        if(/^(?=.*[\d])(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*])[\w!@#$%^&*]{8,}$/.test(password)){
            setPasswordStrength(password);
            return false;
        //else if(blacklistedPasswords.includes(password)){
         //   setPasswordStrength(password);
         //   return false;
        } else {
            setPasswordStrength("");
            return true;
        }
    }

    function validatePasswords(){
        if(passwords.oldPass=="" || user.password!=passwords.oldPass){
            setOldErr('Please enter valid old password!');
        }else{
            setOldErr('');
        }

        if(passwords.newPass=="" && oldErr==""  ){
            setNewErr('Please enter new password.')
        }else if(checkPassword(passwords.newPass)){
            setNewErr('Password must contains at least 8 characters (lowercase letter, capital letter, number and special character) or not be a common password!')

        }else{
            setNewErr('')
        }

        if(passwords.newPass!=passwords.repPass && oldErr==""){
             setRepErr('This password must match the previous.')
        }else{
              setRepErr('')

        }
    }
        return (
            <Container>
                <h2 className="pt-4 pb-3">Change Password</h2>
                <Row className="m-2">
                    <FormControl name="oldPass" type="password" placeholder="Please enter old password"  value={passwords.oldPass} onChange={handleChange}/>
                     {submitted &&  <label className="text-danger">{oldErr}</label>}
                </Row>
                <Row className="m-2">
                    <FormControl name="newPass" type="password" placeholder="Enter new Password" value={passwords.newPass} onChange={handleChange}/>
                    <PasswordStrengthBar password={passwordStrength} />

                    {submitted &&  <label className="text-danger">{newErr}</label>}

                </Row>
                <Row className="m-2">
                    <FormControl name="repPass" type="password" placeholder="Repeat new Password" value={passwords.repPass} onChange={handleChange}/>
                    {submitted &&  <label className="text-danger">{repErr}</label>}
                </Row>
               <button  style={{marginRight:'100px', float:'right'}} type="button" className="btn btn-outline-danger" onClick={activateUpdateMode}>Edit profile</button>
            </Container>
        );
}
export default ChangePassword;