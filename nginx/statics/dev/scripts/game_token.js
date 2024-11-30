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
        console.log(gameData);
    } catch (ex) {
        console.error(ex);
    }
});