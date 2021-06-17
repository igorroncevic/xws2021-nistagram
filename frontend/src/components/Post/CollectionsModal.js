import React, { useEffect, useState } from 'react';
import { Button, Modal, ListGroup } from "react-bootstrap";
import { useSelector } from 'react-redux';
import { ReactComponent as Check } from './../../images/icons/check.svg'
import collectionsService from './../../services/collections.service'
import favoritesService from './../../services/favorites.service'
import toastService from './../../services/toast.service'
import './../../style/CollectionsModal.css'

const CollectionsModal = (props) => {
    const { 
        showModal, 
        setShowModal, 
        postId, 
        isSaved, 
        setIsSaved, 
        collections, 
        setCollections, 
        savedInCollections, 
        setSavedInCollections 
    } = props;

    const [newCollectionName, setNewCollectionName] = useState("");
    const [isNewCollectionButtonDisabled, setIsNewCollectionButtonDisabled] = useState(true);
    const [selectedCollection, setSelectedCollection] = useState({});   // Collection to save a post in 
    const [isButtonRemove, setIsButtonRemove] = useState(false)  // Checking if button should be save or remove

    const store = useSelector(state => state);

    useEffect(()=>{
        setIsNewCollectionButtonDisabled(newCollectionName === "")
    }, [newCollectionName])

    useEffect(() => {
        setIsButtonRemove(savedInCollections.some(id => selectedCollection.id === id))
    }, [selectedCollection])

    const saveToCollection = async () => {
        const newFavorite = {
            userId: selectedCollection.userId,
            postId: postId,
            collectionId: selectedCollection.id == "1" ? null : selectedCollection.id,
            jwt: store.user.jwt,
        }

        const response = await favoritesService.createFavorite(newFavorite);
        
        if(response.status === 200){
            toastService.show("success", "Successfully saved this post to " + selectedCollection.name);
            setSelectedCollection({});
            setShowModal(!showModal);
        }else{
            toastService.show("error", "Could not save this post to " + selectedCollection.name);
        }
    }

    const removeFromCollection = async () => {
        const removingFavorite = {
            userId: selectedCollection.userId,
            postId: postId,
            collectionId: selectedCollection.id == "1" ? null : selectedCollection.id,
            jwt: store.user.jwt,
        }
        
        const response = await favoritesService.removeFavorite(removingFavorite);
        
        if(response.status === 200){
            toastService.show("success", "Successfully removed this post from " + selectedCollection.name);
            setSelectedCollection({});
            setShowModal(!showModal);
        }else{
            toastService.show("error", "Could not remove this post from " + selectedCollection.name);
        }
    }

    const createNewCollection = async () => {
        const collectionRequest = {
            name: newCollectionName,
            userId: store.user.id,
            jwt: store.user.jwt
        }
        const response = await collectionsService.createCollection(collectionRequest)

        if(response.status === 200) {
            setNewCollectionName("");
            setCollections([...collections, {
                id: response.data.id,
                name: collectionRequest.name,
                userId: collectionRequest.userId,
            }])
        }
    }

    const handleNewCollectionName = (e) => {
        setNewCollectionName(e.target.value)
    }

    return (
        <Modal className="saveModal" size="lg" show={showModal} onHide={setShowModal} animation={true} centered>
            <Modal.Header closeButton>
                <Modal.Title>Choose a collection</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <ListGroup variant="flush" className="collectionsList">
                    {collections.map((collection, key) => {
                        return (
                            <ListGroup.Item 
                                action    
                                active={selectedCollection.id === collection.id} 
                                onClick={() => setSelectedCollection(collection)}
                                className="collectionCard">
                                    { savedInCollections.some(id => collection.id === id) ? 
                                        <div className="savedInCollection">{collection.name} <Check className="savedCheck" /> </div> : 
                                        collection.name 
                                    } 
                            </ListGroup.Item>
                        )
                    })}
                </ListGroup>
                <div className="newCollection">
                    <p className="newCollectionLabel">New collection: </p>
                    <input 
                        className="newCollectionInput" 
                        placeholder="Enter collection name..." 
                        value={newCollectionName} 
                        onChange={handleNewCollectionName}
                    />
                    
                    <Button 
                        className="newCollectionButton" 
                        variant="outline-primary" 
                        disabled={isNewCollectionButtonDisabled} 
                        onClick={createNewCollection}>Create</Button>
                </div>
            </Modal.Body>
            <Modal.Footer style={{'background':'#E0E0E0'}}>
                { !isButtonRemove ?
                    <Button variant="primary" onClick={saveToCollection}>Save</Button> : 
                    <Button variant="primary" onClick={removeFromCollection}>Remove</Button>
                }
            </Modal.Footer>
        </Modal>
    )
}

export default CollectionsModal;