import '../../style/Story.css';
import '../../style/ProfileIcon.css';
import Stories from "react-insta-stories";
import React, {useState} from "react";
import {Modal} from "react-bootstrap";


function Story(props) {
    const { story } = props;
    const [showModal, setModal] = useState(false);

    return (
        <div>
            <div className="story">
                <div className={true ? "storyBorder" : ""}>
                    <img className={`profileIcon big`} src=""  alt="profile" onClick={setModal(!showModal)}/>
                </div>
                <span className="accountName">Ime</span>
            </div>
            
            <Modal show={showModal} onHide={setModal(!showModal)}>
                <Modal.Header closeButton >
                    <Modal.Title>Nistagram</Modal.Title>
                </Modal.Header>
                <Modal.Body >
                    <Stories stories={story.stories} defaultInterval={1500} width={432} height={768}/>
                </Modal.Body>
                <Modal.Footer >

                </Modal.Footer>
            </Modal>
        </div>
    );
}

export default Story;