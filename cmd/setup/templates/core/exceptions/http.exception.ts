export class HttpException {
  status: number;
  message: any;
  constructor(status: number, message: any) {
    this.status = status;
    this.message = message;
  }
}
