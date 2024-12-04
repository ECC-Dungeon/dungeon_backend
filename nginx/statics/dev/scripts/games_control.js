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
        const game = await CreateGame(data.name);

        console.log(game);

        // チームリスト更新
        await RefreshGame();

    } catch (ex) {
        console.error(ex);
    }
});

async function RefreshGame() {
    // チームリスト更新
    const gamesData = await GetGames();

    // ゲームリスト更新
    const games = gamesData["msg"];

    // 既存のチームを削除
    game_list.innerHTML = "";

    // チームを回す
    games.forEach((game) => {
        console.log(game);
        ShowGame(game["GameID"],game["Name"],game["CreatorID"]);
    });
}

const game_list = document.getElementById("game_list");
function ShowGame(gameid,Gamename,CreatorID) {
    const tr = document.createElement("tr");

    // チームID
    const idtd = document.createElement("td");
    idtd.textContent = gameid;

    //クリックイベント
    idtd.addEventListener("click",async () => {
        // リロード
        // ゲームの詳細に遷移
        window.location.href = "./team_control.html?gameid=" + gameid;
    });

    // チーム名
    const nametd = document.createElement("td");
    nametd.textContent = Gamename;

    // 作成者   
    const creatortd = document.createElement("td");
    creatortd.textContent = CreatorID;

    // 削除ボタン
    const buttontd = document.createElement("td");
    const delbtn = document.createElement("button");
    delbtn.textContent = "削除";
    delbtn.addEventListener("click",async () => {
        console.log(await DeleteGame(gameid));
        // リロード
        // チームリスト更新
        await RefreshGame();
    });
    buttontd.appendChild(delbtn);

    tr.appendChild(idtd);
    tr.appendChild(nametd);
    tr.appendChild(creatortd);
    tr.appendChild(buttontd);
    game_list.appendChild(tr);
}

// 初期化関数
async function Init() {
    try {
        // チームリスト更新
        await RefreshGame();
    } catch (ex) {
        console.error(ex);
    }
}

Init();