import RootService from './root.service'
import { monitoringService as getMonitoringService } from './../backendPaths';

class MonitoringService extends RootService {
    constructor() {
        super(getMonitoringService())
    }

    // Unimplemented: Ping a server every minute so that we know that he is online. Not needed at the moment.
    activityPing(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt);

        const response = this.apiClient.post("/activity/ping", { userId }, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getUsersRecentActivity(data){
        const { userId, jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = await this.apiClient.get(`/user/activity/${userId}`, { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }

    async getPerformanceMessages(data){
        const { jwt } = data;
        const headers = this.setupHeaders(jwt)

        const response = await this.apiClient.get('/performance', { headers })
            .then(res => {
                return res
            }).catch(err => {
                console.error(err)
                return err
            })
        return response
    }
}

const monitoringService = new MonitoringService()

export default monitoringService;

