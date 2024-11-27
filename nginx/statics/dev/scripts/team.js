const create_team = document.getElementById("create_team");

create_team.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    try {
        // フォームからデータ取得
        const formData = new FormData(create_team);
        const data = {
            name: formData.get("name"),
        };

        // チーム作成実行
        const team = await CreateTeam(data.name);

        console.log(team);

        // チームリスト更新
        await RefreshTeam();

    } catch (ex) {
        console.error(ex);
    }
});


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


// チームリスと取得
const teamlist = document.getElementById("team_list");

// チームを表示する関数
function ShowTeam(id, name, creator, status) {
    const basetr = document.createElement("tr");

    // チームID
    const idtd = document.createElement("td");
    idtd.textContent = id;

    // チーム名
    const nametd = document.createElement("td");
    nametd.textContent = name;

    // 作成者   
    const creatortd = document.createElement("td");
    creatortd.textContent = creator;

    // 状態
    const statustd = document.createElement("td");
    statustd.textContent = status;

    // 削除ボタン
    const deltd = document.createElement("td");
    const delbtn = document.createElement("button");
    delbtn.textContent = "削除";
    delbtn.addEventListener("click",async () => {
        console.log(await DeleteTeam(id));

        // リロード
        RefreshTeam();
    });
    deltd.appendChild(delbtn);

    // 追加
    basetr.appendChild(idtd);
    basetr.appendChild(nametd);
    basetr.appendChild(creatortd);
    basetr.appendChild(statustd);
    basetr.appendChild(deltd);

    teamlist.appendChild(basetr);
}

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

async function RefreshTeam() {
    // チームリストをクリア
    teamlist.innerHTML = "";

    const teams = await ListTeam();
    console.log(teams["msg"]);

    teams["msg"].forEach((team) => {
        ShowTeam(team["TeamID"], team["Name"], team["Creator"], team["Status"]);
    });
}

// 初期化関数
async function Init() {
    await RefreshTeam();
}

Init();