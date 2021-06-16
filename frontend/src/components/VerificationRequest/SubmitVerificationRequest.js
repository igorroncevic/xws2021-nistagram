import React, {useEffect, useState} from 'react';
import { useDispatch, useSelector } from 'react-redux';
import axios from "axios";
import Navigation from "../HomePage/Navigation";
import {Button} from "react-bootstrap";
import {user} from "../../store/reducers/user.reducer";
import userService from "../../services/user.service";


function SubmitVerificationRequest() {
    const categories = ["Sports", "Influencer", "News","Brand","Business","Organization","Government"]
    const [selectedCategory, setSelectedCategory] = useState("")
    const [image, setImage] = useState("")

    const dispatch = useDispatch()
    const store = useSelector(state => state);

    const handleInputChange = (event) => {
        const target = event.target;
        setSelectedCategory(target.value)
    }

    function handleChangeImage(evt) {
        console.log("Uploading");
        var self = this;
        var reader = new FileReader();
        var file = evt.target.files[0];

        reader.onload = function(upload) {
            setImage(upload.target.result)
        };
        reader.readAsDataURL(file);
    }

    function submitVerificationRequest() {
        console.log(store.user.id)
        axios.post("http://localhost:8080/api/users/api/users/submit-verification-request", {
            userId : store.user.id,
            documentPhoto : image,
            category : selectedCategory
        }, {
            headers: userService.setupHeaders(store.user.jwt)
        }).then(res => {
            alert("Verification request submitted successfully!")
            history.push({
                pathname: '/view-my-verification-request'
            })
        }) .catch(err => {
            alert("Error while submitting verification request")
        })
    }

    return (
        <div style={{marginTop:'5%'}}>
            <Navigation/>
            <div style={{marginTop:'5%', marginLeft: '5%'}}>
                <h1>Submit verification request</h1>
                <br/>
                <div className="row" style={{marginTop: '1rem'}}>
                    <label  className="col-sm-2 col-form-label">Category</label>
                    <div className="col-sm-6 mb-2">
                        <select onChange={(e) => handleInputChange(e)} name={"category"} value={selectedCategory}>
                            <option disabled={true} value="">Select category</option>
                            {categories.map(category => {
                                return (
                                    <option value={category}>{category}</option>
                                )
                            })}
                        </select>
                    </div>
                    <div className="col-sm-4">
                    </div>
                </div>

                <div className="row" style={{marginTop: '1rem'}}>
                    <label  className="col-sm-2 col-form-label">Document photo</label>
                    <div className="col-sm-6 mb-2">
                        <input type="file" name="file"
                               className="upload-file"
                               id="file"
                               onChange={handleChangeImage}
                               formEncType="multipart/form-data"
                               required />
                    </div>
                    <div className="col-sm-4">
                    </div>
                </div>

                <br/>
                <Button type={"primary"} onClick={submitVerificationRequest}>Submit</Button>


            </div>
        </div>

    );
}

export default SubmitVerificationRequest;