import { Buffer } from 'buffer';
import { IUser } from '@pyramid-systems/domain';

export const getUserFromToken = (token: string): IUser => {
    const payload = token.split('.')[1]; // lop off header and signature
    const json = Buffer.from(payload, 'base64').toString('ascii');
    return JSON.parse(json);
}