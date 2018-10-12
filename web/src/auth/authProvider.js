import {
    AUTH_LOGIN,
    AUTH_GET_PERMISSIONS,
    AUTH_LOGOUT,
    AUTH_CHECK,
    AUTH_ERROR,
} from 'react-admin';
import decodeJwt from 'jwt-decode';

export const AUTH_KEY = 'AUTHORIZATION';
export const ROLE_KEY = 'ROLE';
const NO_AUTH = 'NO_AUTH';

export default (type, params) => {
    if (type === AUTH_LOGIN) {
        const { username, password } = params;
        const request = new Request('/api/login', {
            method: 'POST',
            body: JSON.stringify({ username, password }),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        });
        return fetch(request).then(response => {
            if (response.status < 200 || response.status >= 300) {
                throw new Error(response.statusText);
            }
            const token = response.headers.get(AUTH_KEY);

            if (token !== NO_AUTH) {
                localStorage.setItem(AUTH_KEY, token);
                localStorage.setItem(ROLE_KEY, response.headers.get(ROLE_KEY));
            } else {
                localStorage.setItem(AUTH_KEY, NO_AUTH);
            }
        });
    }
    if (type === AUTH_GET_PERMISSIONS) {
        const auth = localStorage.getItem(AUTH_KEY);
        if (auth === NO_AUTH) {
            return Promise.resolve(() => {
                return true;
            });
        }
        const roles = localStorage.getItem(ROLE_KEY);

        const checkPermissions = (...authParams) => {
            return authParams.every(a => roles.indexOf(a) !== -1);
        };
        return Promise.resolve(checkPermissions);
    }
    if (type === AUTH_LOGOUT) {
        localStorage.removeItem(AUTH_KEY);
        localStorage.removeItem(ROLE_KEY);
        return Promise.resolve();
    }
    if (type === AUTH_ERROR) {
        const { status } = params;
        return status === 401 || status === 403
            ? Promise.reject()
            : Promise.resolve();
    }
    if (type === AUTH_CHECK) {
        return localStorage.getItem(AUTH_KEY)
            ? Promise.resolve()
            : Promise.reject();
    }
    return Promise.reject('Unkown method');
};
