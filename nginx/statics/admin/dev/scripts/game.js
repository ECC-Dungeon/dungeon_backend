
async function GetTeamInfo() {
    // ローカルストレージから取得
    const token = localStorage.getItem("game_token");

    // トークンを取得
    if (token == null) {
        alert("ゲーム用トークンを取得してください");
        return;
    }

    // 送信
    const req = await fetch("/admin/game/team", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
    });

    return req.json();
}

async function CreateGame(name) {
    // pocketbase からトークンを取得
    const token = await GetToken();

    // 送信
    const req = await fetch("/admin/game/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
        body: JSON.stringify({ "name": name }),
    });

    return req.json();
}

async function CreateGame(name) {
    //pocketbase からトークンを取得
    const token = await GetToken();

    // 送信
    const req = await fetch("/admin/game/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
        body: JSON.stringify({ "name": name }),
    });

    return req.json();
}

async function GetGames() {
    //pocketbase からトークンを取得
    const token = await GetToken();

    // 送信
    const req = await fetch("/admin/game/list", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
    });

    return req.json();
}

async function DeleteGame(gameid) {
    // ゲーム削除
    const token = await GetToken();

    // 送信
    const req = await fetch("/admin/game/delete", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
            "gameid": gameid
        },
    });

    return req.json();
}