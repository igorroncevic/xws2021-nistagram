import axios from 'axios';

class SearchService {
    constructor() {
        this.apiClient = axios.create({
            baseURL: process.env.REACT_APP_CONTENT_SERVICE + "/api/content"
        })
    }

    setupHeaders(jwt) {
        return {
            Accept: 'application/json',
            Authorization: 'Bearer ' + jwt,
        }
    }

    async searchByTag(data){
        const {tag,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/searchByLocation',{
           tag
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }
    async searchByLocation(data){
        const {tag,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/searchByLocation',{
            tag
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }


}const searchService = new SearchService()

export default searchService;