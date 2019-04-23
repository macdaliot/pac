export interface WebStorage {
  getData(): WebStorageData,
  getItem(key: string): string,
  hasItem(key: string): boolean,
  isSupported(): boolean,
  setItem(key: string, value: string): void
}

export interface WebStorageData {
  [index: string]: string
}

const getData = (): WebStorageData => {
  let data: WebStorageData = {}
  for (let i = 0; i < localStorage.length; ++i) {
    let key = localStorage.key(i);
    let value = localStorage.getItem(key);
    data[key] = value;
  }
  return data;
}

const getItem = (key: string): string => {
  return localStorage.getItem(key);
}

const hasItem = (key: string): boolean => {
  return localStorage.getItem(key) !== null
}

const isSupported = (): boolean => {
  return typeof(Storage) !== "undefined";
}

const setItem= (key: string, value: string): void => {
  localStorage.setItem(key, value);
}

export const webStorage: WebStorage = {
  getData: getData,
  getItem: getItem,
  hasItem: hasItem,
  isSupported: isSupported,
  setItem: setItem
}
