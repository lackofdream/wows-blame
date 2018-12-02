import {Component} from '@angular/core';
import {ApiService} from '../api.service';
import {concat, Observable, of, Subscription, timer, zip} from 'rxjs';
import {flatMap, map, tap} from 'rxjs/operators';
import {fromArray} from 'rxjs/internal/observable/fromArray';

@Component({
  selector: 'app-blame',
  templateUrl: './blame.component.html',
  styleUrls: ['./blame.component.css']
})
export class BlameComponent {

  private matchInfo: any = {};
  public wiki = {
    ready: false,
    player: {},
    ship: {},
    playerWithShip: {},
    playerId: {},
  };
  private players = {
    f: [],
    e: [],
  };
  private lastBattleTime = '';
  private matchSub: Subscription;
  private tick: Observable<number> = timer(1000, 5000);
  private getMatch: Observable<any> = this.api.getMatch();
  private battleTimeChecker = (x: any) => {
    if (x.dateTime !== this.lastBattleTime) {
      this.lastBattleTime = x.dateTime;
      this.wiki.ready = false;
      this.players = {
        f: [],
        e: [],
      };
      return x;
    }
    throw null;
  };
  private matchToVehicles = x => {
    return fromArray(x.vehicles);
  };
  private buildShipWiki = (vehicle) => {
    return this.api.getShip(vehicle.shipId).pipe(
      map(x => {
        this.wiki.ship[Object.entries(x)[0][0]] = x[Object.entries(x)[0][0]];
      }),
    );
  };
  private buildPlayerWiki = (vehicle) => {
    return zip(of(vehicle), this.api.getPlayerId(vehicle.name)).pipe(
      flatMap((zipped: [any, any]) => {
        const v = zipped[0];
        const data = zipped[1];
        const player = {
          id: data[0].account_id,
          name: data[0].nickname,
        };
        this.wiki.playerId[data[0].nickname] = data[0].account_id;
        return zip(this.api.getPlayerShipStats(player.id, v.shipId), this.api.getPlayerStats(player.id));
      }),
      map((zipped: [any, any]) => {
        const playerWithShip = zipped[0], player = zipped[1];
        if (Object.entries(playerWithShip).length !== 0) {
          this.wiki.playerWithShip[Object.entries(playerWithShip)[0][0]] = Object.entries(playerWithShip)[0][1][0];
          this.wiki.player[Object.entries(player)[0][0]] = Object.entries(player)[0][1];
        }
      }),
    );
  };

  private buildWiki = (vehicle) => {
    return concat(this.buildShipWiki(vehicle), this.buildPlayerWiki(vehicle));
  };

  private fullCheck = this.getMatch.pipe(
    map(this.battleTimeChecker),
    tap(data => this.matchInfo = data),
    flatMap(this.matchToVehicles),
    flatMap(this.buildWiki),
  );

  private showPlayers = () => {
    const playersByType = {
      AirCarrier: [],
      Battleship: [],
      Cruiser: [],
      Destroyer: [],
    };
    for (const p of this.matchInfo.vehicles) {
      const shipType = this.wiki.ship[p.shipId].type;
      playersByType[shipType].push(p);
    }
    for (const [k, v] of Object.entries(playersByType)) {
      for (const p of v) {
        this.players[p.relation === 2 ? 'e' : 'f'].push(p);
      }
    }
    console.log(this.wiki);
    console.log(this.players);
    this.wiki.ready = true;
  };

  private handleTick = () => {
    if (!this.matchSub || this.matchSub.closed) {
      this.matchSub = this.fullCheck.subscribe(() => {
      }, (err) => {
      }, this.showPlayers);
    }
  };

  constructor(private api: ApiService) {
    this.tick.subscribe(this.handleTick);
  }

}
