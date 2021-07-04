import RootService from './root.service'
import { contentService } from './../backendPaths';

class ComplaintService extends RootService {
    constructor(){
        super(contentService() + "/complaint")
    }


    async createComplaint(data){
        const { id, category, postId, status, isPost, userId, jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post('/create', {
            id, category, postId, status, isPost, userId
        }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getAllContentComplaints(data){
        const { jwt} = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get('/get' , { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const complaintService = new ComplaintService();

export default complaintService;