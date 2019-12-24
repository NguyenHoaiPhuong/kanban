import { UserState, InitialUserState } from './UserState'

export interface AppState {
    user: UserState
}

export const InitialAppState: AppState = {
    user: InitialUserState
}