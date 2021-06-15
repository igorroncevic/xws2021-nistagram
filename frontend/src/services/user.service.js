import RootService from './root.service'

class UserService extends RootService {
    constructor(){
        super("http://localhost:8001/api/users")
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
        const { id, jwt } = data
        const headers = this.setupHeaders(jwt)
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
        const { id, oldPassword, newPassword, repeatedPassword, jwt } = data
        const headers = this.setupHeaders(jwt)
        const response = this.apiClient.post('/approveAccount', {
            password: { id, oldPassword, newPassword, repeatedPassword }
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
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

    async getUserById(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/'+id,{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async editProfile(data){
        const { id,firstName,lastName,email,phoneNumber,username,profilePhoto,sex,website,biography,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/update_profile',{
            user:{id,firstName,lastName,email,phoneNumber,username,profilePhoto,sex,website,biography}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async changePassword(data){
        console.log(data)
        const {id,oldPassword,newPassword,repeatedPassword,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.post('/update_password',{
            password:{id,oldPassword,newPassword,repeatedPassword}
        },{
            headers:headers
        }).then(res => {
            return res
        }).catch(err => {
            return err
        })
        return response
    }

    async getUsernameById(data){
        const { id,jwt} = data
        const headers=this.setupHeaders(jwt)
        const response = this.apiClient.get('/username/'+id,{
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