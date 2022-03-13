# Intro of NAKAMA

NAKAMA is an open-source distributed social and realtime server for games and apps. It includes a
large set of services for users, data storage, and realtime client/server communication; as well as
specialized APIs like realtime multiplayer, groups/guilds, and chat.

NAKAMA has advised that if anyone wants something customize or additional, they should add RPC by
writing Lua/Go/TypeScript scripts in order to customize and change the default behaviour of their
functions.

In this repository I am using Go and I developed multiple function:

- add leaderboard
- server time
- server date

### add leaderboard

`path`:`/addleaderboard?http_key=admin_admin_admin`
<br>
`method`:`POST`
<br>
`desc`:`This api adds leaderboard to nakama`
<br>
`request model : json format`
<br>

`{"id":"LDB_1",
"metadata":{"name":"amqezi"},
"sort":"desc",
"operator":"incr",
"reset":"0 1 * * *"}
`
<br>

`Sort: desc, asc Operator: best, set, incr`
`
<br>

`response model`:
`{
"id": "LDB_1",
"metadata": {
"name": "test"
},
"sort": "desc",
"operator": "incr",
"reset": "0 1 * * *"
}`

### server time

`path`:`/servertime?http_key=admin_admin_admin`
<br>
`method`:`POST`
<br>
`desc`:`This api returns server utc time`
<br>
`response model`:
`{
"payload": "1643523929"
}`

### server date

`path`:`/serverdate?http_key=admin_admin_admin`
<br>
`method`:`POST`
<br>
`desc`:`This api returns the human-readable server date`
<br>
`response model`:
`{
"payload" :
"{\"day\":30, \"hour\":6, \"isdst\":false, \"min\":28, \"month\":1, \"sec\":49, \"wday\":0, \"yday\":30, \"year\":2022}"
}`

