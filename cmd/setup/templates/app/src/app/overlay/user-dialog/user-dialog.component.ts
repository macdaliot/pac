import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Inject } from '@angular/core';
import { FormControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { User } from '../../model/user';
import { UserService } from '../../service/user/user.service';
@Component({
  selector: 'app-user-dialog',
  templateUrl: './user-dialog.component.html',
  styleUrls: ['./user-dialog.component.scss']
})
export class UserDialogComponent implements OnInit {
  form: FormGroup;
  titles = ['Mr.', 'Mrs.', 'Ms.', 'Dr.'];
  roles = ['admin', 'manager', 'user'];
  constructor(
  public dialogRef: MatDialogRef<UserDialogComponent>,
  @Inject(MAT_DIALOG_DATA) public data: any,
  private fb: FormBuilder,
  private userService: UserService) {

  	this.form = this.fb.group({
  		title: '',
        firstName: ['', Validators.required],
        lastName: ['', Validators.required],
        phone: [''],
        email: ['', [Validators.required, Validators.email]],
        birthdate: [null, Validators.required],
        notes: '',
        role: ''
      });

  	this.form.patchValue(data);

  }

  ngOnInit() {
  	this.userService.completionObservable.subscribe(() => this.dialogRef.close());
  }

  onSave(): void {
    const userObject = new User(this.form.value);
    this.userService.editUser(this.data._id, userObject);
  }

  onAdd(): void {
    const userObject = new User(this.form.value);
    this.userService.addUser(userObject);

  }

  onCancel(): void {
    this.dialogRef.close();
  }

}
