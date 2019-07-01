export interface IWebStorage {
  getItem(key: string): string,
  hasItem(key: string): boolean,
  isSupported(): boolean,
  setItem(key: string, value: string): void
}

export class WebStorage implements IWebStorage {
  constructor(private storage: Storage = localStorage) {

  }

  getItem = (key: string): string => {
    return this.storage.getItem(key);
  }

  hasItem = (key: string): boolean => {
    return this.storage.getItem(key) !== null
  }

  isSupported = (): boolean => {
    return typeof (Storage) !== "undefined";
  }

  setItem = (key: string, value: string): void => {
    this.storage.setItem(key, value);
  }

  removeItem = (key: string): void => {
    this.storage.removeItem(key)
  }
}

export default new WebStorage();