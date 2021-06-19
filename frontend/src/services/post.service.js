import RootService from './root.service'
import { contentService } from './../backendPaths';


class PostService extends RootService {
    constructor(){
        super(contentService() + "/posts")
    }

    async getHomepagePosts(data){
        const { jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.get('', { headers })
        .then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }

    async getPostsForUser(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = this.apiClient.get(`/user/${userId}`, { headers })
        .then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }
}

const postService = new PostService()

export default postService;