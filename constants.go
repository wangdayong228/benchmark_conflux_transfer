package main

const CHAIN_ID_TESTNET uint32 = 1
const CHAIN_ID_8888 uint32 = 8888

func GetUrls(chainID uint32) []string {
	switch chainID {
	case CHAIN_ID_TESTNET:
		return []string{
			"https://test.confluxrpc.com/09K6dkMRr9xyz6suyNgqTR",
		}
	case CHAIN_ID_8888:
		return []string{
			"http://net8888cfx.confluxrpc.com",
			"http://101.132.158.162:12537",
			"http://39.100.97.209:12537",
			"http://47.92.105.52:12537",
			"http://47.92.7.84:12537",
		}
	}
	return nil
}
