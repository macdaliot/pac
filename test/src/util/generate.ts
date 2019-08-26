export const randomString = (stringLength: number): string => {
  const validCharacters: string = 'abcdefghijklmnopqrstuvwxyz';
  let randomString: string = '';
  for (let i = 0; i < stringLength; i++) {
    let randomCharacterIndex: number = Math.random() * (validCharacters.length);
    randomString += validCharacters.charAt(randomCharacterIndex);
  }
  return randomString;
}
