import { Injectable } from '@angular/core';
import { IUser } from '@pyramid-systems/domain';
import { UrlService } from '../url/url.service';
import { BehaviorSubject, Subject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { User } from '../../model/user';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private urlService: UrlService, private httpClient: HttpClient) {}

  private _users: User[];
  private endpoint = this.urlService.apiUrl + 'myuser/';
  private usersSubject: BehaviorSubject<User[]> = new BehaviorSubject([]);
  private completionSubject: Subject<any> = new Subject();


  get usersObservable(): Observable<User[]> {
    return this.usersSubject.asObservable();
  }

  get completionObservable(): Observable<User[]> {
    return this.completionSubject.asObservable();
  }

  loadUsers() {
    this.httpClient.get<User[]>(this.endpoint)
          .pipe(map(response => {
                   const u = response.map(data => new User(data));
                   return u;
               })).subscribe(data => {
                             this._users = data;
                             this.usersSubject.next(this._users.slice(0));
                           }, error => console.log('Error loading users.'));
  }


  editUser(id: string, user: User) {
  //  let p = new Promise<boolean>((resolve, reject)=> {});
      const u = {firstName: user.firstName, lastName: user.lastName, birthdate: user.birthdate, phone: user.phone, email: user.email, notes: user.notes, title: user.title, role: user.role}; // Object.assign(Object, user);
      // u.birthdate  = '1970-10-03';
      this.httpClient.put<User>(this.endpoint + id, u)
        .subscribe(data => {
         // user._id = id;
          this._users.forEach((t, i) => {
           // debugger;
            this.completionSubject.next('update');

            if (t._id === id) { this._users[i] = user; }
          });

          this.usersSubject.next(this._users);
        }, error => console.log('Could not update todo.'));
   //   return p;
  }

  addUser(user: User) {
  //  let p = new Promise<boolean>((resolve, reject)=> {});
    //  let u = {firstName: user.firstName, lastName: user.firstName, birthdate: '1970-10-03',phone: user.phone, email: user.email, notes: user.notes, title: user.title, role: user.role};//Object.assign(Object, user);
      // u.birthdate  = '1970-10-03';
     this.httpClient.post<User>(this.endpoint, user)
        .subscribe(data => {
          const u = new User(user);
          u._id = data[0];
          this._users.unshift(u);


          this.usersSubject.next(this._users);
          this.completionSubject.next('update');
        }, error => console.log('Could not update todo.'));
   //   return p;
  }

  deleteUser(user: User) {
      this.httpClient.delete<boolean>(this.endpoint + user._id).subscribe(res => {
        this._users.some((t, i) => {
          if (t._id === user._id) { this._users.splice(i, 1);
                                    this.usersSubject.next(this._users);
        }
       });
      });
  }

  // create(todo: Todo) {
  //   this.http.post<Todo>(`${this.baseUrl}/todos`, JSON.stringify(todo)).subscribe(data => {
  //       this.dataStore.todos.push(data);
  //       this._todos.next(Object.assign({}, this.dataStore).todos);
  //     }, error => console.log('Could not create todo.'));
  // }

}
