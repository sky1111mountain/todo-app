[![Go Test](https://github.com/sky1111mountain/todo-app/actions/workflows/go-test.yml/badge.svg)](https://github.com/sky1111mountain/todo-app/actions/workflows/go-test.yml)

# Todo Management API

Go言語で構築した、堅牢でテスト済みのTodo管理APIです。

## アプリ概要

シンプルなTodo管理機能に加え、優先度やステータスによるフィルタリング機能を備えています。

## こだわりポイント

**開発環境のコンテナ化**:
Docker Composeを採用し、どの環境でも docker-compose up 一発でデータベースを含む開発環境が整うように設計しました。

**CIによるテストの自動化**:
GitHub Actionsを導入し、テストに合格したコードのみがアップロードされる仕組みを構築しています。

**データベースの信頼性確保**:
pingによるDB接続を確認してからテストを実行し、信頼性の高いCIを実現しました。

## 使用している技術

- Language:Go
- Database:MySQL
- Infrastructure:Docker Compose
- CI/CD:GitHub Actions

## 機能詳細

**CRUD操作:タスクの登録・取得・更新・削除**

**フィルタリング**:

- 未完了タスクの全取得
- ID指定取得
- 全タスク取得（all パラメータ）

**ステータス管理**：

- Priority(high/middle/low)
- Status(Done/Not_Done)

## データベース設計

| Column     | Type         | Key | Notes                   |
| ---------- | ------------ | --- | ----------------------- |
| id         | INT          | PK  | Auto Increment          |
| task       | VARCHAR(255) |     |                         |
| priority   | ENUM         |     | 'high', 'middle', 'low' |
| status     | ENUM         |     | 'not_done', 'done'      |
| username   | VARCHAR(100) |     |                         |
| created_at | DATETIME     |     |                         |
