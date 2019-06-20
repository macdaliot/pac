import * as React from 'react';
import { shallow } from 'enzyme';
import {
    LoginCallbackComponent,
    mapDispatchToProps
} from './LoginCallback.component';
import { IUser } from '@pyramid-systems/domain';
import { IWebStorage, WebStorage } from '@app/config';

jest.mock('@app/config');

let StorageMock: jest.Mocked<IWebStorage> = WebStorage as any;

const tokenFromJwtIo =
    'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c';
const sampleUser: IUser = {
    exp: 1234567890,
    groups: [],
    iat: 1516239022,
    iss: 'urn:pacAuth',
    name: 'John Doe',
    sub: 'sample@pyramidsystems.com'
};
describe('Login Callback', () => {
    beforeEach(() => {
        jest.resetAllMocks();
    });

    it('should render if local storage is supported', () => {
        StorageMock.isSupported.mockReturnValue(true);

        const location = {
            search: `/something/something-else?${tokenFromJwtIo}`
        };
        shallow(
            <LoginCallbackComponent
                setToken={() => ({
                    type: 'JWT_RECEIVED',
                    payload: {
                        token: tokenFromJwtIo,
                        user: sampleUser
                    }
                })}
                location={location}
            />
        );

        expect(StorageMock.setItem).toBeCalled();
    });

    it('should not set local storage if it is NOT supported', () => {
        const location = {
            search: `/something/something-else?${tokenFromJwtIo}`
        };
        shallow(
            <LoginCallbackComponent
                setToken={() => ({
                    type: 'JWT_RECEIVED',
                    payload: {
                        token: tokenFromJwtIo,
                        user: sampleUser
                    }
                })}
                location={location}
            />
        );
        expect(StorageMock.setItem).toBeCalledTimes(0);
    });
    it('should map dispatch', () => {
        expect.assertions(2);
        const result = mapDispatchToProps(action => {
            return expect(action.payload).toEqual({
                token: tokenFromJwtIo,
                user: sampleUser
            });
        });
        expect(result.setToken).toBeTruthy();
        result.setToken(tokenFromJwtIo, sampleUser);
    });
});
