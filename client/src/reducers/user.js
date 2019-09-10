const user = {
    id: -1,
    first: "",
    last: "",
    username: "",
    thumbnail: ""
}

export const getUser = () => (user);

export const setUser = (u) => {
    return Object.assign(user, u)
};