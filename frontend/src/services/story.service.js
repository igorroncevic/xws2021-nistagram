import RootService from './root.service'

class StoryService extends RootService {
    constructor(){
        super("http://localhost:8002/stories")
    }

    async getHomepageStories(data){
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

const storyService = new StoryService()

export default storyService;