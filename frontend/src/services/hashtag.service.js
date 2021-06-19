import RootService from './root.service'
import { contentService } from './../backendPaths';

class HashtagService extends RootService {
    constructor(){
        super(contentService())
    }

    async getAllHashtags(data){
        const { id, category, postId, status, isPost, userId, jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get('/hashtag/get-all', { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const hashtagService = new HashtagService();

export default hashtagService;