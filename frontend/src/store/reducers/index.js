import { combineReducers } from 'redux';

import { user } from './user.reducer'
import {followers} from "./followers.reducer";

const rootReducer = combineReducers({
    user,followers
})

export default rootReducer;