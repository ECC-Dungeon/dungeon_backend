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

const floors_form = document.getElementById("floors_form");

floors_form.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    try {
        // フォームからデータ取得
        const formData = new FormData(floors_form);
        const data = {
            floors: formData.get("floors"),
        };

        console.log(data.floors);

        // チームリスト更新
        await RefreshTeam(gameid);

    } catch (ex) {
        console.error(ex);
    }
});

async function GetFloors() {
    //pocketbaseから取得
    const token = await GetToken();

    // フロアのリストを取得
    const req = await fetch("/admin/game/floor", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
            "gameid": gameid
        },
    });

    return req.json();
}

const floors = document.getElementById("floors");
// フロアを表示する関数
function ShowFloor(floorNum,floorName) {
    const floor_div = document.createElement("div");
    
    // チェックボックス
    const floorCheckbox = document.createElement("input");
    floorCheckbox.type = "checkbox";
    floorCheckbox.name = "floors";
    floorCheckbox.value = floorNum;
    floor_div.appendChild(floorCheckbox);

    // フロア名
    const floorLabel = document.createElement("label");
    floorLabel.textContent = floorName;
    floor_div.appendChild(floorLabel);

    // 名前
    const floorInput = document.createElement("input");
    floorInput.type = "text";
    floorInput.value = floorName;
    floor_div.appendChild(floorInput);

    // 追加
    floors.appendChild(floor_div);
}

// 初期化関数
async function Init() {
    try {
        // チームリスト更新
        await RefreshTeam(gameid);
        
        // フロア取得
        const floors = await GetFloors();

        // 回す
        floors["msg"].forEach((floor) => {
            console.log(floor);
            ShowFloor(floor["FloorNum"],floor["Name"]);
        });
    } catch (ex) {
        console.error(ex);
        // alert("取得に失敗しました");
        // 認証していない場合ログインに飛ばす
        // window.location.href = LoginURL;
    }
}

Init();