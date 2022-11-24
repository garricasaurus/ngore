package ncgore

const locationLogin = "login.php"
const locationIndex = "index.php"
const locationLoginProblem = "problema"

const urlLogin = "/login.php"
const urlIndex = "/index.php"
const urlSearch = "/torrents.php"

const errUserNotLoggedIn = "user is not logged in"
const errLoginMissingCredentials = "login failed: user name or password is empty"
const errLoginInvalidCredentials = "login failed: invalid BasicAuth"
const errLoginUnexpectedResponse = "login failed: unexpected response"
const errSearchUnexpectedResponseCode = "search failed: unexpected response code: %d"
