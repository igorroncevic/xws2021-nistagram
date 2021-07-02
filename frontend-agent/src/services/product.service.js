import RootService from './root.service'
import {agentService} from '../backendPaths';
import axios from "axios";

class ProductService extends RootService {
    constructor(){
        super(agentService() + "/api/agent/product")
    }


    async createProduct(data){
        const { id, name, price, isActive, quantity, photo, jwt, agentId} = data
        const response = axios.post('http://localhost:8080/api/agent/product/create-product',{
            id,name, price, isActive, quantity, photo, agentId
        }, {
            headers : this.setupHeaders(jwt)
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getProductsByAgent(data) {
        const { id, jwt} = data
        const response = axios.post('http://localhost:8080/api/agent/product/get-by-agent',{
            id
        }, {
            headers : this.setupHeaders(jwt)
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getAllProducts(data) {
        const { jwt} = data
        const response = axios.get('http://localhost:8080/api/agent/product/get-all', {
            headers : this.setupHeaders(jwt)
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }
}

const productService = new ProductService()

export default productService;