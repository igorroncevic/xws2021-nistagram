import { userConstants } from "../constants";

const initialState = {
    storyId : ""
}

export const storyForReport = (state = initialState, action) => {
    switch(action.type) {
        case userConstants.SET_STORY:
            return {
                ...state,
                storyId: action.data.storyId
            }
        default: 
            return state;
    }
}