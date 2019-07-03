export interface IWebStorage {
  getItem(key: string): string,
  hasItem(key: string): boolean,
  isSupported(): boolean,
  setItem(key: string, value: string): void
}
export interface IStorageBackend {
  getItem(key: string): string,
  setItem(key: string, value: string): void
}

export class WebStorageImpl implements IWebStorage {
  private storage: IStorageBackend;
  constructor(storage: IStorageBackend = localStorage) {
    this.storage = storage;
  }

  getItem = (key: string): string => {
    return this.storage.getItem(key);
  }

  hasItem = (key: string): boolean => {
    return this.storage.getItem(key) !== null
  }

  isSupported = (): boolean => {
    return typeof (Storage) !== "undefined" && (this.storage.getItem != null);
  }

  setItem = (key: string, value: string): void => {
    this.storage.setItem(key, value);
  }
}

export default new WebStorageImpl();