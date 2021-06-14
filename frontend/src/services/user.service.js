import axios from 'axios';

class UserService {
    constructor(){
        this.apiClient = axios.create({
            baseURL: "http://localhost:8001/api/users"
        })
    }

     setupHeaders(jwt){
        return {
            Accept: 'application/json',
            Authorization: 'Bearer ' + jwt,
        }
    }

    async login(data){
        const { email, password } = data
        const response = this.apiClient.post('/login', {
            email,
            password
        }).then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }

    async checkIsApproved(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/checkIsApproved', {
            id
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }


    async approveAccount(data){
        const { id,oldPassword,newPassword,repeatedPassword,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/approveAccount', {
            password:{id,oldPassword,newPassword,repeatedPassword}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            console.error(err)
            return err
        })
        return response
    }

    async getUserByUsername(data){
        const { username,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/getUserByUsername/'+username,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }
}

const userService = new UserService()

export default userService;