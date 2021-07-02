import React, { useState, useEffect } from 'react';
import { Modal } from 'react-bootstrap';

import "./../../style/PostPreviewGrid.css"
import ProductPreviewThumbnail from "./ProductPreviewThumbnail";
import ProductPreviewModal from "./ProductPreviewModal";
import {useHistory} from "react-router-dom";

const ProductPreviewGrid = (props) => {
    const { posts, shouldReload } = props;

    const [selectedPost, setSelectedPost] = useState({});
    const [showModal, setShowModal] = useState(false);
    const history = useHistory()


    const openPost = (post) => {
        history.push({ pathname: '/product/' + post.id })
    }

    return (
        ( posts && 
        <div class="postPreviewGrid__Wrapper">
            { posts.map(post => <ProductPreviewThumbnail post={post} openPost={openPost} /> ) }

            {/*{ showModal && */}
            {/*<ProductPreviewModal */}
            {/*    post={selectedPost} */}
            {/*    postUser={{ id: selectedPost.agentId }} */}
            {/*    shouldReload={shouldReload}*/}
            {/*    showModal={showModal}*/}
            {/*    setShowModal={setShowModal}*/}
            {/*/> */}
            {/*}*/}
        </div>
        )
    )
}

export default ProductPreviewGrid;