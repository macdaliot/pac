export const decodeBase64 = (data: string): string => {
  return Buffer.from(data, 'base64').toString();
}

export const encodeBase64 = (data: string): string => {
  return Buffer.from(data).toString('base64');
}
