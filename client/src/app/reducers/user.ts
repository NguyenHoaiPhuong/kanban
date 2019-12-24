import { UserAction } from "../actions/UserAction";
import { InitialUserState } from "../states/UserState";

// user object is set server side and is never updated client side but this empty reducer is still needed
const user = (prevState = InitialUserState, action: UserAction) => {
    switch (action.type) {
        default:
            return prevState;
    }
};
  
export default user;