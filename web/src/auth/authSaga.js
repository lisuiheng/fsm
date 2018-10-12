import { all, takeEvery } from 'redux-saga/effects';
import { FETCH_ERROR } from 'react-admin';

export default function* serverSaga() {
    yield takeEvery(FETCH_ERROR, function*() {
        yield all([]);
    });
}
