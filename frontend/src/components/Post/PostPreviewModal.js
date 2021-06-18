import React from 'react';
import { Modal } from 'react-bootstrap';
import Post from './Post';

const PostPreviewModal = (props) => {
    const { postUser, post, setShowModal } = props;

    return (
        <Modal onHide={() => setShowModal(false)}>
            <Post post={post} postUser={postUser} />
        </Modal>
    )
}

export default PostPreviewModal;