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
