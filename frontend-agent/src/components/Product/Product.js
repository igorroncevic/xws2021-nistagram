import React, { useEffect, useState } from "react";
import toastService from './../../services/toast.service';

import '../../style/Profile.css';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import {useHistory, useParams} from 'react-router-dom'
import Spinner from "../../helpers/spinner";
import ProductPreviewGrid from "./ProductPreviewGrid";
import productService from "../../services/product.service";
import {Button} from "react-bootstrap";


const Product = () => {
    const {id} = useParams()

    const dispatch = useDispatch()
    const store = useSelector(state => state);
    const [product, setProduct] = useState({});
    const history = useHistory()

    useEffect(() => {
        getProductById(id);
    },[]);

    async function getProductById(id) {
        const response = await productService.getProductById({
            id : id,
            jwt: store.user.jwt,
        })

        if (response.status === 200) {
            setProduct(response.data)
            console.log(response.data)
        } else {
            console.log(response);
            toastService.show("error", "Could not retrieve products")
        }
    }

    async function deleteProduct() {
        if (window.confirm('Are you sure you want to delete this product?')) {
            const response = await productService.deleteProduct({
                id : id,
                jwt: store.user.jwt,
            })

            if (response.status === 200) {
                toastService.show("success", "Product successfully deleted")
                history.push({ pathname: '/profile/' + store.user.username })
            } else {
                console.log(response);
                toastService.show("error", "Could not retrieve products")
            }
        }
    }

    return (
        <div>
            <Navigation/>
            <div className="profileGrid">
                <div className="card">
                    {/*<img src={""} alt="product photo" style="width:100%"/>*/}
                        <h1>{product.name}</h1>

                    {product.agentId === store.user.id && <Button style={{width: '250px'}} variant={"outline-primary"}>Edit</Button>}
                    {product.agentId === store.user.id && <Button style={{width: '250px'}} variant={"outline-danger"} onClick={deleteProduct}>Delete</Button>}
                        <p className="price">${product.price}</p>
                        <p>Quantity : {product.quantity}</p>
                        <p>
                            <button>Add to Cart</button>
                        </p>


                </div>
            </div>
        </div>
    );
}
export default Product;
