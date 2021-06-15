import RootService from './root.service'

class PrivacyService extends RootService {
    constructor() {
        super("http://localhost:8001/api/privacy")
    }

    async getUserPrivacy(data){
        const { userId, jwt } = data
        const headers = this.setupHeaders(jwt)
        const response = this.apiClient.post('/isProfilePublic', {
            userId
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

}const privacyService = new PrivacyService()

export default privacyService;