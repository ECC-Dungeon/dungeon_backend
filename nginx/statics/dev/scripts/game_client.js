const read_form = document.getElementById("read_token_form");

read_form.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    try {
        // フォームからデータ取得
        const formData = new FormData(read_form);
        const data = {
            token: formData.get("token"),
        };

        // ローカルストレージに保存
        localStorage.setItem("game_token", data.token);

    } catch (ex) {
        console.error(ex);
    }
});

async function GetTeam() {
    // チーム情報取得
    const teamData = await GetTeamInfo();

    console.log(teamData["msg"]);
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

const floors_form = document.getElementById("floors_form");

floors_form.addEventListener("submit", async (evt) => {
    evt.preventDefault();

    try {
        // トークンを取得
        const token = localStorage.getItem("game_token");

        const req2 = await fetch("/admin/game/info", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": token,
            },
        });

        // ゲームIDを取得
        const gameData = await req2.json();

        // フォームからデータ取得
        const req = await fetch("/admin/game/floor", {
            method: "GET",
            headers: {
                "gameid": gameData["msg"]["GameID"],
            },
        });

        console.log(await req.json());

    } catch (ex) {
        console.error(ex);
    }
});

async function StartGame() {
    const req = await fetch("/admin/game/start2", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": localStorage.getItem("game_token"),
        },
    });

    console.log(await req.json());
}

async function GameTest() {
    const req = await fetch("/game/", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": localStorage.getItem("game_token"),
        },
    });

    console.log(await req.json());
}