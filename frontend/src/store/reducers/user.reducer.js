import { userConstants } from "../constants";

const initialState = {
    user: {},
    jwt: "",
    id: "",
    role: "",
    isSSO: false,
}

export const user = (state = initialState, action) => {
    switch(action.type) {
        case userConstants.REGISTER_REQUEST:
            return {
                ...state,
                jwt: action.data.jwt,
                id: action.data.id,
                role: action.data.role,
                isSSO: action.data.isSSO
            }
        case userConstants.LOGIN_REQUEST:
            return {
                ...state,
                jwt: action.data.jwt,
                id: action.data.id,
                role: action.data.role,
                isSSO: action.data.isSSO
        }
        default: 
            return state;
    }
}