const read_form = document.getElementById("read_token_form");

read_form.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    try {
        // フォームからデータ取得
        const formData = new FormData(read_form);
        const data = {
            token: formData.get("token"),
        };

        // トークン読み込み実行
        const gameData = await GenGameToken(data.token);
        if (gameData["msg"] == undefined) {
            alert("ゲーム用トークンを取得できませんでした");
            return;
        }

        console.log(gameData["msg"]);

        // ローカルストレージに保存
        localStorage.setItem("game_token", gameData["msg"]);

    } catch (ex) {
        console.error(ex);
    }
});

async function GetTeam() {
    // チーム情報取得
    const teamData = await GetTeamInfo();

    console.log(teamData["msg"]);
}

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

const nickname_form = document.getElementById("nickname_form");

nickname_form.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    try {
        // トークンを取得
        const token = localStorage.getItem("game_token");

        // フォームからデータ取得
        const formData = new FormData(nickname_form);
        const data = {
            nickname: formData.get("nickname"),
        };

        // ニックネーム保存実行
        const req = await fetch("/admin/game/tname", {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": token,
            },
            body: JSON.stringify({
                "name": data.nickname
            }),
        });

        console.log(await req.json());

    } catch (ex) {
        console.error(ex);
    }
});