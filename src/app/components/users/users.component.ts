import {Component} from '@angular/core';

import {User, UsersService} from '../../services/users.service';

@Component({
  selector: 'users',
  templateUrl: './users.component.html',
  providers: [UsersService]
})

export class UsersComponent {
  users: User[];
  newUser: User = new User();
  modifyingUser: User;
  creating: boolean = false;

  constructor(private usersService: UsersService) {}

  ngOnInit() {
    this.usersService
        .list()
        .subscribe(
          users => this.users = users);
  }

  create() {
    if (this.creating && this.newUser.name) {
      this.usersService
          .create(this.newUser)
          .subscribe(
            users => this.users = users,
            null,
            () => {
              this.creating = false;
              this.newUser = new User();
            });

      return;
    }

    this.creating = !this.creating;
  }

  remove(id: string) {
    this.usersService
        .remove(id)
        .subscribe(
          users => this.users = users);
  }

  modify(user: User) {
    user.modifying = true;
    this.modifyingUser = Object.assign({}, user);
  }

  cancel(user: User) {
    user = Object.assign(user, this.modifyingUser);
    user.modifying = false;
  }

  update(user: User) {
    if (user.name) {
      this.usersService
          .update(user)
          .subscribe(
            users => this.users = users,
            null,
            () => user.modifying = false);
    }
  }
}
