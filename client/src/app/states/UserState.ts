export interface UserState {
    FirstName: string,
    LastName: string,
    UserName: string,
    Email?: string,
    Password: string
}

export const InitialUserState: UserState = {
    FirstName: "",
    LastName: "",
    UserName: "",
    Email: "",
    Password: "",
}