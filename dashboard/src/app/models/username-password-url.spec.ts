import { UsernamePasswordUrl } from './username-password-url';

describe('UsernamePasswordUrl', () => {
  it('should create an instance', () => {
    expect(new UsernamePasswordUrl('', '', '', '', '')).toBeTruthy();
  });
});
