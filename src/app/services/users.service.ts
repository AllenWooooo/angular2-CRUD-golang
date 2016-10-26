import {Injectable} from '@angular/core';
import {Http} from '@angular/http';

import {Observable} from 'rxjs/Observable';

import {AppService} from './app.service';


export class User {
  constructor(
    public id: string = '',
    public name: string = '',
    public balance: number = 0,
    public modifying: boolean = false) {}
}

@Injectable()
export class UsersService {
  private prefix = '/users';

  constructor(private service: AppService) {
  }

  list(): Observable<User[]> {
    return this.service.get(this.prefix);
  }

  create(newUser: User): Observable<User[]> {
    return this.service.post(this.prefix, newUser);
  }

  remove(id: string): Observable<User[]> {
    return this.service.delete(`${this.prefix}/${id}`);
  }

  update(user: User): Observable<User[]> {
    return this.service.put(`${this.prefix}/`, user);
  }
}
