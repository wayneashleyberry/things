package url

const Protocol = "things:///"

const ShowInbox = Protocol + "show?id=inbox"
const ShowToday = Protocol + "show?id=today"
const ShowUpcoming = Protocol + "show?id=upcoming"
const ShowAnytime = Protocol + "show?id=anytime"
const ShowSomeday = Protocol + "show?id=someday"

const ShowLogbook = Protocol + "show?id=logbook"
const ShowTrash = Protocol + "show?id=trash"

const ShowID = Protocol + "show?id=%s"
const ShowQuery = Protocol + "show?query=%s"
