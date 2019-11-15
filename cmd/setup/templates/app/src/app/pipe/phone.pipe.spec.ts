import { PhonePipe } from './phone.pipe';

describe('PhonePipe', () => {
  it('create an instance', () => {
    const pipe = new PhonePipe();
    expect(pipe).toBeTruthy();
  });

  it('valid phone number', () => {
    const pipe = new PhonePipe();

    const case1 = '1234567890';
    expect(pipe.transform(case1)).toBe('(123) 456-7890');
  });

  it('invalid phone number', () => {
    const pipe = new PhonePipe();

    const case1 = '123456789';
    expect(pipe.transform(case1)).toBe('123456789');
  });

  it('empty phone number', () => {
    const pipe = new PhonePipe();

    const case1 = '';
    expect(pipe.transform(case1)).toBe('');
  });

});
