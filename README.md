# Kanban metrics

- This project is composed of 2 major packages: client and server.

- Client is mainly programmed by reactjs. Some frameworks being used are listed below:

    * redux
    * router

- Server is programmed by GO langague.

## Installation
Firstly make sure that you have Node, Docker and Git installed. Next clone this repo https://github.com/NguyenHoaiPhuong/warehouse.git. You can do this by going into your shell of choice and entering

```
git clone https://github.com/NguyenHoaiPhuong/kanban.git
```

Then, access the project directory (kanban) and run docker-compose up as shown below:

```
cd kanban
docker-compose up
```

Next, open a new terminal and access the project directory again. Restore the backup database by make command:

```
make restore-backup-db
```

Finally, open your web browser and access http://localhost:5001/. At the beginning, you will be requested to login. The default username and password are ***admin***. After login successfully, you will be able to access the home page.

## Tech stack

### Frontend

* typescript
* @types/react
* @types/react-dom
* @types/jest
* @types/node
* @types/react-router-dom
* @types/redux

* @material-ui/core
* @material-ui/icons

* react-redux
* redux-thunk

* react-router-dom

- Refer to following link for "How to use Immutable.js Records with React and Redux":

https://medium.com/azendoo-team/immutable-record-react-redux-99f389ed676


- Axios Interceptors:

    https://medium.com/swlh/handling-access-and-refresh-tokens-using-axios-interceptors-3970b601a5da

- Webpack:

    https://www.pluralsight.com/guides/react-typescript-webpack

### Backend

### Docker

- Docker env and arg:

    https://vsupalov.com/docker-arg-env-variable-guide/#arg-and-env-availability

- Add user and group:

    https://stackoverflow.com/questions/49955097/how-do-i-add-a-user-when-im-using-alpine-as-a-base-image

- docker for nodejs:

    https://github.com/nodejs/docker-node/blob/master/docs/BestPractices.md#global-npm-dependencies

- RUN vs CMD vs ENTRYPOINT

    https://goinbigdata.com/docker-run-vs-cmd-vs-entrypoint/
