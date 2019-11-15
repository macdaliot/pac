export class User {
	_id: string;
	firstName: string;
	lastName: string;
	phone: string;
	email: string;
	birthdate: Date;
	notes: string;
	title: string;
	role: string;

	constructor(data: any = null) {
		Object.assign(this, data);
		this.birthdate = new Date(data.birthdate);
	}

}
