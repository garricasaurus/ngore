package internal

const LocationLogin = "login.php"
const LocationIndex = "index.php"
const LocationLoginProblem = "problema"

const UrlLogin = "/login.php"
const UrlIndex = "/index.php"
const UrlSearch = "/torrents.php"
const UrlActivity = "/hitnrun.php"

const ErrUserNotLoggedIn = "user is not logged in"
const ErrLoginMissingCredentials = "login failed: user name or password is empty"
const ErrLoginInvalidCredentials = "login failed: invalid BasicAuth"
const ErrLoginUnexpectedResponse = "login failed: unexpected response"
const ErrSearchUnexpectedResponseCode = "search failed: unexpected response code: %d"
const ErrActivityUnexpectedResponseCode = "fetching activity failed: unexpected response code: %d"
