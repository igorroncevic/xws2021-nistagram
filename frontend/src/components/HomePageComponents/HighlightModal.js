import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { Modal, Button, ListGroup } from 'react-bootstrap';
import { ReactComponent as Check } from './../../images/icons/check.svg'

import highlightsService from '../../services/highlights.service';
import toastService from '../../services/toast.service';

import "./../../style/HighlightModal.css"

const HighlightModal = (props) => {
    const { 
        story,
        highlights,
        setHighlights,
        showModal,
        setShowModal,
        storiesToRender,
        setStoriesToRender,
        removeFromRender // if highlight is rendered (not all stories), remove it from there 
    } = props;

    const [selectedHighlight, setSelectedHighlight] = useState({});
    const [savedInHighlights, setSavedInHighlights] = useState([]);
    const [newHighlightName, setNewHighlightName] = useState("");
    const [isNewHighlightButtonDisabled, setIsNewHighlightButtonDisabled] = useState(true);
    const [isButtonRemove, setIsButtonRemove] = useState(false);
    const [isSaveButtonDisabled, setIsSaveButtonDisabled] = useState(true);

    const store = useSelector(state => state);

    useEffect(() => {
        const tempSavedIn = [];
        highlights.forEach(highlight => {
            const contained = highlight.stories.some(highlightStory => highlightStory.id === story.id)
            if(contained) tempSavedIn.push(highlight.id) 
        })
        setSavedInHighlights(tempSavedIn)
    }, [story])

    useEffect(() => {
        newHighlightName && setIsNewHighlightButtonDisabled(false);
    }, [newHighlightName])

    
    useEffect(() => {
        setIsButtonRemove(savedInHighlights.some(id => selectedHighlight.id === id))
        selectedHighlight && setIsSaveButtonDisabled(false);
    }, [selectedHighlight])

    const handleNewHighlightName = (e) => {
        setNewHighlightName(e.target.value)
    }

    const saveToHighlight = async () => {
        const highlightRequest = {
            userId: store.user.id,
            highlightId: selectedHighlight.id,
            storyId: story.id,
            jwt: store.user.jwt,
        };

        const response = await highlightsService.saveStoryToHighlight(highlightRequest)

        if(response.status === 200) {
            toastService.show("success", "Successfully saved this story to " + selectedHighlight.name + ".");
            setSavedInHighlights([...savedInHighlights, selectedHighlight.id])
            const updateHighlight = highlights.filter(highlight => highlight.id === selectedHighlight.id)[0]
            updateHighlight.stories ? updateHighlight.stories = [...updateHighlight.stories, story] : updateHighlight.stories = [story];
            setHighlights([...highlights.filter(highlight => highlight.id !== updateHighlight.id), updateHighlight])
            setSelectedHighlight({});
            setShowModal(!showModal);
        }else{
            toastService.show("error", "Could not save this story to " + selectedHighlight.name + ".");
        }
    }

    const removeFromHighlight = async () => {
        const highlightRequest = {
            userId: store.user.id,
            highlightId: selectedHighlight.id,
            storyId: story.id,
            jwt: store.user.jwt,
        };
        const response = await highlightsService.removeStoryFromHighlight(highlightRequest)

        if(response.status === 200) {
            toastService.show("success", "Successfully removed this story from " + selectedHighlight.name + ".");
            setSavedInHighlights([...savedInHighlights.filter(highlight => highlight.id !== selectedHighlight.id)])
            const updateHighlight = highlights.filter(highlight => highlight.id === selectedHighlight.id)[0]
            updateHighlight.stories = updateHighlight.stories.filter(highlightStory => highlightStory.id !== story.id)
            setHighlights([...highlights.filter(highlight => highlight.id !== updateHighlight.id), updateHighlight])
            setSelectedHighlight({});
            removeFromRender && setStoriesToRender(...storiesToRender.filter(storyRender => storyRender.id !== story.id))
            setShowModal(!showModal);
        }else{
            toastService.show("error", "Could not removed this story from " + selectedHighlight.name + ".");
        }
    }

    const createNewHighlight = async () => {
        const highlightRequest = {
            name: newHighlightName,
            userId: store.user.id,
            jwt: store.user.jwt
        }
        const response = await highlightsService.createHighlight(highlightRequest)

        if(response.status === 200) {
            setNewHighlightName("");
            setIsNewHighlightButtonDisabled(true);
            setHighlights([...highlights, {
                id: response.data.id,
                name: highlightRequest.name,
                userId: highlightRequest.userId,
                stories: [],
            }])
            toastService.show("success", `You have created a new highlight called ${highlightRequest.name}!`)
        }else{
            toastService.show("error", `Could not create a new highlight`)
        }
    }

    return(
        <Modal className="saveModal" size="lg" show={showModal} onHide={() => setShowModal(!showModal)} animation={true} centered>
            <Modal.Header closeButton>
                <Modal.Title>Choose a highlight</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <ListGroup variant="flush" className="collectionsList">
                    {highlights.map((highlight, key) => {
                        return (
                            <ListGroup.Item 
                                action    
                                active={selectedHighlight.id === highlight.id} 
                                onClick={() => setSelectedHighlight(highlight)}
                                className="highlightCard">
                                    { savedInHighlights.some(id => highlight.id === id) ? 
                                        <div className="savedInHighlight">{highlight.name}<Check className="savedCheck" /> </div> : 
                                        highlight.name 
                                    } 
                            </ListGroup.Item>
                        )
                    })}
                </ListGroup>
                <div className="newHighlight">
                    <p className="newHighlightLabel">New highlight: </p>
                    <input 
                        className="newHighlightInput" 
                        placeholder="Enter highlight name..." 
                        value={newHighlightName} 
                        onChange={handleNewHighlightName}
                    />
                    
                    <Button 
                        className="newHighlightButton" 
                        variant="outline-primary" 
                        disabled={isNewHighlightButtonDisabled} 
                        onClick={createNewHighlight}>Create</Button>
                </div>
            </Modal.Body>
            <Modal.Footer style={{'background':'#E0E0E0'}}>
                { !isButtonRemove ?
                    <Button variant="primary" disabled={isSaveButtonDisabled} onClick={saveToHighlight}>Save</Button> : 
                    <Button variant="primary" disabled={isSaveButtonDisabled} onClick={removeFromHighlight}>Remove</Button>
                }
            </Modal.Footer>
        </Modal>
    )
}

export default HighlightModal;