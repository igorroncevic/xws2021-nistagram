import { userConstants } from "../constants";

const initialState = {
    userId: "",
    followerId: "",
}

export const followers = (state = initialState, action) => {
    switch(action.type) {
        case userConstants.FOLLOW_REQUEST:
            return {
                ...state,
                userId: action.data.userId,
                followerId: action.data.followerId,
            }
        default:
            return state;
    }
}