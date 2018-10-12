import { fetchUtils } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import { AUTH_KEY } from '../auth';

export const BASE_PATH = '/api';

export const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }
    // add your own headers here
    options.headers.set(AUTH_KEY, localStorage.getItem(AUTH_KEY));
    return fetchUtils.fetchJson(url, options);
};

const restProvider = simpleRestProvider(BASE_PATH, httpClient);
export default (type, resource, params) =>
    new Promise(resolve =>
        setTimeout(() => resolve(restProvider(type, resource, params)), 500)
    );
