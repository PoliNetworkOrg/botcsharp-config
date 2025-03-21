# botcsharp-config

I was tired of messing out with k8s resources to correctly pass the configuration json file for
the [csharp telegram bot](https://github.com/PoliNetworkOrg/PoliNetworkBot_CSharp), so I created a simple
Golang script that creates those configuration files from environment variabales.

## Usage

Use this as an `initContainer` in your k8s deployment.  
See [this example](https://github.com/PoliNetworkOrg/polinetwork-cd/blob/56ab55f29b1e83fcc9d6d1480836e75912178df6/bot-mat/src/deployment.yaml#L62-L112).

## Available configuration

These are the env variables you can set.

> [!NOTE]  
> Even if the var type is `bool`, `int` or others, you still need to set their k8s value as string

Columns notes:

- `Required` means that you must specify it, otherwise the container **PANICS**.
- `Default` is the default value for non-required fields.
- `Enum` points to the config enum if exists.

### Basic

| Key                 | Required | Default | Type     | Description                                                                                  |
| ------------------- | -------- | ------- | -------- | -------------------------------------------------------------------------------------------- |
| `OUT_DIR`           | Yes      |         | `string` | Path where configuration files are written.                                                  |
| `CREATE_BOT_CONFIG` | No       | `true`  | `bool`   | Whether to create configuration file `bots_info.json` (required to run the bot)              |
| `CREATE_DB_CONFIG`  | No       | `false` | `bool`   | Whether to create configuration file `dbconfig.json` (database config)                       |
| `CREATE_MAT_CONFIG` | No       | `false` | `bool`   | Whether to create configuration file `materialbotconfig.json` (Material Bot specific config) |

### Telegram Bot

| Key                            | Required | Default | Type      | Enum                                                                                                                                                |
| ------------------------------ | -------- | ------- | --------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `BOT_TOKEN`                    | Yes      |         | `string`  |                                                                                                                                                     |
| `BOT_ON_MESSAGES`              | Yes      |         | `string`  | [BotStartMethods](https://github.com/PoliNetworkOrg/PoliNetworkBot_CSharp/blob/master/PoliNetworkBot_CSharp/Code/Data/Constants/BotStartMethods.cs) |
| `BOT_TYPE_API`                 | No       | `1`     | `number`  | [BotConfigAPI](https://github.com/PoliNetworkOrg/PoliNetworkBot_CSharp/blob/master/PoliNetworkBot_CSharp/Code/Enums/BotTypeAPI.cs)                  |
| `BOT_ACCEPTED_MESSAGES`        | No       | `true`  | `bool`    |                                                                                                                                                     |
| `BOT_WEBSITE`                  | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_CONTACT_STRING`           | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_SESSION_USER_ID`          | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_API_HASH`                 | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_NUMBER_COUNTRY`           | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_NUMBER_NUMBER`            | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_PASSWORD_TO_AUTHENTICATE` | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_METHOD`                   | No       | `null`  | `string?` |                                                                                                                                                     |
| `BOT_USER_ID`                  | No       | `null`  | `int?`    |                                                                                                                                                     |
| `BOT_API_ID`                   | No       | `null`  | `int?`    |                                                                                                                                                     |

### DB Config

Ensure `CREATE_DB_CONFIG` is set to `true`. Otherwise the following vars will be useless.

| Key           | Required | Type     | Description             |
| ------------- | -------- | -------- | ----------------------- |
| `DB_HOST`     | Yes      | `string` | Database Host           |
| `DB_PORT`     | Yes      | `int`    | Database Port           |
| `DB_DATABASE` | Yes      | `string` | Database name           |
| `DB_USER`     | Yes      | `string` | Database Login User     |
| `DB_PASSWORD` | Yes      | `string` | Database Login Password |

### Material Bot (MAT) Config

Ensure `CREATE_MAT_CONFIG` is set to `true`. Otherwise the following vars will be useless.

| Key            | Required | Type     | Description                                                           |
| -------------- | -------- | -------- | --------------------------------------------------------------------- |
| `MAT_ROOT`     | Yes      | `string` | The git root directory on the container where repositories are stored |
| `MAT_PASSWORD` | Yes      | `string` | GitLab `polibot` password                                             |
