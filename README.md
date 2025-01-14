# Pocketbase 認証キットリポジトリ

## 環境構築
1. 自己署名SSL の鍵を生成するために
   - Mac Linux の人
     - ```./Genkey.sh ```
     を実行して全てエンター
    - Windows の人
      - ``` ./Genkey.bat ```
      を実行して全てエンター
    - しっぱしいた場合
        ファイル内容に書いてあるものをコピーして docker-compose.yaml があるディレクトリで実行してください。
2. コンテナを起動
    - 起動コマンド 
    ```
    python3 ./scripts/develop.py
    ```
3. 各種データ設定
   1. (必須) [Google Cloud Console](https://console.cloud.google.com/welcome) にアクセスして　Oauth 用の Client ID と シークレットを作成
   2. (オプション) [Github開発者設定](https://github.com/settings/developers) にアクセスして　Oauth 用の Client ID と シークレットを作成
   3. (オプション) [Discord開発者コンソール](https://discord.com/developers/applications) にアクセスして　Oauth 用の Client ID と シークレットを作成
   - リダイレクト URL には全て 
    ```
    https://localhost:8520/auth/api/oauth2-redirect
    ```
    を設定してください。
4. Pocketbase の初期設定
    1. [管理画面](https://localhost:8520/auth/_)にアクセスしてユーザーを作成
    2. [設定画面](https://localhost:8520/auth/_/#/settings) にアクセス
    3. Application URL を 
    ```
    https://localhost:8520/auth/
    ```
    にする
    4. [認証プロバイダ管理画面](https://localhost:8520/auth/_/#/settings/auth-providers)で (使うもののみ) Google, Github,Discord を設定する
5. 認証テスト
   - [ホーム](https://localhost:8520/statics/)にアクセス
   - ログインしてみる

## 機能を追加する場合
- utils.py を実行する
```
python3 ./scripts/utils.py
```
- 表示される指示に従ってプロジェクトを作成する
# 終わり！