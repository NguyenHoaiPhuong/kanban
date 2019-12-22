export const getHeader = (jsonContentType: boolean = true) => {
    const headers = {
        'Accept': 'application/json',
        'Authorization': ""
    };
    if (jsonContentType) {
        headers['Content-Type'] = 'application/json'
    } else {
        headers['Content-Type'] = 'multipart/form-data; boundary="--"'
    }
    if (window.localStorage.jwt) {
        headers.Authorization =`JWT ${window.localStorage.jwt}`
    }
}