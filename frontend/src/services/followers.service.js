import RootService from './root.service'

class FollowersService extends RootService {
    constructor() {
        super("http://localhost:8005/api/followers")
    }

    async getFollowers(data){
        const { userId,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/get_followers',{
            user: { UserId: userId}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getFollowing(data){
        const { userId,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/get_followings',{
            user: { UserId: userId}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async follow(data){
        const { userId,followerId,isApprovedRequest,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/create_connection',{
            userId,followerId,isApprovedRequest,
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async unfollow(data){
        const { userId,followerId,isApprovedRequest,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/delete_directed',{
            userId,followerId,isApprovedRequest,
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }


}const followersService = new FollowersService()

export default followersService;