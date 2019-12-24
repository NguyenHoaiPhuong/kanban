import { UserState, InitialUserState } from './UserState'

export interface AppState {
    user: UserState,
    isGuest: boolean,
    isUser: boolean
}

export const InitialAppState: AppState = {
    user: InitialUserState,
    isGuest: false,
    isUser: false
}