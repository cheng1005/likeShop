package pkg

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/smartwalle/alipay/v3"
	"strconv"
)

func AliPay(title string, total int) string {
	var (
		appId      = "9021000143642539"
		privateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDF0n1CdY6oYCF2gRJ5r1jE/AyCkolN9/Y16pKbckef5q7P0K3nXyq9oYFCu3+AI3dqKRSfrluRjz9pU7bsHJvXuwOTdNS37G0OZT1wjTwQKLl6oqsS6tNK/g4wJxNMnum6vlGxa916HCRErY9kbrHIU6riib879WQSDyeAzlNFFX+lV7pJneq7EajjMO3kKpqs72EIvkpeiOwAFMrSioVMfIn/nbJrKyTMlq6Dac9aIf4sgMtwRWXVk5ExD0S53caZRvy1cPOan4Ts/HYuvtzxgkoQqEyCd+G4zKwkLYL4vOIzS2jPL/6bCfpuRsG4uTi3E8pQvfwGpCbUen7oDh9TAgMBAAECggEBAITnW5xQ2p0oqJKzOLRkJ+F/7O+0f2bQlqu2gvTInHJDS2L3Mj7l4xMv9GXfDlQKf+nk26REQeTTNqkOwzByMAmJ7Wcd8OP4hXRMdDCv0pCQq7xGmWTGO1qAwQt81Sx+/0ylHMF0VIH7cipp6JJzBCIwg/hazw37WVYAtGuHMs6P6oBRSJitgNphttZkk5iyodL9HOS+1kZo+fGKukFi/VGYVklRjGmWaFDnbVteuHP2S5OQmUff4x+AYnKuX3RHQ3nXeb28iUlpKi9ijiNoJULdEtwTxFKRCO8CDZwQmO4566r2s22ekclycSd4QXJb6RLcO8Gj3CuuvqfpvUAUv/ECgYEA6DO7y+t7czpdK01j2bCDsRHKVkJ9zGPs3pQqk2gy64+M0zWeatqahH8HAYQPP5mT/2PD0JpfdqkVH2pS7ANh71vs+qZS7yFrmeKIRz0pnsu8IvQUXZ3ufkEg3XQQ3VyCPKwBNthtIFQ+yylCrbSmQrY+gIZlNsnxV9qVyHFAl68CgYEA2hi79k9mnLbj53wVGAggHNfW9TRZ/MGrY2Ywl8qMcprd+D9dK0oUiiqbjvrXasUP6SrtfR4mhkEFXpVMaRH71h3tEEv9kXT3qXQJJXp1quz6k0vAM0cRcbNOwMeWqIYzK418UrUtRZtRx8D3g58/ZYD7h9QvyATaeJuhikbtt50CgYBStZq/GDHw6Wkmc3qNoAJIoD2iXCzTEQTeg4hm2UbKAWf4E7FJ/nPrnhOwxoln7hx5a9/j/hji+3c6qXS//LR1vhi7b2M7KnxZZAeG2JJqrcdl6+sVLZw1/JzoRN98+eaqFdqY3p6AVYHIe/n0RdSzDpdU0Ipc6bG9yW5w52ZKoQKBgB54PipXRXq3gnegukG1QOdGsF1phvjtBcjJShiqz5xjToeUMqwNmXkEIt+C4/2ismpc+Pj18WPoSZvBN/+l540ueSRBuZdMxB0EfRKjXTUYDpnth1iioSlFZ8c/GylINnXrOmdDsKN845wItYvJ/81qi7maR2Kmau9WrubmqedxAoGBAMTvsoRejr3uEWIDtD7LFAjuF6cencE/IOP3WvBYU2Djm33gbX7sSyg2gktbh+4b7R5qZM+33I30eoRy3N664qg6+i816Ld4TaWspDoHuNo6EST2M4oWCbSOVrSXeLoGsgsTUujB6vrBDm8gsfMmDH7FGk37NW4XRqi60rF17mIS" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	)
	client, _ := alipay.New(appId, privateKey, false)
	id := uuid.NewString()
	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://xxx"
	p.Subject = title
	p.OutTradeNo = id
	p.TotalAmount = strconv.Itoa(total)
	p.ProductCode = "QUICK_WAP_WAY"

	var url, err = client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}
	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
	return payURL
}
