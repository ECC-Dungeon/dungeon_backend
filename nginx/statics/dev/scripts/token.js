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
        const linkData = await InitLink(data.token);
        console.log(linkData);
    } catch (ex) {
        console.error(ex);
    }
});

async function InitLink(token) {
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