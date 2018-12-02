import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {map} from 'rxjs/operators';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  token: string = null;
  private goServer = 'http://localhost:8080/api';
  private wgServer = 'https://api.worldofwarships.asia/wows';
  private wgStatusCheck = res => {
    if (res.status === 'ok') {
      return res.data;
    } else {
      throw null;
    }
  };

  constructor(
    private http: HttpClient,
  ) {
  }

  getToken(): Observable<string> {
    return this.http.get<any>(
      `${this.goServer}/token`
    ).pipe(map((response) => {
      if (response.ok) {
        this.token = response.data.application_id;
        return this.token;
      } else {
        throw null;
      }
    }));
  }

  setup(param: any) {
    return this.http.post<any>(
      `${this.goServer}/setup`,
      param,
    );
  }

  getMatch() {
    return this.http.get<any>(
      `${this.goServer}/match`
    ).pipe(map(res => {
      if (res.ok) {
        return res.data;
      } else {
        throw null;
      }
    }));
  }

  getShip(shipId: number) {
    return this.http.get<any>(
      `${this.wgServer}/encyclopedia/ships/`, {
        params: {
          application_id: this.token,
          ship_id: `${shipId}`,
          language: 'ja',
        }
      }
    ).pipe(map(this.wgStatusCheck));
  }

  getPlayerId(playerName: string) {
    return this.http.get<any>(
      `${this.wgServer}/account/list/`, {
        params: {
          application_id: this.token,
          search: playerName,
        }
      }
    ).pipe(map(this.wgStatusCheck));
  }

  getPlayerShipStats(playerId: number, shipId: number) {
    return this.http.get<any>(
      `${this.wgServer}/ships/stats/`, {
        params: {
          application_id: this.token,
          account_id: `${playerId}`,
          ship_id: `${shipId}`,
        }
      }
    ).pipe(map(this.wgStatusCheck));
  }

  getPlayerStats(playerId: number) {
    return this.http.get<any>(
      `${this.wgServer}/account/info/`, {
        params: {
          application_id: this.token,
          account_id: `${playerId}`,
        }
      }
    ).pipe(map(this.wgStatusCheck));
  }
}
