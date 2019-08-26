import * as axios from 'axios';

export const httpGet = async (url: string): Promise<axios.AxiosResponse<any>> => {
  try {
    return await axios.default.get(url);
  } catch(err) {
    const errorContext: string = '[ http.ts ]';
    console.error(errorContext, err);
  }
}
