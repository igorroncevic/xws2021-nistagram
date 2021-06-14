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
}

const userService = new UserService()

export default userService;