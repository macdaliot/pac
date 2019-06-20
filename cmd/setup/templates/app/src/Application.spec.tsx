import * as React from 'react';
import { shallow } from 'enzyme';
import { ApplicationComponent } from './Application';
import { ApplicationStore } from './redux/Store';
import { WebStorage, IWebStorage } from './config';
import { Store } from 'redux';
jest.mock('./config');
jest.mock('./redux/Store');

let StorageMock: jest.Mocked<IWebStorage> = WebStorage as any;
let StoreMock: jest.Mocked<Store> = ApplicationStore as any;

describe('Application component', () => {
  beforeEach(() => {
    jest.resetAllMocks();
  });

  describe('Token is stored in local storage', () => {
    it('should dispatch an action to the store', () => {
      StorageMock.isSupported.mockReturnValue(true);
      StorageMock.hasItem.mockReturnValue(true);
      StorageMock.getItem.mockReturnValue(
          'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c'
      );
      shallow(<ApplicationComponent />);
      expect(StoreMock.dispatch).toHaveBeenCalledTimes(1);
    });
  });

  describe('Token not stored in local storage', () => {
    it('should not dispatch an action to the store', () => {
      shallow(<ApplicationComponent />);
      expect(StoreMock.dispatch).toHaveBeenCalledTimes(0);
    });
  });
});
