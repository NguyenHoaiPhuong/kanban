import {combineReducers} from 'redux';
import {getUser, setUser} from './user'

const AllReducers = combineReducers({
    user: getUser,
    setUser: setUser
})

export default AllReducers;