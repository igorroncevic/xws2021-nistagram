import axios from 'axios';

class UserService {
    constructor(){
        this.apiClient = axios.create({
            baseURL: "http://localhost:8080/api/users/api/users"
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
        }).then(err => {
            console.error(err)
            return {}
        })
        return response
    }
}

const userService = new UserService()

export default userService;