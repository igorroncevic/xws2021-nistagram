import React, { useState, useEffect } from 'react';
import { Modal } from 'react-bootstrap';

import PostPreviewThumbnail from './../Post/PostPreviewThumbnail';
import PostPreviewModal from './../Post/PostPreviewModal';

import "./../../style/PostPreviewGrid.css"

const PostPreviewGrid = (props) => {
    const { posts } = props;

    const [selectedPost, setSelectedPost] = useState({});
    const [showModal, setShowModal] = useState(false);

    const openPost = (post) => {
        setSelectedPost(post);
        setShowModal(true);
    }

    return (
        ( posts && 
        <div class="postPreviewGrid__Wrapper">
            { posts.map(post => <PostPreviewThumbnail post={post} openPost={openPost} /> ) }    
            { showModal && 
            <PostPreviewModal 
                post={selectedPost} 
                postUser={{ id: selectedPost.userId }} 
                setShowModal={setShowModal}
            /> }
        </div>
        )
    )
}

export default PostPreviewGrid;