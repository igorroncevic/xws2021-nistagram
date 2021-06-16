import axios from 'axios';

class FollowersService {
    constructor() {
        this.apiClient = axios.create({
            baseURL: "http://localhost:8005/api/followers"
        })
    }

    setupHeaders(jwt) {
        return {
            Accept: 'application/json',
            Authorization: 'Bearer ' + jwt,
        }
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

    async createConnection(data){
        const { userId,followerId,isApprovedRequest,isCloseFriends,isNotificationEnabled,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/create_connection',{
          follower:{  userId,followerId,isApprovedRequest,isCloseFriends,isNotificationEnabled}
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
             userId,followerId,isApprovedRequest
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }
    async updateUserConnection(data){
        const { userId,followerId,isApprovedRequest,isCloseFriends,isNotificationEnabled,jwt} = data
        const headers=this.setupHeaders(jwt)

        const response = this.apiClient.post('/update_follower',{
             userId,followerId,isApprovedRequest,isCloseFriends,isNotificationEnabled
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }


    async getFollowersConnection(data){
        const { userId,followerId,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/connection',{
             userId,followerId
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