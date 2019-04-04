export interface IUser {
  name: string;
  sub: string;
  groups: Array<string>;
}

export class User implements IUser {
  name: string;
  sub: string;
  groups: string[];
}
