import RootService from './root.service'

class PostService extends RootService {
    constructor(){
        super("http://localhost:8002/posts")
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
}

const postService = new PostService()

export default postService;