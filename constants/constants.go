package constants

const (
	IMAGE_BASE_URL     = "https://rpa-voucher-exchange.s3.ap-southeast-1.amazonaws.com/"
	JWT_DATE_FORMAT    = "2006-01-02 15:04:05" //yyyy-mm-dd HH:mm:ss
	SEARCH_DATE_FORMAT = "02-01-2006"          //dd-mm-yyyy
	ROLE_ADMIN         = 1
	ROLE_SALE          = 2
	STATUS_PENDING     = uint(1)
	STATUS_APPROVED    = uint(2)
	STATUS_REJECTED    = uint(3)
)
