// リンク用のトークンからゲーム用のトークンを生成
async function GenGameToken(token) {
    try {
        // 送信
        const req = await fetch("/admin/initlink", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ "token": token }),
        });

        return req.json();
    } catch (ex) {
        console.error(ex);
    }
}

// リンクトークンを作成する
async function GenLinkToken(teamid) {
    // トークンを取得
    const token = await GetToken();

    // リクエストを送る
    const req = await fetch("/admin/link/token", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
        },
        body: JSON.stringify({ "teamid": teamid }),
    });

    return req.json();
}