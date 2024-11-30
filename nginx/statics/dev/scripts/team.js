// ゲームトークンを取得

async function ListTeam() {
    // トークンを取得
    const token = await GetToken();

    // リクエストを送る
    const req = await fetch("/admin/team/list", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
    });

    return req.json();
}

// チームリストを更新
async function RefreshTeam() {
    // チームリストをクリア
    teamlist.innerHTML = "";

    const teams = await ListTeam();
    console.log(teams["msg"]);

    teams["msg"].forEach((team) => {
        ShowTeam(team["TeamID"], team["Name"], team["Creator"], team["Status"]);
    });
}

// チームを削除する
async function DeleteTeam(teamid) {
    // トークンを取得
    const token = await GetToken();

    // リクエストを送る
    const req = await fetch("/admin/team/delete", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
        body: JSON.stringify({ "teamid": teamid }),
    });

    return req.json();
}


// 関数をまとめた処理
async function CreateTeam(name) {
    // トークンを取得
    const token = await GetToken();

    // リクエストを送る
    const req = await fetch("/admin/team/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
        body: JSON.stringify({ "name": name }),
    });

    return req.json();
}

// リンク解除
async function RemoveLink(teamid) {
    // トークンを取得
    const token = await GetToken();

    // リクエストを送る
    const req = await fetch("/admin/link/remove", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
        body: JSON.stringify({ "teamid": teamid }),
    });

    return req.json();
}