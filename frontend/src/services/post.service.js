import RootService from './root.service'

class PostService extends RootService {
    constructor(){
        super(process.env.REACT_APP_CONTENT_SERVICE + "/posts")
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