import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { WowsBlameSetupParam } from './models';

@Injectable()
export class AppService {

    private host = 'localhost';
    private port = 8080;

    constructor(
        private http: Http,
    ) { }

    isSetup() {
        return this.http.get(
            `http://${this.host}:${this.port}/api/version`
        ).map(response => response.json()['ok'] === true);
    }

    setup(param: WowsBlameSetupParam) {
        return this.http.post(
            `http://${this.host}:${this.port}/api/setup`,
            param,
        );
    }
}
