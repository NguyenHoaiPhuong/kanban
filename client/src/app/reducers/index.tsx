import { combineReducers } from "redux";
// import cardsById from "./cardsById";
// import listsById from "./listsById";
// import boardsById from "./boardsById";
import user from "./user";
import isGuest from "./isGuest";
// import currentBoardId from "./currentBoardId";

export const reducers = combineReducers({
//   cardsById,
//   listsById,
//   boardsById,
// currentBoardId,
  user,
  isGuest
});