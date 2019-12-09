## API

### CLI Commands

The following table contains short descriptions for all cwctl operations:

| Command     | Alias | Usage                                                               |
| ----------- | ----- | ------------------------------------------------------------------- |
| connections | `con` | Manage connections configuration list                               |
| install     | `in`  | Pull pfe & performance images from dockerhub                        |
| project     |       | Manage Codewind projects                                            |
| remove      | `rm`  | Remove Codewind/Project docker images and the codewind network      |
| secclient   | `sc`  | Manage new or existing APPLICATION access configurations            |
| seckeyring  | `sk`  | Manage Codewind keys in the desktop keyring                         |
| secrealm    | `sr`  | Manage new or existing REALM configurations                         |
| secrole     | `sl`  | Manage realm based ACCESS roles                                     |
| sectoken    | `st`  | Authenticate with username and password to obtain an access_token   |
| secuser     | `su`  | Manage new or existing USER access configurations                   |
| start       |       | Start the Codewind containers                                       |
| status      |       | Print the installation status of Codewind                           |
| stop        |       | Stop the running Codewind containers                                |
| stop-all    |       | Stop all of the Codewind and project containers                     |
| templates   |       | Manage project templates                                            |
| version     |       | Print the versions of Codewind containers, for a given container    |
| help        | `h`   | Shows a list of commands or help for one command                    |

### Operation details




## connections
| Subcommand  | Alias | Syntax                                                             |
| ----------- | ----- | -------------------------------------------------------------------|
| add         |  a    | cwcwtl connections add --label DISPLAY_NAME --url CODEWIND_GATEKEEPER_INGRESS --username USERNAME|
| get         |  g    | cwctl connections get --conid CONNECTION_ID |
| list        |  ls   | cwctl connections list |
| remove      |  rm   | cwctl connections remove --conid CONNECTION_ID|
| reset       |       | cwctl connections reset |
| update      |  u    | cwctl connections update --conid CONNECTION_ID --label DISPLAY_NAME --url CODEWIND_GATEKEEPER_INGRESS --username USERNAME|

Subcommands:</br>

`add/a` - Add a new connection to the list

> **Flags:**
> --label value     A displayable name
> --url value       The ingress URL of the PFE instance

`update/u` - Update an existing connection

> **Flags:**
> --conid value     The Connection ID to update
> --label value     A displayable name
> --url value       The ingress URL of the PFE instance

`get/g` - Get a connection using its ID

> **Flags:**
> --conid value     The Connection ID to retrieve

`remove/rm` - Remove a connection from the list

> **Flags:**
> --conid value     A Connection ID

`list/ls` - List known connections

>**Note:** No additional flags

`reset` - Resets the connections list to a single local connection

>**Note:** No additional flags










### project

`--url/-u <value>` - URL of project to download

Subcommands:</br>

`bind` - Bind a project to Codewind for building and running
> **Flags:**
> --name,-n value               Project name
> --language,-l value           Project language
> --type,-t value               Project Type
> --path,-p value               Project Path
> --conid value                 Connection ID

`sync` - Synchronize a bound project to its connection
> **Flags:**
> --path,-p value               Project Path
> --id,-i value                 Project ID
> --time,-t value               Time of last project sync

`connection/con` - Manage the connection targets for a project

`set,s` - Sets the connection for a projectID
> **Flags:**
> --id,i value                  Project ID
> --conid value                 Connection ID

`get,g` - Gets connections for a projectID
> **Flags:**
> --id,i value                  Project ID

`remove,r` - Removes the connection from a projectID
> **Flags:**
> --id,i value                  Project ID

## install

`--tag/-t <value>` - Dockerhub image tag (default: "latest")</br>
`--json/-j` - Specify terminal output

Subcommands:</br>

`remote` - Install a remote deployment of Codewind
> **Flags:**
> --namespace,-n value          Kubernetes namespace to install into
> --session,-ses value          Codewind session secret to encrypt session store
> --ingress,-i value            Ingress Domain eg: 10.22.33.44.nip.io
> --kadminuser,-au value        Keycloak admin user
> --kadminpass,-ap value        Keycloak admin password
> --kdevuser,-du value          Keycloak developer username
> --kdevpass,-dp value          Keycloak developer username initial password
> --krealm,-r value             Keycloak realm to setup
> --kclient,-c value            Keycloak client to setup
> --pvcsize,-p value            Codewind PVC size (integer between 1 and 999 Gigabytes)
> --kurl value                  Don't deploy a new Keycloak pod, use an existing one at this URL
> --konly                       Install a deployment of Keycloak only

### start

`--tag/-t <value>` - Dockerhub image tag (default: "latest")</br>
`--debug/-d` - Add debug output

### status

`--json/-j` - Specify terminal output

