import RootService from './root.service'

class LikeService extends RootService {
    constructor(){
        super("http://localhost:8002/likes")
    }


    async addLike(data){
        const { userId, postId, isLike, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('', { userId, postId, isLike }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const likeService = new LikeService();

export default likeService;