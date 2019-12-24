import { ActionType } from './ActionType'

export interface UserAction {
    type: ActionType
    payload?: any
}

