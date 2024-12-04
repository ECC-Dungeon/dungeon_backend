//ゲームid取得 (queryパラメータ)
const gameid = new URLSearchParams(window.location.search).get("gameid");

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
        const team = await CreateTeam(data.name,gameid);

        console.log(team);

        // チームリスト更新
        await RefreshTeam(gameid);

    } catch (ex) {
        console.error(ex);
    }
});

// チームリスと取得
const teamlist = document.getElementById("team_list");

// チームを表示する関数
function ShowTeam(id, name, creator, status, nickName) {
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

    // ニックネーム表示
    const nicktd = document.createElement("td");
    nicktd.textContent = nickName;

    // 削除ボタン
    const buttontd = document.createElement("td");
    const delbtn = document.createElement("button");
    delbtn.textContent = "削除";
    delbtn.addEventListener("click",async () => {
        console.log(await DeleteTeam(id));

        // リロード
        await RefreshTeam(gameid);
    });
    buttontd.appendChild(delbtn);

    // トークン生成ボタン
    const tokenbtn = document.createElement("button");
    tokenbtn.textContent = "トークン生成";
    tokenbtn.addEventListener("click", async () => {
        try {
            // トークンを取得
            const token_data = await GenGameToken(id);

            // トークンを表示
            console.log(token_data["msg"]);
        } catch (ex) {
            console.error(ex);
        }
    });
    buttontd.appendChild(tokenbtn);

    // リンク解除用ボタン    
    const linkbtn = document.createElement("button");
    linkbtn.textContent = "リンク解除";
    linkbtn.addEventListener("click", async () => {
        try {
            // トークンを取得
            const unlink_data = await RemoveLink(id);

            // トークンを表示
            console.log(unlink_data);
        } catch (ex) {
            console.error(ex);
        }
    });
    buttontd.appendChild(linkbtn);

    // 追加
    basetr.appendChild(idtd);
    basetr.appendChild(nametd);
    basetr.appendChild(creatortd);
    basetr.appendChild(statustd);
    basetr.appendChild(nicktd);
    basetr.appendChild(buttontd);

    teamlist.appendChild(basetr);
}

// 初期化関数
async function Init() {
    try {
        // チームリスト更新
        await RefreshTeam(gameid);

    } catch (ex) {
        console.error(ex);
        // alert("取得に失敗しました");
        // 認証していない場合ログインに飛ばす
        // window.location.href = LoginURL;
    }
}

Init();