package easycodefgo

type messageConstant struct {
	Code         string
	Message      string
	ExtraMessage string
}

func new(code, msg string) *messageConstant {
	return &messageConstant{code, msg, ""}
}

var (
	messageOK                  = new("CF-00000", "성공")
	messageInvalidJson         = new("CF-00002", "json형식이 올바르지 않습니다.")
	messageInvalidParameter    = new("CF-00007", "요청 파라미터가 올바르지 않습니다.")
	messageUnsupportedEncoding = new("CF-00009", "지원하지 않는 형식으로 인코딩된 문자열입니다.")
	messageEmptyClientInfo     = new("CF-00014", "상품 요청을 위해서는 클라이언트 정보가 필요합니다. 클라이언트 아이디와 시크릿 정보를 설정하세요.")
	messageEmptyPublicKey      = new("CF-00015", "상품 요청을 위해서는 퍼블릭키가 필요합니다. 퍼블릭키 정보를 설정하세요.")
	messageInvalid2WayInfo     = new("CF-03003", "2WAY 요청 처리를 위한 정보가 올바르지 않습니다. 응답으로 받은 항목을 그대로 2way요청 항목에 포함해야 합니다.")
	messageInvalid2WayKeyword  = new("CF-03004", "추가 인증(2Way)을 위한 요청은 requestCertification메서드를 사용해야 합니다.")
	messageBadRequest          = new("CF-00400", "클라이언트 요청 오류로 인해 요청을 처리 할 수 ​​없습니다.")
	messageUnauthorized        = new("CF-00401", "요청 권한이 없습니다.")
	messageForbidden           = new("CF-00403", "잘못된 요청입니다.")
	messageNotFound            = new("CF-00404", "요청하신 페이지(Resource)를 찾을 수 없습니다.")
	messageMethodNotAllowed    = new("CF-00405", "요청하신 방법(Method)이 잘못되었습니다.")
	messageLibrarySenderError  = new("CF-09980", "통신 요청에 실패했습니다. 응답정보를 확인하시고 올바른 요청을 시도하세요.")
	messageServerError         = new("CF-09999", "서버 처리중 에러가 발생 했습니다. 관리자에게 문의하세요.")
)