### stop

>**Note:** No additional flags

### stop-all

>**Note:** No additional flags

### remove

`--tag/-t <value>` - Dockerhub image tag.</br>
**Note:** Failing to specify a `--tag`, will remove all Codewind images on the host machine.

### templates

>**Note:** No additional flags

Subcommands:</br>

`list/ls` - List available templates

### version

> **Flags:**
> --conid value                 Connection ID (see the connections cmd)

## sectoken

Subcommands:</br>

`get/g` - Authenticate and obtain an access_token.

>**Note 1:**: The preferred way to authenticate is by supplying just the connection ID (conid) and username. In this mode the command will use the stored password from the platform keyring
>**Note 2:**: If you dont have a connection ID (conid) you must supply use the host, realm and client flags
>**Note 3:**: You can use a combination of both the connection ID (conid) and host/realm/client flags. In this mode, the host/realm/client flags take precedence override the connection defaults
>**Note 4:**: The password flag is optional when used with the connection ID (conid) flag and when a password already exists in the platform keyring. Including the password flag will update the keychain password after a successful login or add a password to the keychain if one does not exist

> **Flags:**
> --host value                  URL or ingress to Keycloak service
> --realm value                 Application realm
> --username value              Account Username
> --password value              Account Password
> --client value                Client
> --conid  value                Connection ID (see the connections cmd)

## sectoken

Subcommands:</br>

`refresh/r` - Refresh access_token using cached refresh_token

Refresh tokens are automatically stored in the platform keychain. This command will use the refresh token to obtain a new access token from the authentication service. The access_token can then be used by curl or socket connections when accessing Codewind.

> **Flags:**
> --conid  value                Connection ID (see the connections cmd)

## secrealm

Subcommands:</br>

`create/c` - Create a new realm (requires either admin_token or username/password)

> **Flags:**
> --host value                   URL or ingress to Keycloak service
> --newrealm value               Application realm to be created
> --accesstoken value            Admin access_token

## secclient

Subcommands:</br>

`create/c` - Create a new client in an existing Keycloak realm (requires either admin_token or username/password)

> --host value                   URL or ingress to Keycloak service
> --realm value                  Application realm where client should be created
> --newclient value              New client ID to create
> --redirect value               Allowed redirect callback URL eg: `http://127.0.0.1:9090/*`
> --accesstoken value            Admin access_token

`get/g` - Get client id (requires either admin_token or username/password)

> --host value                   URL or ingress to Keycloak service
> --realm value                  Application realm
> --clientid value               Client ID to retrieve
> --accesstoken value            Admin access_token
> --username value               Admin Username
> --password value               Admin Password

`secret/s` - Get client secret (requires either admin_token or username/password)

> --host value                   URL or ingress to Keycloak service
> --realm value                  Application realm
> --clientid value               Client ID to retrieve
> --accesstoken value            Admin access_token
> --username value               Admin Username
> --password value               Admin Password

## seckeyring

Subcommands:</br>

`update/u` - Add new or update existing Codewind credentials key in keyring

> --conid  `<value>`              Connection ID (see the connections cmd)
> --username `<value>`            Username
> --password `<value>`            Password

`validate/v` - Checks if credentials key exist in the keyring

> --conid  `<value>`              Connection ID (see the connections cmd)
> --username `<value>`            Username

## secuser

Subcommands:</br>

`create/c` - Create a new user in an existing Keycloak realm (requires either admin_token or username/password)

> --host value                   URL or ingress to Keycloak service
> --realm value                  Application realm
> --accesstoken value            Admin access_token
> --username value               Admin Username
> --password value               Admin Password
> --name value                   Username to add

`get/g` - Gets an existing Keycloak user from an existing realm (requires either admin_token or username/password)

> --host value                   URL or ingress to Keycloak service
> --realm value                  Application realm
> --accesstoken value            Admin access_token
> --username value               Admin Username
> --password value               Admin Password
> --name value                   Username to query

`setpw/p` - Reset an existing users password (requires either admin_token or username/password)

> --host value                   URL or ingress to Keycloak service
> --realm value                  Application realm
> --accesstoken value            Admin access_token
> --username value               Admin Username
> --password value               Admin Password
> --name value                   Username to query
> --newpw value                  New replacement password

`addrole/p` - Adds an existing role to a user (requires either admin_token or username/password)

> --host value                   URL or ingress to Keycloak service
> --realm value                  Application realm
> --accesstoken value            Admin access_token
> --name value                   Username to target
> --role value                   Name of an existing role to add
## help

`--help/-h` - Shows a list of commands or help for one command

## Contributing

Submit issues and contributions:

1. [Submitting issues](https://github.com/eclipse/codewind/issues)
2. [Contributing](CONTRIBUTING.md)
