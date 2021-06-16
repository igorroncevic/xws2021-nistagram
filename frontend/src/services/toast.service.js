import { toast } from 'react-toastify';

const defaultOptions = {
    autoClose: 3000,
    hideProgressBar: true,
}

/** 
 * Types: ```success```, ```info```, ```warning```, ```error```
*/
const show = (type, message, options) => {
    let additionalOptions = {};
    if(options) {
        additionalOptions = { ...defaultOptions, ...options }
    }
    const toastOptions = (!options) ? defaultOptions : additionalOptions;
    toast[type](message, toastOptions);  
}

export default {
    show,
};