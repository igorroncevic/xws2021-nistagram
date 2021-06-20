import { combineReducers } from 'redux';

import { user } from './user.reducer'
import {followers} from "./followers.reducer";
import {storyForReport} from "./storyForReport.reducer";

const rootReducer = combineReducers({
    user,followers, storyForReport
})

export default rootReducer;