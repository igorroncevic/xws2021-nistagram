import React, {useState} from "react";
import {Col, Container, FormControl, Row} from "react-bootstrap";
import axios from "axios";

function ChangePassword(props) {
    const [user, setUser] = useState(props.user);
    const[passwords,setPasswords]=useState({oldPass:'',newPass:'',repPass:''})
    const[submitted,setSubmitted]=useState(false)
    const[oldErr,setOldErr]=useState('Please enter valid old password!');
    const[newErr,setNewErr]=useState('Please enter new password!');
    const[repErr,setRepErr]=useState('Please repeat new password!');

    function changePass(){
        axios
            .post('http://localhost:8080/api/users/api/users/update_password', {
                id:'893b1f54-7b74-4476-85b6-b5d4f798fb29',
                old: passwords.oldPass,
                new: passwords.newPass,
                repeated: passwords.repPass
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
    }

    function  activateUpdateMode(event){
        setSubmitted(true);
        validatePasswords();

        if(newErr=='' && oldErr=='' && repErr==''){
            changePass();
        }
    }

    function validatePasswords(){
        if(passwords.oldPass=="" || user.password!=passwords.oldPass){
            setOldErr('Please enter valid old password!');
        }else{
            setOldErr('');
        }

        if(passwords.newPass=="" && oldErr==""){
            setNewErr('Please enter new password.')
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