export interface IUser {
  name: string;
  sub: string;
  groups: Array<string>;
  exp?: number;
  iat?: number;
  iss?: string;
}

export class User implements IUser {
  name: string;
  sub: string;
  groups: string[];
  exp?: number;
  iat?: number;
  iss?: string;
}
