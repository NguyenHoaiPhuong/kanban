import { UserAction } from "../actions/UserAction";
import { ActionType } from "../actions/ActionType";

const isGuest = (prevState = false, action: UserAction) => {
    switch (action.type) {
        case ActionType.ENTER_AS_GUEST:
            return true;
        default:
            return prevState;
    }
};
  
export default isGuest;