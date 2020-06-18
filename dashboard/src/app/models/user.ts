export class User {
  constructor(
    public username: string,
    public masterpassword: string,
    public twofa: boolean,
  ) {
  }
}
