import { TokenCredentials } from '@azure/ms-rest-js';

export default function ApplicationTokenCredentials(token: string) {
    return new TokenCredentials(token, 'Bearer');
}