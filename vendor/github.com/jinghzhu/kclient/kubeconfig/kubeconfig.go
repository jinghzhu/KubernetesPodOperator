package kubeconfig

import (
	"encoding/base64"
)

func DecodeKubeconfigBase64FromStrToStr(kubeconfig string) (string, error) {
	byteArr, err := base64.StdEncoding.DecodeString(kubeconfig)
	if err != nil {
		return "", err
	}

	return string(byteArr), nil
}

func DecodeKubeconfigBase64FromStrToBytes(kubeconfig string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(kubeconfig)
}

func DecodeKubeconfigBase64FromBytesToStr(kubeconfig []byte) (string, error) {
	return DecodeKubeconfigBase64FromStrToStr(string(kubeconfig))
}

func DecodeKubeconfigBase64FromBytesToBytes(src []byte) (dst []byte, err error) {
	_, err = base64.StdEncoding.Decode(dst, src)

	return dst, err
}

func EncodeKubeconfigBase64FromStrToStr(kubeconfig string) string {
	return base64.StdEncoding.EncodeToString([]byte(kubeconfig))
}

func EncodeKubeconfigBase64FromStrToBytes(src string) (dst []byte) {
	base64.StdEncoding.Encode(dst, []byte(src))

	return dst
}

func EncodeKubeconfigBase64FromBytesToStr(kubeconfig []byte) string {
	return base64.StdEncoding.EncodeToString(kubeconfig)
}

func EncodeKubeconfigBase64FromBytesToBytes(src []byte) (dst []byte) {
	base64.StdEncoding.Encode(dst, src)

	return dst
}
