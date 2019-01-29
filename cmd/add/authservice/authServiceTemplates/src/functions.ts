export const getError = error => {
	/* put error mapping here */
	return 500;
}
export const errorHandler = (err, req, res, next) => {
	console.log(err.toString());
	res.status(getError(err)).send(err);
}

export const generateRandomString = () => {
    return Math.random().toString(36).substring(2, 15);
}