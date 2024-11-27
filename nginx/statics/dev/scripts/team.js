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

// 初期化関数
async function Init() {
    const teams = await ListTeam();
    console.log(teams["msg"]);

    // チームリスと取得
    const teamlist = document.getElementById("team_list");

    teams["msg"].forEach((team) => {
        const basetr = document.createElement("tr");

        // チームID
        const idtd = document.createElement("td");
        idtd.textContent = team["TeamID"];

        // チーム名
        const nametd = document.createElement("td");
        nametd.textContent = team["Name"];

        // 作成者   
        const creatortd = document.createElement("td");
        creatortd.textContent = team["Creator"];

        // 状態
        const statustd = document.createElement("td");
        statustd.textContent = team["Status"];

        // 削除ボタン
        const deltd = document.createElement("td");
        const delbtn = document.createElement("button");
        delbtn.textContent = "削除";
        delbtn.addEventListener("click", () => {
            DeleteTeam(team["TeamID"]);
        });
        deltd.appendChild(delbtn);

        // 追加
        basetr.appendChild(idtd);
        basetr.appendChild(nametd);
        basetr.appendChild(creatortd);
        basetr.appendChild(statustd);
        basetr.appendChild(deltd);
        
        teamlist.appendChild(basetr);
    });
}

Init();