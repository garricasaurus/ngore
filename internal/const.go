package internal

const LocationLogin = "login.php"
const LocationIndex = "index.php"
const LocationLoginProblem = "problema"

const UrlLogin = "/login.php"
const UrlIndex = "/index.php"
const UrlTorrents = "/torrents.php"
const UrlActivity = "/hitnrun.php"

const ErrUserNotLoggedIn = "user is not logged in"
const ErrApiKeyEmpty = "api key is empty"
const ErrLoginMissingCredentials = "login failed: user name or password is empty"
const ErrLoginInvalidCredentials = "login failed: invalid BasicAuth"
const ErrLoginUnexpectedResponse = "login failed: unexpected response"
const ErrLoginUnableToFetchIndex = "login failed: unable to fetch index page"
const ErrLoginKeyMissing = "login failed: unable to find login key in response"
const ErrLoginKeyParse = "login failed: login key cannot be parsed"
const ErrSearchUnexpectedResponseCode = "search failed: unexpected response code: %d"
const ErrActivityUnexpectedResponseCode = "fetching activity failed: unexpected response code: %d"
const ErrDownloadUnexpectedResponseCode = "download failed: unexpected response code: %d"
