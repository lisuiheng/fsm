import 'babel-polyfill';
import React, { Component } from 'react';
import { Admin, Resource } from 'react-admin';

import './App.css';

import { authProvider } from './auth';
import sagas from './sagas';
import themeReducer from './themeReducer';
import Login from './Login';
import Menu from './Menu';
import defultMessages from './i18n/cn';
import customRoutes from './routes';

import {
    UserList,
    UserCreate,
    UserShow,
    UserEdit,
    UserResourceName,
    UserIcon,
} from './user';
import dataProviderFactory from './dataProvider';

const i18nProvider = locale => {
    // if (locale === 'cn') {
    //     return import('./i18n/cn').then(messages => messages.default);
    // }

    // Always fallback on english
    return defultMessages;
};

class App extends Component {
    state = { dataProvider: null };

    async UNSAFE_componentWillMount() {
        // this.restoreFetch = await fakeServerFactory(
        //     process.env.REACT_APP_DATA_PROVIDER
        // );

        const dataProvider = await dataProviderFactory('rest');

        this.setState({ dataProvider });
    }

    render() {
        const { dataProvider } = this.state;

        if (!dataProvider) {
            return (
                <div className="loader-container">
                    <div className="loader">Loading...</div>
                </div>
            );
        }

        return (
            <Admin
                title="FSM"
                dataProvider={dataProvider}
                customReducers={{ theme: themeReducer }}
                customSagas={sagas}
                customRoutes={customRoutes}
                authProvider={authProvider}
                // dashboard={Dashboard}
                loginPage={Login}
                // appLayout={Layout}
                menu={Menu}
                locale="cn"
                i18nProvider={i18nProvider}
            >
                {permissions => [
                    permissions('usersALL') ? (
                        <Resource
                            name={UserResourceName}
                            list={UserList}
                            create={
                                permissions('usersPOST') ? UserCreate : null
                            }
                            edit={permissions('usersPUT') ? UserEdit : null}
                            show={UserShow}
                            icon={UserIcon}
                        />
                    ) : null,
                ]}
            </Admin>
        );
    }
}

export default App;
