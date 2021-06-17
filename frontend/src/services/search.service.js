import axios from 'axios';

class SearchService {
    constructor() {
        this.apiClient = axios.create({
            baseURL: "http://localhost:8080/api/content"
        })
    }

    setupHeaders(jwt) {
        return {
            Accept: 'application/json',
            Authorization: 'Bearer ' + jwt,
        }
    }

    async searchByTag(data){
        const {text,jwt} = data
        const headers=this.setupHeaders(jwt)
        return this.apiClient.post('/posts-by-hashtag', {
            text
        }, {
            headers: headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
    }
    async searchByLocation(data){
        const {location,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/location',{
            location
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