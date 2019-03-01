export interface User {
  exp?: number;
  groups?: Array<any>;
  iat?: number;
  iss?: string;
  name: string;
  sub?: string;
}
