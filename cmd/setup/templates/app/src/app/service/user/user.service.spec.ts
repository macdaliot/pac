import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { UserService } from './user.service';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { User } from '../../model/user';
import { toArray } from 'rxjs/operators';
describe('UserService', () => {
//  let injector: TestBed;
 // let service: GithubApiService;
  let httpMock: HttpTestingController;

  const dummyUsers = [
 	 {_id: '5678', firstName: 'Mari', lastName: 'Merritt', phone: '6943344705', email: 'eu@ligulaAliquamerat.co.uk', title: 'Ms.', role: 'manager', notes: 'Morbi neque tellus, imperdiet non, vestibulum nec, euismod in, dolor. Fusce feugiat. Lorem ipsum', birthdate: '1958-02-05'},
     {_id: '1234', firstName: 'Hiram', lastName: 'Mclaughlin', phone: '3591368426', email: 'hendrerit.a@non.net', title: 'Mr.', role: 'manager', notes: 'vitae sodales nisi magna sed dui. Fusce aliquam, enim nec tempus scelerisque,', birthdate: '1953-11-25'}
      // { firstName: 'Doe' }
    ];
  beforeEach(() => {TestBed.configureTestingModule({
  	imports: [RouterTestingModule, HttpClientTestingModule],
  	providers: [{ provide: 'HOST', useValue: 'mockdata' }]

  });
                    httpMock = TestBed.get(HttpTestingController);
});

  afterEach(() => {
	 httpMock.verify();
  });

  it('should be created', () => {
    const service: UserService = TestBed.get(UserService);
    expect(service).toBeTruthy();
  });

  it('get should send messages on observable', () => {
    const service: UserService = TestBed.get(UserService);

    service.loadUsers();
    const req = httpMock.expectOne(`https://mockdata/api/myuser/`);
    expect(req.request.method).toBe('GET');
    req.flush(dummyUsers);

    service.usersObservable.subscribe(users => {
      expect(users.length).toBe(2);
      expect(users[0]).toEqual(new User(dummyUsers[0]));
    });

  });

  it('edit should send messages on observable', (done) => {
    const service: UserService = TestBed.get(UserService);

    service.loadUsers();
    const req = httpMock.expectOne(`https://mockdata/api/myuser/`);
    expect(req.request.method).toBe('GET');
    req.flush(dummyUsers);

    const newUser = Object.assign({}, dummyUsers[0]);
    newUser.firstName = 'David';
    service.editUser('5678', new User(newUser));

    const req2 = httpMock.expectOne(`https://mockdata/api/myuser/5678`);
    expect(req2.request.method).toBe('PUT');

    req2.flush([true]);

     service.usersObservable.subscribe(users => {
      expect(users.length).toBe(2);
      expect(users[0]).toEqual(new User(newUser));
      done();
    });

  });

  it('delete should send messages on observable', (done) => {
    const service: UserService = TestBed.get(UserService);

    service.loadUsers();
    const req = httpMock.expectOne(`https://mockdata/api/myuser/`);
    expect(req.request.method).toBe('GET');
    req.flush(dummyUsers);

    service.deleteUser(new User(dummyUsers[0]));

    const req2 = httpMock.expectOne(`https://mockdata/api/myuser/5678`);
    expect(req2.request.method).toBe('DELETE');

    req2.flush([true]);

    service.usersObservable.subscribe(users => {
      expect(users.length).toBe(1);
      done();
    });

  });

  // it('edit should send messages on observables', () => {
  //   const service: UserService = TestBed.get(UserService);

  //   service.loadUsers();
  //   const req = httpMock.expectOne(`https://mockdata/api/myuser/`);
  //   expect(req.request.method).toBe("GET");
  //   req.flush(dummyUsers);

  //    service.usersObservable.subscribe(users => {
  //     expect(users.length).toBe(2);
  //     expect(users[0]).toEqual(new User(dummyUsers[0]));
  //   });

  // });
});
