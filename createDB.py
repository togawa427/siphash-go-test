import sqlite3
import os

DB_PATH = "/Users/togawa/GitHub/siphash-go-test/tmpLog.db"

def create_database():
    # 既存のDBがあれば削除（初期化したい場合のみ有効化）
    if os.path.exists(DB_PATH):
        os.remove(DB_PATH)

    # SQLiteデータベースに接続（なければ作成される）
    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()

    # usersテーブルの作成
    cur.execute("""
    CREATE TABLE IF NOT EXISTS users (
        uuid TEXT PRIMARY KEY,
        rssi INTEGER,
        count INTEGER
    )
    """)

    # 変更を保存して接続を閉じる
    conn.commit()
    conn.close()

    print(f"データベース '{DB_PATH}' に users テーブルを作成しました！")


if __name__ == "__main__":
    create_database()