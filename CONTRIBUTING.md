# Install and run

WARNING !!! Do not use system backup in Rapido, there is a [critical bug](https://framagit.org/InfoLibre/rapido/-/issues/110) which causes the partition to fill up with an ever growing file.


Install [Go](https://golang.org), [GCC](https://gcc.gnu.org), [Git](https://git-scm.com) and [npm](https://npmjs.com):
```bash
sudo apt-get update && sudo apt-get -y install golang gcc git npm
```

Download current version of Rapido with Git, or [older versions](https://framagit.org/InfoLibre/rapido/-/tags):
```bash
git clone https://framagit.org/InfoLibre/rapido.git .
```

Download and install Go dependencies, start API server:
```bash
go run -v main.go
```
By default, API server uses port 8080. If you already use this port number for another application, you can change it in [main.go](main.go) and [client/config/dev.env.js](client/config/dev.env.js) files.

In another terminal, install Javascript dependencies and start front end server:
``` bash
cd client
npm install
npm run dev
```
By default, front end server uses port 3000. Rapido will be available at [http://localhost:3000](http://localhost:3000).
If you already use this port number for another application, you can change it in [client/config/index.js](client/config/index.js).


# Report a bug / Request a feature

You can report a bug or ask for a new feature.
Search in the [list of open issues](https://framagit.org/InfoLibre/rapido/issues?scope=all&state=opened) to see if it is already reported by someone else.
If it is the case, you can add information or feedback.
Otherwise, you can [create a new ticket](https://framagit.org/InfoLibre/rapido/issues/new).

# Translate and develop

<a href="https://matrix.to/#/#rapido:matrix.org">Chat with us on <img src="docs/Matrix.svg" width="53.821521731" height="23" alt="Matrix"></a>

List of languages and translation files are in [lang/](lang/) directory. You can add a JSON file for each language, named according to [ISO 639-3](https://iso639-3.sil.org/code_tables/639/data) standard and otherwise, according to [ELP](https://endangeredlanguages.com/lang/region) standard.

Update [npm packages](https://framagit.org/InfoLibre/rapido/-/blob/master/client/package.json) and [Go modules](https://framagit.org/InfoLibre/rapido/-/blob/master/go.mod) with their latest version. Verify that they are all distributed under a Free Software License or under an Open Source license and that they do not contain any tracker, external API, external font or other external ressources. Do not use Font Awesome but [Fork Awesome](https://forkaweso.me/Fork-Awesome/), do not use Google Fonts API but [self hosted Google Fonts](https://google-webfonts-helper.herokuapp.com).

Use CSS or the CSS preprocessor [Stylus](https://stylus-lang.com). Do not use Less, Sass/SCSS.

Use a text editor that comes bundled with native support for [EditorConfig](https://editorconfig.org) files or that can use it as a plugin. Use the [.editorconfig](.editorconfig) file at the root of this project.

Use [DB Browser for SQLite](https://sqlitebrowser.org) to edit [database/Rapido.sqlite](database/Rapido.sqlite).


### Document

Comment all your code and update comments used to generate the [API documentation](models/docs/) with [apiDoc](https://apidocjs.com).

### Before pushing your code to the Git server

* If you want to create a new tag with the command ``git tag vx.x.x``, follow [Semantic Versioning](https://semver.org) and update [version.json](version.json) file before creating the tag.
* Delete files, source code, comments, variables and database tables that are no longer used.
* Delete all personnal data, passwords, SMTP configuration, test data...
* Format Go source code with the command ``go fmt ./...``
