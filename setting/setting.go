package setting

/* ----- Server Setting ----- */
const ServerPort string = "443"

const ModelPath string = "./models/model_list.json"
const UrlPath string = "./urls/url_list.json"

/* ----- Triton Server Setting ----- */
const TritonUrl string = "localhost:2000"

const batchSize int = 1
const Samples int = 1
const Steps int = 45
const GuidanceScale float64 = 7.5
const seed int = -1
