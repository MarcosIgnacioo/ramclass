class UserData {
 username: string | FormDataEntryValue
 password: string | FormDataEntryValue
 public constructor(username: string | FormDataEntryValue, password: string | FormDataEntryValue) {
  this.username = username;
  this.password = password;
 }
}
// interface UserData {
//     username: string | FormDataEntryValue
//     password: string | FormDataEntryValue
// }
export default UserData
