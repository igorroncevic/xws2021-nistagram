import {Alert, Button, Container, Form, Table} from "react-bootstrap";
import {useState} from "react";
import axios from "axios";
import {ResetPasswordCode} from "./ResetPasswordCode";
import {PasswordChange} from "./PasswordChange";


export function ForgotPasswordPage() {
    const[steps,setSteps]=useState({step1:true, step2:false, step3:false, step4:false});
    const[email,setEmail]=useState('');
    const[emailError,setEmailError]=useState('Enter email');
    const[submitted,setSubmitted]=useState(false);
    const[success,setSuccess]=useState(false);

    const handleEmailChange = event => {
        setEmail(event.target.value);
    }

    const handleSubmit = () => {
        setSubmitted(true);
        if(emailError=="") {
           // this.sendMail();  SALJI MEJL
        }else{
            setEmailError("Please enter valid email address!");
        }

    }

    const sendMail=()=>{
        axios
            .put("http://localhost:8080/api/email/send", {
                'to': email,
                'subject': "Recover your password",
            })
            .then(res => {
                setSteps({
                    ...steps,
                    step2:true,
                    step1:false,
                });
                setEmailError('');
            }).catch(res => {
             setEmailError('Sorry, we don\'t recognize this email. Please try again.')
        })
    }

    const nextStep=()=>{
        setSteps({
            ...steps,
            step1:false,
            step2:false,
            step3:true,

        });
    }

    const setStateFromChild=()=>{
        setSteps({
            ...steps,
            step4:true,
            step3:false,
        });
    }

    const setSuccessFromChild=()=>{
        setSuccess(true);
    }


    return (
        <div style={{'background-color' : '#AEB6BF'}}>
            <div style={{ overflowY: "auto", height: "500px", width:"1000px", marginLeft:'auto', marginRight:'auto'}}>
                <Table striped bordered hover variant="dark" >
                    <tbody>
                    {!steps.step4 &&
                    <tr>
                        <td colSpan="2">
                            <p style={{textAlign: 'center', margin: 20}}> Follow these instructions if you forgot
                                your password and need to create a new one. </p>
                        </td>
                        {!steps.step2 &&  !steps.step3 &&
                        <td>
                            <a href={'/'} style={{'color': '#08B8A2', float: "right"}}> Back to login page?</a>

                        </td>
                        }
                    </tr>
                    }

                    {steps.step1 &&
                    <tr>
                        <td> Please enter your email address:</td>
                        <td>
                            <Form.Control autoFocus type="email" name="email" value={email} onChange={handleEmailChange}/>
                            {submitted && <span className="text-danger">{emailError}</span>}
                        </td>
                    </tr>

                    }
                    {steps.step1 &&
                    <tr>
                        <td >
                            <p style={{fontSize:"10px", color:"#08B8A2"}}>
                                Reset code will be sent to your email. It could take 10 to 30 seconds to be delivered.
                            </p>
                        </td>
                        <td colSpan="2">
                            <Button variant="info" style={{display:'block', margin:'auto', backgroundColor:"#08B8A2"}}  onClick={handleSubmit}> Confirm </Button>
                        </td>
                    </tr>
                    }
                    {steps.step2 &&
                    <tr>
                        <p style={{color:"#08B8A2", margin:"20px"}}>Please check your email for a text message with your reset code.</p>
                        <p style={{color:"#08B8A2", margin:"20px"}}>Your old password has been locked for security reasons.To unlock your profile you must verify your identity.</p>
                        <td colSpan="2">
                            <Button style={{backgroundColor:"#08B8A2", color:"white"}} variant="outline-primary" onClick={nextStep} >Next step</Button>
                        </td>
                    </tr>
                    }

                    {steps.step3 && <ResetPasswordCode email={email}  onChangeValue={setStateFromChild}/>}
                    {steps.step4 && <PasswordChange email={email} onChangeValue={setSuccessFromChild}/>}

                    </tbody>

                </Table>
            </div>
        </div>
    );
}