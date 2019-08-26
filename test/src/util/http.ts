import * as axios from 'axios';
import { encodeBase64 } from './encoding';
import { logAndQuit } from './errors';

export const httpGet = async (url: string): Promise<axios.AxiosResponse<any>> => {
  try {
    return await axios.default.get(url);
  } catch (error) {
    logAndQuit(error);
  }
}

export const httpGetWithAuth = async (url: string, authHeader: string): Promise<axios.AxiosResponse<any>> => {
  try {
    return await axios.default.get(url, {
      headers: {
        'Authorization': authHeader
      }
    });
  } catch (error) {
    logAndQuit(error);
  }
}

export const httpGetWithBasicAuth = async (url: string, username: string, password: string): Promise<axios.AxiosResponse<any>> => {
  try {
    const encodedCredentials: string = encodeBase64(`${username}:${password}`);
    const basicAuth: string = `Basic ${encodedCredentials}`;
    return await axios.default.get(url, {
      headers: {
        'Authorization': basicAuth
      }
    });
  } catch (error) {
    logAndQuit(error);
  }
}
