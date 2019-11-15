import { Component, OnInit } from '@angular/core';
import { UserService } from '../../service/user/user.service';
import { User } from '../../model/user';
import { faEdit, faTrash, faPlus } from '@fortawesome/free-solid-svg-icons';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { UserDialogComponent } from '../../overlay/user-dialog/user-dialog.component';
// import { userList } from '../../model/user-list';
@Component({
  selector: 'app-page-two',
  templateUrl: './page-two.component.html',
  styleUrls: ['./page-two.component.scss']
})
export class PageTwoComponent implements OnInit {
  users: User[];
  faEdit = faEdit;
  faTrash = faTrash;
  faPlus = faPlus;
  displayedColumns: string[] = ['name', 'role', 'phone', 'email', 'date', 'actions'];
  constructor(private userService: UserService, public dialog: MatDialog) { }

  ngOnInit() {
    this.userService.usersObservable.subscribe(res => { this.users = res; });
    this.userService.loadUsers();

  }

  editUser(user) {
    const dialogRef = this.dialog.open(UserDialogComponent, {
      width: '620px',
      data: user
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
     // this.animal = result;
    });

  }

  addUser() {
     const dialogRef = this.dialog.open(UserDialogComponent, {
      width: '620px',
      data: new User({})
    });

     dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
     // this.animal = result;
    });

  }
  deleteUser(user) {
    this.userService.deleteUser(user);
  }



}
