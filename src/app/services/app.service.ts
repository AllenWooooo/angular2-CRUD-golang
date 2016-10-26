import {Injectable} from '@angular/core';
import {Http, Response} from '@angular/http';

import {Observable} from 'rxjs/Observable';

export class ApiHttpService {
  private methods: Array<string> = ['get', 'delete'];
  private bodyMethods: Array<string> = ['post', 'put', 'patch'];

  public get: any;
  public delete: any;
  public post: any;
  public put: any;
  public patch: any;

  constructor(private http: Http) {
    this.init();
  }

  init() {
    for (let method of this.methods) {
      this[method] = (url: string, option?: Object) => {
        return this.request(method, url, option);
      };
    }

    for (let method of this.bodyMethods) {
      this[method] = (url: string, body:any, option?: Object) => {
        return this.requestWithBody(method, url, body, option);
      };
    }
  }

  private extractData(res: Response) {
    let data = res.json();
    return data || {};
  }

  private handleError(error: Response | any) {
    let errMsg: string;
    const body = error.json() || '';
    const err = body.error || JSON.stringify(body);

    errMsg = `${error.status} - ${error.statusText || ''} ${err}`;

    return Observable.throw(errMsg);
  }

  private request(method: string, url: string, option?: Object) {
    return this.http[method](url, option)
               .map(this.extractData)
               .catch(this.handleError);
  }

  private requestWithBody(method: string, url: string, body: any, option?: Object) {
    return this.http[method](url, body, option)
               .map(this.extractData)
               .catch(this.handleError);
  }
}

@Injectable()
export class AppService extends ApiHttpService {
  constructor(http: Http) {
    super(http);
  }
}