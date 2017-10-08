export class WowsBlameSetupResponse {
    ok: boolean;
    app_id_ok: boolean;
    app_id_message: string;
    path_ok: boolean;
    path_message: string;
    message: string;

    static createFrom(source: any) {
        const result = new WowsBlameSetupResponse();
        result.ok = source['ok'];
        result.app_id_ok = source['app_id_ok'];
        result.app_id_message = source['app_id_message'];
        result.path_ok = source['path_ok'];
        result.path_message = source['path_message'];
        result.message = source['message'];
        return result;
    }
}

export class WowsBlameSetupParam {
    application_id: string;
    game_path: string;

    static createFrom(source: any) {
        const result = new WowsBlameSetupParam();
        result.application_id = source['application_id'];
        result.game_path = source['game_path'];
        return result;
    }
}

export class ArenaVehiclesInfo {
    shipId: number;
    relation: number;
    id: number;
    name: string;

    static createFrom(source: any) {
        const result = new ArenaVehiclesInfo();
        result.shipId = source['shipId'];
        result.relation = source['relation'];
        result.id = source['id'];
        result.name = source['name'];
        return result;
    }
}

export class ArenaInfo {
    clientVersionFromXml: string;
    gameMode: number;
    clientVersionFromExe: string;
    mapDisplayName: string;
    mapId: number;
    matchGroup: string;
    duration: number;
    gameLogic: string;
    name: string;
    scenario: string;
    playerID: number;
    vehicles: ArenaVehiclesInfo[];
    playersPerTeam: number;
    dateTime: string;
    mapName: string;
    playerName: string;
    scenarioConfigId: number;
    teamsCount: number;
    logic: string;
    playerVehicle: string;

    static createFrom(source: any) {
        const result = new ArenaInfo();
        result.clientVersionFromXml = source['clientVersionFromXml'];
        result.gameMode = source['gameMode'];
        result.clientVersionFromExe = source['clientVersionFromExe'];
        result.mapDisplayName = source['mapDisplayName'];
        result.mapId = source['mapId'];
        result.matchGroup = source['matchGroup'];
        result.duration = source['duration'];
        result.gameLogic = source['gameLogic'];
        result.name = source['name'];
        result.scenario = source['scenario'];
        result.playerID = source['playerID'];
        result.vehicles = source['vehicles'] ? source['vehicles']
            .map(function (element) { return ArenaVehiclesInfo.createFrom(element); }) : null;
        result.playersPerTeam = source['playersPerTeam'];
        result.dateTime = source['dateTime'];
        result.mapName = source['mapName'];
        result.playerName = source['playerName'];
        result.scenarioConfigId = source['scenarioConfigId'];
        result.teamsCount = source['teamsCount'];
        result.logic = source['logic'];
        result.playerVehicle = source['playerVehicle'];
        return result;
    }
}

export class WowsBlameMatchResponse {
    ok: boolean;
    message: string;
    data: ArenaInfo;

    static createFrom(source: any) {
        const result = new WowsBlameMatchResponse();
        result.ok = source['ok'];
        result.message = source['message'];
        result.data = source['data'] ? ArenaInfo.createFrom(source['data']) : null;
        return result;
    }
}

export class WowsBlamePlayerPayload {
    win_rate: number;
    total_battle_count: number;
    account_name: string;
    account_id: string;
    ship_name: string;
    ship_id: string;
    ship_type: string;
    ship_win_rate: number;
    ship_battle_count: number;

    static createFrom(source: any) {
        const result = new WowsBlamePlayerPayload();
        result.win_rate = source['win_rate'];
        result.total_battle_count = source['total_battle_count'];
        result.account_name = source['account_name'];
        result.account_id = source['account_id'];
        result.ship_name = source['ship_name'];
        result.ship_id = source['ship_id'];
        result.ship_type = source['ship_type'];
        result.ship_win_rate = source['ship_win_rate'];
        result.ship_battle_count = source['ship_battle_count'];
        return result;
    }
}

export class WowsBlamePlayer {
    ok: boolean;
    data: WowsBlamePlayerPayload;
    message: string;

    static createFrom(source: any) {
        const result = new WowsBlamePlayer();
        result.ok = source['ok'];
        result.data = source['data'] ? WowsBlamePlayerPayload.createFrom(source['data']) : null;
        result.message = source['message'];
        return result;
    }
}
