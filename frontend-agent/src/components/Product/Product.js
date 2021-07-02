import React, { useEffect, useState } from "react";
import toastService from './../../services/toast.service';

import '../../style/Profile.css';
import Navigation from "../HomePage/Navigation";
import {useDispatch, useSelector} from "react-redux";
import { useParams } from 'react-router-dom'
import Spinner from "../../helpers/spinner";
import ProductPreviewGrid from "./ProductPreviewGrid";


const Product = () => {
    const {id} = useParams()

    return (
        <div>
            <Navigation/>
            <div className="profileGrid">
                aa
            </div>
        </div>
    );
}
export default Product;
