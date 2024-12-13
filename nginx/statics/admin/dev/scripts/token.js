// リンク用のトークンからゲーム用のトークンを生成
async function GenGameToken(teamid) {
    // トークンを取得 (Pocketbase から取得)
    const token = await GetToken();

    // 送信
    const req = await fetch("/admin/team/link", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": token,
            "teamid": teamid,
        },
    });

    return req.json();
}
