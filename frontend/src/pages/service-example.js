import axios from "axios";
import store from './../store/index'
import { getBackendPath } from './backendPath'

class CounselingService {
    constructor() {
        this.apiClient = axios.create({
            baseURL: getBackendPath() + "/api/counseling"
        });
    }

    async getAllAvailablePharmacies(data) {
        let headers = this.setupHeaders()
        let pharmacies = await this.apiClient
            .post("/available", data, {
                headers
            })
            .then(response => {
                return response;
            })
            .catch(err => {
                return err.response;
            });

        return pharmacies;
    }

    async getAllAvailablePharmacistsInPharmacy(id, data) {
        let headers = this.setupHeaders()
        let counselings = await this.apiClient
            .post(`/available/${id}`, data, {
                headers
            })
            .then(response => {
                return response;
            })
            .catch(err => {
                return err.response;
            });

        return counselings;
    }

    async scheduleCounseling(data) {
        let headers = this.setupHeaders()
        let success = await this.apiClient
            .post(`/schedule`, data, {
                headers
            })
            .then(response => {
                return response;
            })
            .catch(err => {
                return err.response;
            });

        return success;
    }

    async cancelCounseling(data) {
        let headers = this.setupHeaders()
        let success = this.apiClient
            .put("/cancel", data, {
                headers
            })
            .then(response => {
                return response
            })
            .catch(err => {
                return err.response
            });
        return success;
    }

    async getAllPatientsCounselings(patientId) {
        let headers = this.setupHeaders()
        let checkups = await this.apiClient
            .get(`/patient/${patientId}`, {
                headers
            })
            .then(response => {
                return response.data;
            })
            .catch(err => {
                console.log(err);
                return [];
            });

        return checkups;
    }

    async getAllPatientsPastCounselingsPaginated(data) {
        let headers = this.setupHeaders()
        let checkups = await this.apiClient
            .post(`/past`, data, {
                headers
            })
            .then(response => {
                return response;
            })
            .catch(err => {
                console.log(err);
                return [];
            });

        return checkups;
    }

    async getAllPatientsUpcomingCounselingsPaginated(data) {
        let headers = this.setupHeaders()
        let checkups = await this.apiClient
            .post(`/upcoming`, data, {
                headers
            })
            .then(response => {
                return response;
            })
            .catch(err => {
                console.log(err);
                return [];
            });

        return checkups;
    }

    setupHeaders() {
        const jwt = store.getters.getJwt;
        let headers = {
            Accept: "application/json",
            Authorization: "Bearer " + jwt,
        };
        return headers;
    }
}

const counselingService = new CounselingService();

export default counselingService;