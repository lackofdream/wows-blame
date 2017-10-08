import { Component, OnInit } from '@angular/core';
import { AppService } from '../app.service';
import { WowsBlameMatchResponse, ArenaInfo, WowsBlamePlayerPayload, WowsBlamePlayer } from '../models';
import { Observable } from 'rxjs/Rx';
import { Response } from '@angular/http';

@Component({
  selector: 'app-blame',
  templateUrl: './blame.component.html',
  styleUrls: ['./blame.component.css']
})
export class BlameComponent implements OnInit {

  private cache: ArenaInfo;
  playerInfo = new Object();
  shyPlayers: string[] = [];

  players: WowsBlamePlayerPayload[][][] = [[[], [], [], []], [[], [], [], []]];

  constructor(
    private appService: AppService
  ) { }

  private cacheResponse = response => {
    const match = response.json() as WowsBlameMatchResponse;
    if (!match.ok || (this.cache && match.data.dateTime === this.cache.dateTime)) {
      return undefined;
    }
    this.cache = match.data;
    return this.cache;
  }

  ngOnInit() {
    Observable.interval(3000).exhaustMap(() => this.appService.match().map(
      this.cacheResponse
    ).flatMap(this.lookup).map(this.handlePlayer)).subscribe(this.updateView);
  }

  private lookup = (info: ArenaInfo) => {
    if (!info) {
      return Observable.of(undefined);
    }
    return Observable.merge(...info.vehicles.map(player => this.appService.player(player.name, player.shipId)));
  }

  private handlePlayer = response => {
    if (!response) {
      return false;
    }
    const body = response.json() as WowsBlamePlayer;
    const player = body.data;
    if (!player.account_name) {
      return false;
    }
    this.playerInfo[player.account_name] = player;
    return true;
  }

  private updateView = (result) => {
    if (!result) {
      return;
    }
    this.players = [[[], [], [], []], [[], [], [], []]];
    this.shyPlayers = [];

    this.cache.vehicles.map(v => {
      if (!this.playerInfo[v.name]) {
        this.shyPlayers.push(v.name);
        return;
      }
      const info = this.playerInfo[v.name] as WowsBlamePlayerPayload;
      const firstIndex = v.relation === 0 ? 0 : v.relation - 1;
      switch (info.ship_type) {
        case 'AirCarrier':
          this.players[firstIndex][0].push(info);
          break;
        case 'Battleship':
          this.players[firstIndex][1].push(info);
          break;
        case 'Cruiser':
          this.players[firstIndex][2].push(info);
          break;
        case 'Destroyer':
          this.players[firstIndex][3].push(info);
          break;
      }
    });

    for (let idx = 0; idx < 2; idx++) {
      for (let _idx = 0; _idx < 4; _idx++) {
        this.players[idx][_idx] = this.players[idx][_idx].sort((a: WowsBlamePlayerPayload, b: WowsBlamePlayerPayload) => {
          if (a.ship_win_rate > b.ship_win_rate) {
            return -1;
          } else if (a.ship_win_rate === b.ship_win_rate) {
            return 0;
          }
          return 1;
        });
      }
    }
  }

}
