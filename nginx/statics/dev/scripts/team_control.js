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
    const buttontd = document.createElement("td");
    const delbtn = document.createElement("button");
    delbtn.textContent = "削除";
    delbtn.addEventListener("click",async () => {
        console.log(await DeleteTeam(id));

        // リロード
        await RefreshTeam();
    });
    buttontd.appendChild(delbtn);

    // トークン生成ボタン
    const tokenbtn = document.createElement("button");
    tokenbtn.textContent = "トークン生成";
    tokenbtn.addEventListener("click", async () => {
        try {
            // トークンを取得
            const token_data = await GenLinkToken(id);

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
    basetr.appendChild(buttontd);

    teamlist.appendChild(basetr);
}


// 使用する階を設定するフォーム
const floors_form = document.getElementById("floors_form");

floors_form.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    try {
        // フォームからデータ取得
        const formData = new FormData(floors_form);

        // リストにする
        const floors = [];

        // for of を使ってチェックボックスを取得
        for (const floor of formData.getAll("floors")) {
            floors.push(Number(floor));
        }

        // トークンを取得
        const token = await GetToken();

        // リクエストを送る
        const req = await fetch("/admin/game/floors", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": token,
            },
            body: JSON.stringify({ "floors": floors }),
        });

        console.log(await req.json());

    } catch (ex) {
        console.error(ex);
    }
});

// 階のチェックを表示するエリア
const floors_area = document.getElementById("floors_area");

// 初期化関数
async function Init() {
    try {
        // チームリスト更新
        await RefreshTeam();

        // 使用する階を設定を取得
        const floors = await GetFloors();
        
        floors["msg"].forEach((floor) => {
            // チェックボックスを作成
            const checkbox = document.createElement("input");
            checkbox.type = "checkbox";
            checkbox.id = "floor_" + floor["FloorNum"];
            checkbox.name = "floors";
            checkbox.value = floor["FloorNum"];
            // チェックボックスにチェックを入れる
            checkbox.checked = floor["IsUsed"];
            checkbox.classList.add("form-check-input");

            // ラベルを作成
            const label = document.createElement("label");
            label.htmlFor = "floor_" + floor["FloorNum"];
            label.textContent = floor["FloorNum"];
            label.classList.add("form-check-label");

            // チェックボックスとラベルを追加
            floors_area.appendChild(checkbox);
            floors_area.appendChild(label);

            // 改行を追加
            floors_area.appendChild(document.createElement("br"));
        })
    } catch (ex) {
        console.error(ex);
        // alert("取得に失敗しました");
        // 認証していない場合ログインに飛ばす
        // window.location.href = LoginURL;
    }
}

Init();