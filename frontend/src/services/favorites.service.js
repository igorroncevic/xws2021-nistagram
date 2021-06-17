import RootService from './root.service'

class FavoritesService extends RootService {
    constructor(){
        super("http://localhost:8002/favorites")
    }

    async createFavorite(data){
        const { userId, postId, collectionId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post("/create", { userId, postId, collectionId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async removeFavorite(data){
        const { userId, postId, collectionId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post("/remove", { userId, postId, collectionId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
    
    async getUserFavorites(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.get(`/${userId}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const collectionService = new FavoritesService();

export default collectionService;